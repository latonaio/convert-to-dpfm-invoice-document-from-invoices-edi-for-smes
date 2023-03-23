package dpfm_api_processing_formatter

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"
)

type ProcessingFormatter struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewProcessingFormatter(ctx context.Context, db *database.Mysql, l *logger.Logger) *ProcessingFormatter {
	return &ProcessingFormatter{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (p *ProcessingFormatter) ProcessingFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *ProcessingFormatterSDC,
) error {
	var err error
	var e error

	if bpIDIsNull(sdc) {
		return xerrors.New("business_partner is null")
	}

	wg := sync.WaitGroup{}

	psdc.Header = p.Header(sdc, psdc)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.ConversionProcessingHeader, e = p.ConversionProcessingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.Item = p.Item(sdc, psdc)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// Ref: Item
			psdc.ConversionProcessingItem, e = p.ConversionProcessingItem(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}(wg)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.Address = p.Address(sdc, psdc)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.Partner = p.Partner(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	p.l.Info(psdc)

	return nil
}

func (p *ProcessingFormatter) Header(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Header {
	data := sdc.Header

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	res := Header{
		ConvertingInvoiceDocument:         data.ExchangedInvoiceDocumentIdentifier,
		CreationDate:                      systemDate,
		CreationTime:                      systemTime,
		LastChangeDate:                    systemDate,
		LastChangeTime:                    systemTime,
		ConvertingBillToParty:             data.TradeBillToPartyRegisteredIdentifier,
		ConvertingBillFromParty:           data.TradeBillFromPartyIdentifier,
		ConvertingPayer:                   data.TradeBillToPartyRegisteredIdentifier,
		ConvertingPayee:                   data.TradeBillFromPartyIdentifier,
		InvoiceDocumentDate:               data.ExchangedInvoiceDocumentIssueDate,
		InvoicePeriodStartDate:            data.InvoiceDocumentSpecifiedPeriodStartDate,
		InvoicePeriodEndDate:              data.InvoiceDocumentSpecifiedPeriodEndDate,
		AccountingPostingDate:             data.ExchangedInvoiceDocumentIssueDate,
		TotalNetAmount:                    data.TradeInvoiceDocumentSettlementMonetarySummationTaxBasisTotalAmount,
		TotalTaxAmount:                    data.TradeTaxCalculatedAmount,
		TotalGrossAmount:                  data.TradeInvoiceDocumentSettlementMonetarySummationGrandTotalAmount,
		TransactionCurrency:               data.SupplyChainTradeSettlementPaymentCurrencyCode,
		PaymentTerms:                      data.TradePaymentTermsTypeCode,
		PaymentDueDate:                    data.TradePaymentTermsDueDate,
		ConvertingPaymentMethod:           data.TradePaymentTermsTypeCode,
		DocumentHeaderText:                data.InvoiceDocument,
		HeaderIsCleared:                   getBoolPtr(false),
		HeaderPaymentBlockStatus:          getBoolPtr(false),
		HeaderPaymentRequisitionIsCreated: getBoolPtr(false),
		IsCancelled:                       getBoolPtr(false),
	}

	return &res
}

func (p *ProcessingFormatter) ConversionProcessingHeader(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*ConversionProcessingHeader, error) {
	dataKey := make([]*ConversionProcessingKey, 0)

	p.appendDataKey(&dataKey, sdc, "ExchangedInvoiceDocumentIdentifier", "InvoiceDocument", psdc.Header.ConvertingInvoiceDocument)
	p.appendDataKey(&dataKey, sdc, "TradeBillToPartyRegisteredIdentifier", "BillToParty", psdc.Header.ConvertingBillToParty)
	p.appendDataKey(&dataKey, sdc, "TradeBillFromPartyIdentifier", "BillFromParty", psdc.Header.ConvertingBillFromParty)
	p.appendDataKey(&dataKey, sdc, "TradePaymentTermsTypeCode", "PaymentMethod", psdc.Header.ConvertingPaymentMethod)

	dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
	if err != nil {
		return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
	}

	data, err := p.ConvertToConversionProcessingHeader(dataKey, dataQueryGets)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
	}

	return data, nil
}

func (psdc *ProcessingFormatter) ConvertToConversionProcessingHeader(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingHeader, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("Value of %s is not in the database", v.LabelConvertTo)
		}
	}

	res := &ConversionProcessingHeader{}

	if _, ok := data["InvoiceDocument"]; ok {
		res.ConvertingInvoiceDocument = data["InvoiceDocument"].CodeConvertFromString
		res.ConvertedInvoiceDocument = data["InvoiceDocument"].CodeConvertToInt
	}
	if _, ok := data["BillToParty"]; ok {
		res.ConvertingBillToParty = data["BillToParty"].CodeConvertFromString
		res.ConvertedBillToParty = data["BillToParty"].CodeConvertToInt
		res.ConvertedPayer = data["BillToParty"].CodeConvertToInt
	}
	if _, ok := data["BillFromParty"]; ok {
		res.ConvertingBillFromParty = data["BillFromParty"].CodeConvertFromString
		res.ConvertedBillFromParty = data["BillFromParty"].CodeConvertToInt
		res.ConvertedPayee = data["BillFromParty"].CodeConvertToInt
	}
	if _, ok := data["Payer"]; ok {
		res.ConvertingPayer = data["Payer"].CodeConvertFromString
		res.ConvertedPayer = data["Payer"].CodeConvertToInt
	}
	if _, ok := data["Payee"]; ok {
		res.ConvertingPayee = data["Payee"].CodeConvertFromString
		res.ConvertedPayee = data["Payee"].CodeConvertToInt
	}
	if _, ok := data["PaymentMethod"]; ok {
		res.ConvertingPaymentMethod = data["PaymentMethod"].CodeConvertFromString
		res.ConvertedPaymentMethod = data["PaymentMethod"].CodeConvertToString
	}

	return res, nil
}

func (p *ProcessingFormatter) Item(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) []*Item {
	res := make([]*Item, 0)
	dataHeader := psdc.Header
	data := sdc.Header.Item

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	for _, data := range data {

		res = append(res, &Item{
			ConvertingInvoiceDocument:              dataHeader.ConvertingInvoiceDocument,
			ConvertingInvoiceDocumentItem:          data.InvoiceDocumentItemIdentifier,
			InvoiceDocumentItemText:                data.TradeProductName,
			ConvertingProduct:                      data.TradeProductIdentifier,
			CreationDate:                           systemDate,
			CreationTime:                           systemTime,
			LastChangeDate:                         systemDate,
			LastChangeTime:                         systemTime,
			ConvertingBuyer:                        sdc.Header.TradeBuyerIdentifier,
			ConvertingSeller:                       sdc.Header.TradeSellerIdentifier,
			ConvertingDeliverToParty:               data.TradeShipToPartyIdentifier,
			InvoiceQuantity:                        data.SupplyChainTradeDeliveryItemLineProductBilledQuantity,
			InvoiceQuantityUnit:                    data.SupplyChainTradeDeliveryItemLineProductQuantityUnitCode,
			InvoiceQuantityInBaseUnit:              data.TradePriceBasisQuantity,
			BaseUnit:                               data.TradePriceBasisQuantityUnitCode,
			NetAmount:                              data.ItemLineTradeTaxBasisAmount,
			GrossAmount:                            data.ItemLineTradeTaxGrandTotalAmount,
			TransactionCurrency:                    sdc.Header.SupplyChainTradeSettlementPaymentCurrencyCode,
			ConvertingTransactionTaxClassification: data.ItemLineTradeTaxCategoryCode,
			ConvertingProject:                      sdc.Header.ProjectIdentifier,
			ConvertingOrderID:                      data.ItemLineReferencedOrdersDocumentIssuerAssignedIdentifier,
			ConvertingOrderItem:                    data.ItemLineReferencedOrdersDocumentItemLineIssuerAssignedIdentifier,
			InvoicePeriodStartDate:                 data.ItemLineSpecifiedPeriodStartDate,
			InvoicePeriodEndDate:                   data.ItemLineSpecifiedPeriodEndDate,
			ConvertingDeliveryDocument:             data.ItemLineReferencedOrdersDocumentIssuerAssignedIdentifier,
			ConvertingDeliveryDocumentItem:         data.ItemLineReferencedOrdersDocumentItemLineIssuerAssignedIdentifier,
			ConvertingOriginDocument:               data.ItemLineReferencedOrdersDocumentIssuerAssignedIdentifier,
			ConvertingOriginDocumentItem:           data.ItemLineReferencedOrdersDocumentItemLineIssuerAssignedIdentifier,
			ConvertingReferenceDocument:            data.ItemLineReferencedOrdersDocumentIssuerAssignedIdentifier,
			ConvertingReferenceDocumentItem:        data.ItemLineReferencedOrdersDocumentItemLineIssuerAssignedIdentifier,
			ItemPaymentRequisitionIsCreated:        getBoolPtr(false),
			ItemIsCleared:                          getBoolPtr(false),
			ItemPaymentBlockStatus:                 getBoolPtr(false),
			IsCancelled:                            getBoolPtr(false),
		})
	}

	return res
}

func (p *ProcessingFormatter) ConversionProcessingItem(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItem, error) {
	data := make([]*ConversionProcessingItem, 0)

	for _, item := range psdc.Item {
		dataKey := make([]*ConversionProcessingKey, 0)

		p.appendDataKey(&dataKey, sdc, "InvoiceDocumentItemIdentifier", "InvoiceDocumentItem", item.ConvertingInvoiceDocumentItem)
		p.appendDataKey(&dataKey, sdc, "TradeProductIdentifier", "Product", item.ConvertingProduct)
		p.appendDataKey(&dataKey, sdc, "TradeBuyerIdentifier ", "Buyer", item.ConvertingBuyer)
		p.appendDataKey(&dataKey, sdc, "TradeSellerIdentifier", "Seller", item.ConvertingSeller)
		p.appendDataKey(&dataKey, sdc, "TradeShipToPartyIdentifier", "DeliverToParty", item.ConvertingDeliverToParty)
		p.appendDataKey(&dataKey, sdc, "ItemTradeTaxCategoryCode", "TransactionTaxClassification", item.ConvertingTransactionTaxClassification)
		p.appendDataKey(&dataKey, sdc, "ProjectIdentifier", "Project", item.ConvertingProject)
		p.appendDataKey(&dataKey, sdc, "ItemLineReferencedOrdersDocumentIssuerAssignedIdentifier", "OrderID", item.ConvertingOrderID)
		p.appendDataKey(&dataKey, sdc, "ItemLineReferencedOrdersDocumentItemLineIssuerAssignedIdentifier", "OrderItem", item.ConvertingOrderItem)

		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
		if err != nil {
			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
		}

		datum, err := p.ConvertToConversionProcessingItem(dataKey, dataQueryGets)
		if err != nil {
			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
		}

		data = append(data, datum)
	}

	return data, nil
}

func (p *ProcessingFormatter) ConvertToConversionProcessingItem(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingItem, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	res := &ConversionProcessingItem{}

	if _, ok := data["InvoiceDocumentItem"]; ok {
		res.ConvertingInvoiceDocumentItem = data["InvoiceDocumentItem"].CodeConvertFromString
		res.ConvertedInvoiceDocumentItem = data["InvoiceDocumentItem"].CodeConvertToInt
	}
	if _, ok := data["Product"]; ok {
		res.ConvertingProduct = data["Product"].CodeConvertFromString
		res.ConvertedProduct = data["Product"].CodeConvertToString
	}
	if _, ok := data["Buyer"]; ok {
		res.ConvertingBuyer = data["Buyer"].CodeConvertFromString
		res.ConvertedBuyer = data["Buyer"].CodeConvertToInt
	}
	if _, ok := data["Seller"]; ok {
		res.ConvertingSeller = data["Seller"].CodeConvertFromString
		res.ConvertedSeller = data["Seller"].CodeConvertToInt
	}
	if _, ok := data["DeliverToParty"]; ok {
		res.ConvertingDeliverToParty = data["DeliverToParty"].CodeConvertFromString
		res.ConvertedDeliverToParty = data["DeliverToParty"].CodeConvertToInt
	}
	if _, ok := data["Project"]; ok {
		res.ConvertingProject = data["Project"].CodeConvertFromString
		res.ConvertedProject = data["Project"].CodeConvertToString
	}
	if _, ok := data["TransactionTaxClassification"]; ok {
		res.ConvertingTransactionTaxClassification = data["TransactionTaxClassification"].CodeConvertFromString
		res.ConvertedTransactionTaxClassification = data["TransactionTaxClassification"].CodeConvertToString
	}
	if _, ok := data["OrderID"]; ok {
		res.ConvertingOrderID = data["OrderID"].CodeConvertFromString
		res.ConvertedOrderID = data["OrderID"].CodeConvertToInt
		res.ConvertedDeliveryDocument = data["OrderID"].CodeConvertToInt
		res.ConvertedOriginDocument = data["OrderID"].CodeConvertToInt
		res.ConvertedReferenceDocument = data["OrderID"].CodeConvertToInt
	}
	if _, ok := data["OrderItem"]; ok {
		res.ConvertingOrderItem = data["OrderItem"].CodeConvertFromString
		res.ConvertedOrderItem = data["OrderItem"].CodeConvertToInt
		res.ConvertedDeliveryDocumentItem = data["OrderItem"].CodeConvertToInt
		res.ConvertedOriginDocumentItem = data["OrderItem"].CodeConvertToInt
		res.ConvertedReferenceDocumentItem = data["OrderItem"].CodeConvertToInt
	}

	return res, nil
}

func (p *ProcessingFormatter) Address(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) []*Address {
	res := make([]*Address, 0)

	deliverToPartyAddress := deliverToPartyAddress(sdc, psdc)
	if !postalCodeContains(deliverToPartyAddress.PostalCode, res) {
		res = append(res, deliverToPartyAddress)
	}

	billFromPartyAddress := billFromPartyAddress(sdc, psdc)
	if !postalCodeContains(billFromPartyAddress.PostalCode, res) {
		res = append(res, billFromPartyAddress)
	}

	buyerAddress := buyerAddress(sdc, psdc)
	if !postalCodeContains(buyerAddress.PostalCode, res) {
		res = append(res, buyerAddress)
	}

	sellerAddress := sellerAddress(sdc, psdc)
	if !postalCodeContains(sellerAddress.PostalCode, res) {
		res = append(res, sellerAddress)
	}

	return res
}

func deliverToPartyAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header
	dataInputHeader := sdc.Header
	dataInputItem := dataInputHeader.Item

	res := &Address{
		ConvertingInvoiceDocument: dataHeader.ConvertingInvoiceDocument,
		PostalCode:                dataInputItem[0].ShipToPartyAddressPostalCode,
	}

	return res
}

func billFromPartyAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header

	res := &Address{
		ConvertingInvoiceDocument: dataHeader.ConvertingInvoiceDocument,
		PostalCode:                sdc.Header.BillFromPartyAddressPostalCode,
	}

	return res
}

func buyerAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header

	res := &Address{
		ConvertingInvoiceDocument: dataHeader.ConvertingInvoiceDocument,
		PostalCode:                sdc.Header.BuyerAddressPostalCode,
	}

	return res
}

func sellerAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header

	res := &Address{
		ConvertingInvoiceDocument: dataHeader.ConvertingInvoiceDocument,
		PostalCode:                sdc.Header.SellerAddressPostalCode,
	}

	return res
}

func (p *ProcessingFormatter) Partner(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) []*Partner {
	res := make([]*Partner, 0)
	dataHeader := psdc.Header

	res = append(res, &Partner{
		ConvertingInvoiceDocument: dataHeader.ConvertingInvoiceDocument,
	})

	return res
}

func (p *ProcessingFormatter) appendDataKey(dataKey *[]*ConversionProcessingKey, sdc *dpfm_api_input_reader.SDC, labelConvertFrom string, labelConvertTo string, codeConvertFrom any) {
	switch v := codeConvertFrom.(type) {
	case int, float32:
		if v == 0 {
			return
		}
	case string:
		if v == "" {
			return
		}
	case *int, *float32:
		if v == nil {
			return
		}
	case *string:
		if v == nil || *v == "" {
			return
		}
	default:
		return
	}
	*dataKey = append(*dataKey, p.ConversionProcessingKey(sdc, labelConvertFrom, labelConvertTo, codeConvertFrom))
}

func postalCodeContains(postalCode *string, addresses []*Address) bool {
	for _, address := range addresses {
		if address.PostalCode == nil || postalCode == nil {
			return true
		}
		if *address.PostalCode == *postalCode {
			return true
		}
	}

	return false
}

func bpIDIsNull(sdc *dpfm_api_input_reader.SDC) bool {
	return sdc.BusinessPartnerID == nil
}
