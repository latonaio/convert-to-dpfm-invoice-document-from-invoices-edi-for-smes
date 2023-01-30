package dpfm_api_processing_formatter

import (
	"convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
	"strconv"
)

// データ連携基盤とInvoices EDI For SMEsとの項目マッピング
// Header
func (psdc *SDC) ConvertToMappingHeader(sdc *dpfm_api_input_reader.SDC) (*MappingHeader, error) {
	data := sdc.InvoicesEDIForSMEsHeader

	systemDate := GetSystemDatePtr()
	systemTime := GetSystemTimePtr()

	res := MappingHeader{
		CreationDate:           systemDate,
		CreationTime:           systemTime,
		LastChangeDate:         systemDate,
		LastChangeTime:         systemTime,
		InvoiceDocumentDate:    data.ExchangedInvoiceDocumentIssueDate,
		InvoicePeriodStartDate: data.InvoiceDocumentSpecifiedPeriodStartDate,
		InvoicePeriodEndDate:   data.InvoiceDocumentSpecifiedPeriodEndDate,
		AccountingPostingDate:  data.ExchangedInvoiceDocumentIssueDate,
		TotalNetAmount:         data.TradeInvoiceDocumentSettlementMonetarySummationTaxBasisTotalAmount,
		TotalTaxAmount:         data.TradeTaxCalculatedAmount,
		TotalGrossAmount:       data.TradeInvoiceDocumentSettlementMonetarySummationGrandTotalAmount,
		TransactionCurrency:    data.SupplyChainTradeSettlementPaymentCurrencyCode,
		PaymentTerms:           data.TradePaymentTermsTypeCode,
		PaymentDueDate:         data.TradeSettlementPaymentMeansTypeCode,
		DocumentHeaderText:     data.InvoiceDocument,
	}

	return &res, nil
}

// Item
func (psdc *SDC) ConvertToMappingItem(sdc *dpfm_api_input_reader.SDC) (*[]MappingItem, error) {
	var res []MappingItem
	data := sdc.InvoicesEDIForSMEsHeader
	dataItem := sdc.InvoicesEDIForSMEsHeader.InvoicesEDIForSMEsItem

	for _, dataItem := range dataItem {
		systemDate := GetSystemDatePtr()
		systemTime := GetSystemTimePtr()

		res = append(res, MappingItem{
			ExchangedInvoiceDocumentIdentifier: data.ExchangedInvoiceDocumentIdentifier,
			InvoiceDocumentItemText:            dataItem.TradeProductName,
			CreationDate:                       systemDate,
			CreationTime:                       systemTime,
			LastChangeDate:                     systemDate,
			LastChangeTime:                     systemTime,
			InvoiceQuantity:                    dataItem.SupplyChainTradeDeliveryItemLineProductBilledQuantity,
			InvoiceQuantityUnit:                dataItem.SupplyChainTradeDeliveryItemLineProductQuantityUnitCode,
			InvoiceQuantityInBaseUnit:          dataItem.TradePriceBasisQuantity,
			BaseUnit:                           dataItem.TradePriceBasisQuantityUnitCode,
			NetAmount:                          dataItem.ItemLineTradeTaxBasisAmount,
			GrossAmount:                        dataItem.ItemLineTradeTaxGrandTotalAmount,
			TransactionCurrency:                data.SupplyChainTradeSettlementPaymentCurrencyCode,
			InvoicePeriodStartDate:             dataItem.ItemLineSpecifiedPeriodStartDate,
			InvoicePeriodEndDate:               dataItem.ItemLineSpecifiedPeriodEndDate,
		})
	}

	return &res, nil
}

// ItemPricingElement
func (psdc *SDC) ConvertToMappingItemPricingElement(sdc *dpfm_api_input_reader.SDC) (*[]MappingItemPricingElement, error) {
	var res []MappingItemPricingElement
	data := sdc.InvoicesEDIForSMEsHeader
	dataItem := sdc.InvoicesEDIForSMEsHeader.InvoicesEDIForSMEsItem

	for _, dataItem := range dataItem {

		res = append(res, MappingItemPricingElement{
			ExchangedInvoiceDocumentIdentifier: data.ExchangedInvoiceDocumentIdentifier,
			InvoiceDocumentItemIdentifier:      dataItem.InvoiceDocumentItemIdentifier,
			ConditionRateValue:                 dataItem.TradePriceChargeAmount,
			ConditionCurrency:                  data.SupplyChainTradeSettlementPaymentCurrencyCode,
			ConditionQuantity:                  dataItem.TradePriceBasisQuantity,
			ConditionQuantityUnit:              dataItem.TradePriceBasisQuantityUnitCode,
			ConditionAmount:                    dataItem.ItemLineTradeTaxGrandTotalAmount,
			TransactionCurrency:                data.SupplyChainTradeSettlementPaymentCurrencyCode,
		})
	}

	return &res, nil
}

// Partner
func (psdc *SDC) ConvertToMappingPartner(sdc *dpfm_api_input_reader.SDC) (*[]MappingPartner, error) {
	var res []MappingPartner
	data := sdc.InvoicesEDIForSMEsHeader
	dataPartner := sdc.InvoicesEDIForSMEsHeader.InvoicesEDIForSMEsItem

	for range dataPartner {
		res = append(res, MappingPartner{
			ExchangedInvoiceDocumentIdentifier: data.ExchangedInvoiceDocumentIdentifier,
		})
	}

	return &res, nil
}

// 番号処理
func (psdc *SDC) ConvertToCodeConversionKey(sdc *dpfm_api_input_reader.SDC, labelConvertFrom, labelConvertTo string, codeConvertFrom any) *CodeConversionKey {
	pm := &requests.CodeConversionKey{
		SystemConvertTo:   "DPFM",
		SystemConvertFrom: "SAP",
		BusinessPartner:   *sdc.BusinessPartnerID,
	}

	pm.LabelConvertFrom = labelConvertFrom
	pm.LabelConvertTo = labelConvertTo

	switch codeConvertFrom := codeConvertFrom.(type) {
	case string:
		pm.CodeConvertFrom = codeConvertFrom
	case int:
		pm.CodeConvertFrom = strconv.FormatInt(int64(codeConvertFrom), 10)
	case float32:
		pm.CodeConvertFrom = strconv.FormatFloat(float64(codeConvertFrom), 'f', -1, 32)
	case bool:
		pm.CodeConvertFrom = strconv.FormatBool(codeConvertFrom)
	case *string:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = *codeConvertFrom
		}
	case *int:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatInt(int64(*codeConvertFrom), 10)
		}
	case *float32:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatFloat(float64(*codeConvertFrom), 'f', -1, 32)
		}
	case *bool:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatBool(*codeConvertFrom)
		}
	}

	data := pm
	res := CodeConversionKey{
		SystemConvertTo:   data.SystemConvertTo,
		SystemConvertFrom: data.SystemConvertFrom,
		LabelConvertTo:    data.LabelConvertTo,
		LabelConvertFrom:  data.LabelConvertFrom,
		CodeConvertFrom:   data.CodeConvertFrom,
		BusinessPartner:   data.BusinessPartner,
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionQueryGets(rows *sql.Rows) (*[]CodeConversionQueryGets, error) {
	defer rows.Close()
	var res []CodeConversionQueryGets

	i := 0
	for rows.Next() {
		i++
		pm := &requests.CodeConversionQueryGets{}

		err := rows.Scan(
			&pm.CodeConversionID,
			&pm.SystemConvertTo,
			&pm.SystemConvertFrom,
			&pm.LabelConvertTo,
			&pm.LabelConvertFrom,
			&pm.CodeConvertFrom,
			&pm.CodeConvertTo,
			&pm.BusinessPartner,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, CodeConversionQueryGets{
			CodeConversionID:  data.CodeConversionID,
			SystemConvertTo:   data.SystemConvertTo,
			SystemConvertFrom: data.SystemConvertFrom,
			LabelConvertTo:    data.LabelConvertTo,
			LabelConvertFrom:  data.LabelConvertFrom,
			CodeConvertFrom:   data.CodeConvertFrom,
			CodeConvertTo:     data.CodeConvertTo,
			BusinessPartner:   data.BusinessPartner,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_code_conversion_code_conversion_data'テーブルに対象のレコードが存在しません。")
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionHeader(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionHeader, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionHeader{}

	pm.ExchangedInvoiceDocumentIdentifier = dataMap["InvoiceDocument"].CodeConvertFrom
	pm.InvoiceDocument, err = ParseInt(dataMap["InvoiceDocument"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.BillToParty, err = ParseIntPtr(GetStringPtr(dataMap["BillToParty"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.BillFromParty, err = ParseIntPtr(GetStringPtr(dataMap["BillFromParty"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.Payer, err = ParseIntPtr(GetStringPtr(dataMap["Payer"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.Payee, err = ParseIntPtr(GetStringPtr(dataMap["Payee"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.PaymentTerms = GetStringPtr(dataMap["PaymentTerms"].CodeConvertFrom)
	pm.PaymentMethod = GetStringPtr(dataMap["PaymentMethod"].CodeConvertFrom)

	data := pm
	res := CodeConversionHeader{
		ExchangedInvoiceDocumentIdentifier: data.ExchangedInvoiceDocumentIdentifier,
		InvoiceDocument:                    data.InvoiceDocument,
		BillToParty:                        data.BillToParty,
		BillFromParty:                      data.BillFromParty,
		Payer:                              data.Payer,
		Payee:                              data.Payee,
		PaymentTerms:                       data.PaymentTerms,
		PaymentMethod:                      data.PaymentMethod,
	}

	return &res, nil
}

//func (psdc *SDC) ConvertToCodeConversionPartner(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionPartner, error) {
//	var err error
//
//	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
//	for _, v := range *dataQueryGets {
//		dataMap[v.LabelConvertTo] = v
//	}
//
//	pm := &requests.CodeConversionPartner{}
//
//	pm.PartnerFunction = dataMap["PartnerFunction"].CodeConvertTo
//	pm.BusinessPartner, err = ParseInt(dataMap["BusinessPartner"].CodeConvertTo)
//	if err != nil {
//		return nil, err
//	}
//
//	data := pm
//	res := CodeConversionPartner{
//		PartnerFunction: data.PartnerFunction,
//		BusinessPartner: data.BusinessPartner,
//	}
//
//	return &res, nil
//}

func (psdc *SDC) ConvertToCodeConversionItem(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItem, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItem{}

	pm.InvoiceDocumentItemIdentifier = dataMap["InvoiceDocumentItem"].CodeConvertFrom
	pm.InvoiceDocumentItem, err = ParseInt(dataMap["InvoiceDocumentItem"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.InvoiceDocumentItemCategory = GetStringPtr(dataMap["InvoiceDocumentItemCategory"].CodeConvertTo)
	pm.Buyer, err = ParseIntPtr(GetStringPtr(dataMap["Buyer"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.Seller, err = ParseIntPtr(GetStringPtr(dataMap["Seller"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.DeliverToParty, err = ParseIntPtr(GetStringPtr(dataMap["DeliverToParty"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.TransactionTaxClassification = GetStringPtr(dataMap["TransactionTaxClassification"].CodeConvertTo)
	pm.Project = GetStringPtr(dataMap["Project"].CodeConvertTo)
	pm.OrderID, err = ParseIntPtr(GetStringPtr(dataMap["OrderID"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.OrderItem, err = ParseIntPtr(GetStringPtr(dataMap["OrderItem"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.DeliveryDocument, err = ParseIntPtr(GetStringPtr(dataMap["DeliveryDocument"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.DeliveryDocumentItem, err = ParseIntPtr(GetStringPtr(dataMap["DeliveryDocumentItem"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItem{
		InvoiceDocumentItemIdentifier: data.InvoiceDocumentItemIdentifier,
		InvoiceDocumentItem:           data.InvoiceDocumentItem,
		InvoiceDocumentItemCategory:   data.InvoiceDocumentItemCategory,
		Buyer:                         data.Buyer,
		Seller:                        data.Seller,
		DeliverToParty:                data.DeliverToParty,
		TransactionTaxClassification:  data.TransactionTaxClassification,
		Project:                       data.Project,
		OrderID:                       data.OrderID,
		OrderItem:                     data.OrderItem,
		DeliveryDocument:              data.DeliveryDocument,
		DeliveryDocumentItem:          data.DeliveryDocumentItem,
		OriginDocument:                data.OrderID,
		OriginDocumentItem:            data.OrderItem,
		ReferenceDocument:             data.OrderID,
		ReferenceDocumentItem:         data.OrderItem,
	}

	return &res, nil
}

// func (psdc *SDC) ConvertToCodeConversionItemPricingElement(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItemPricingElement, error) {
// 	var err error
//
// 	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
// 	for _, v := range *dataQueryGets {
// 		dataMap[v.LabelConvertTo] = v
// 	}
//
// 	pm := &requests.CodeConversionItemPricingElement{}
//
// 	pm.PricingProcedureCounter, err = ParseInt(dataMap["PricingProcedureCounter"].CodeConvertTo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	pm.ConditionRecord, err = ParseIntPtr(GetStringPtr(dataMap["ConditionRecord"].CodeConvertTo))
// 	if err != nil {
// 		return nil, err
// 	}
// 	pm.ConditionSequentialNumber, err = ParseIntPtr(GetStringPtr(dataMap["ConditionSequentialNumber"].CodeConvertTo))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	data := pm
// 	res := CodeConversionItemPricingElement{
// 		PricingProcedureCounter:   data.PricingProcedureCounter,
// 		ConditionRecord:           data.ConditionRecord,
// 		ConditionSequentialNumber: data.ConditionSequentialNumber,
// 	}
//
// 	return &res, nil
// }

func (psdc *SDC) ConvertToConversionData() *[]ConversionData {
	var res []ConversionData

	for _, v := range *psdc.CodeConversionItem {
		pm := &requests.ConversionData{}

		pm.ExchangedInvoiceDocumentIdentifier = psdc.CodeConversionHeader.ExchangedInvoiceDocumentIdentifier
		pm.InvoiceDocument = psdc.CodeConversionHeader.InvoiceDocument
		pm.InvoiceDocumentItemIdentifier = v.InvoiceDocumentItemIdentifier
		pm.InvoiceDocumentItem = v.InvoiceDocumentItem

		data := pm
		res = append(res, ConversionData{
			ExchangedInvoiceDocumentIdentifier: data.ExchangedInvoiceDocumentIdentifier,
			InvoiceDocument:                    data.InvoiceDocument,
			InvoiceDocumentItemIdentifier:      data.InvoiceDocumentItemIdentifier,
			InvoiceDocumentItem:                data.InvoiceDocumentItem,
		})
	}

	return &res
}

// 個別処理
