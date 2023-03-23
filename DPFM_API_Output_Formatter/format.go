package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Processing_Formatter"
)

func OutputFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_formatter.ProcessingFormatterSDC,
	osdc *Output,
) error {
	header := ConvertToHeader(*sdc, *psdc)
	item := ConvertToItem(*sdc, *psdc)
	address := ConvertToAddress(*sdc, *psdc)
	partner := ConvertToPartner(*sdc, *psdc)

	osdc.DataConcatenation = DataConcatenation{
		Header:  header,
		Item:    item,
		Address: address,
		Partner: partner,
	}

	osdc.ServiceLabel = "FUNCTION_INVOICE_DOCUMENT_DATA_CONCATENATION"
	osdc.APISchema = "DPFMDataConcatenation"
	osdc.APIProcessingResult = getBoolPtr(true)

	return nil
}

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) *Header {
	dataProcessingHeader := psdc.Header
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

	header := &Header{
		InvoiceDocument:                   *dataConversionProcessingHeader.ConvertedInvoiceDocument,
		CreationDate:                      dataProcessingHeader.CreationDate,
		CreationTime:                      dataProcessingHeader.CreationTime,
		LastChangeDate:                    dataProcessingHeader.LastChangeDate,
		LastChangeTime:                    dataProcessingHeader.LastChangeTime,
		BillToParty:                       dataConversionProcessingHeader.ConvertedBillToParty,
		BillFromParty:                     dataConversionProcessingHeader.ConvertedBillFromParty,
		Payer:                             dataConversionProcessingHeader.ConvertedPayer,
		Payee:                             dataConversionProcessingHeader.ConvertedPayee,
		InvoiceDocumentDate:               dataProcessingHeader.InvoiceDocumentDate,
		InvoicePeriodStartDate:            dataProcessingHeader.InvoicePeriodStartDate,
		InvoicePeriodEndDate:              dataProcessingHeader.InvoicePeriodEndDate,
		AccountingPostingDate:             dataProcessingHeader.AccountingPostingDate,
		TotalNetAmount:                    dataProcessingHeader.TotalNetAmount,
		TotalTaxAmount:                    dataProcessingHeader.TotalTaxAmount,
		TotalGrossAmount:                  dataProcessingHeader.TotalGrossAmount,
		TransactionCurrency:               dataProcessingHeader.TransactionCurrency,
		PaymentTerms:                      dataProcessingHeader.PaymentTerms,
		PaymentDueDate:                    dataProcessingHeader.PaymentDueDate,
		PaymentMethod:                     dataConversionProcessingHeader.ConvertedPaymentMethod,
		DocumentHeaderText:                dataProcessingHeader.DocumentHeaderText,
		HeaderIsCleared:                   dataProcessingHeader.HeaderIsCleared,
		HeaderPaymentBlockStatus:          dataProcessingHeader.HeaderPaymentBlockStatus,
		HeaderPaymentRequisitionIsCreated: dataProcessingHeader.HeaderPaymentRequisitionIsCreated,
		IsCancelled:                       dataProcessingHeader.IsCancelled,
	}

	return header
}

func ConvertToItem(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Item {
	dataProcessingItem := psdc.Item
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	dataConversionProcessingItem := psdc.ConversionProcessingItem

	items := make([]*Item, 0)
	for i := range dataProcessingItem {
		item := &Item{
			InvoiceDocument:                 *dataConversionProcessingHeader.ConvertedInvoiceDocument,
			InvoiceDocumentItem:             *dataConversionProcessingItem[i].ConvertedInvoiceDocumentItem,
			InvoiceDocumentItemText:         dataProcessingItem[i].InvoiceDocumentItemText,
			Product:                         dataConversionProcessingItem[i].ConvertedProduct,
			CreationDate:                    dataProcessingItem[i].CreationDate,
			CreationTime:                    dataProcessingItem[i].CreationTime,
			LastChangeDate:                  dataProcessingItem[i].LastChangeDate,
			LastChangeTime:                  dataProcessingItem[i].LastChangeTime,
			Buyer:                           dataConversionProcessingItem[i].ConvertedBuyer,
			Seller:                          dataConversionProcessingItem[i].ConvertedSeller,
			DeliverToParty:                  dataConversionProcessingItem[i].ConvertedDeliverToParty,
			InvoiceQuantity:                 dataProcessingItem[i].InvoiceQuantity,
			InvoiceQuantityUnit:             dataProcessingItem[i].InvoiceQuantityUnit,
			InvoiceQuantityInBaseUnit:       dataProcessingItem[i].InvoiceQuantityInBaseUnit,
			BaseUnit:                        dataProcessingItem[i].BaseUnit,
			NetAmount:                       dataProcessingItem[i].NetAmount,
			GrossAmount:                     dataProcessingItem[i].GrossAmount,
			TransactionCurrency:             dataProcessingItem[i].TransactionCurrency,
			Project:                         dataConversionProcessingItem[i].ConvertedProject,
			OrderID:                         dataConversionProcessingItem[i].ConvertedOrderID,
			OrderItem:                       dataConversionProcessingItem[i].ConvertedOrderItem,
			InvoicePeriodStartDate:          dataProcessingItem[i].InvoicePeriodStartDate,
			InvoicePeriodEndDate:            dataProcessingItem[i].InvoicePeriodEndDate,
			DeliveryDocument:                dataConversionProcessingItem[i].ConvertedDeliveryDocument,
			DeliveryDocumentItem:            dataConversionProcessingItem[i].ConvertedDeliveryDocumentItem,
			OriginDocument:                  dataConversionProcessingItem[i].ConvertedOriginDocument,
			OriginDocumentItem:              dataConversionProcessingItem[i].ConvertedOriginDocumentItem,
			ReferenceDocument:               dataConversionProcessingItem[i].ConvertedReferenceDocument,
			ReferenceDocumentItem:           dataConversionProcessingItem[i].ConvertedReferenceDocumentItem,
			ItemPaymentRequisitionIsCreated: dataProcessingItem[i].ItemPaymentRequisitionIsCreated,
			ItemIsCleared:                   dataProcessingItem[i].ItemIsCleared,
			ItemPaymentBlockStatus:          dataProcessingItem[i].ItemPaymentBlockStatus,
			IsCancelled:                     dataProcessingItem[i].IsCancelled,
		}

		items = append(items, item)
	}

	return items
}

func ConvertToItemPricingElement(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*ItemPricingElement {
	dataProcessingItemPricingElement := psdc.ItemPricingElement
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	dataConversionProcessingItem := psdc.ConversionProcessingItem

	dataConversionProcessingItemMap := make(map[string]*dpfm_api_processing_formatter.ConversionProcessingItem, len(dataConversionProcessingItem))
	for _, v := range dataConversionProcessingItem {
		dataConversionProcessingItemMap[*v.ConvertingInvoiceDocumentItem] = v
	}

	itemPricingElements := make([]*ItemPricingElement, 0)
	for i, v := range dataProcessingItemPricingElement {
		if _, ok := dataConversionProcessingItemMap[v.ConvertingInvoiceDocumentItem]; !ok {
			continue
		}

		itemPricingElements = append(itemPricingElements, &ItemPricingElement{
			InvoiceDocument:            *dataConversionProcessingHeader.ConvertedInvoiceDocument,
			InvoiceDocumentItem:        *dataConversionProcessingItemMap[v.ConvertingInvoiceDocumentItem].ConvertedInvoiceDocumentItem,
			ConditionRateValue:         dataProcessingItemPricingElement[i].ConditionRateValue,
			ConditionCurrency:          dataProcessingItemPricingElement[i].ConditionCurrency,
			ConditionQuantity:          dataProcessingItemPricingElement[i].ConditionQuantity,
			ConditionQuantityUnit:      dataProcessingItemPricingElement[i].ConditionQuantityUnit,
			ConditionAmount:            dataProcessingItemPricingElement[i].ConditionAmount,
			TransactionCurrency:        dataProcessingItemPricingElement[i].TransactionCurrency,
			ConditionIsManuallyChanged: dataProcessingItemPricingElement[i].ConditionIsManuallyChanged,
		})
	}

	return itemPricingElements
}

func ConvertToAddress(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Address {
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	data := psdc.Address

	addresses := make([]*Address, 0)
	for _, data := range data {
		addresses = append(addresses, &Address{
			InvoiceDocument: *dataConversionProcessingHeader.ConvertedInvoiceDocument,
			PostalCode:      data.PostalCode,
		})
	}

	return addresses
}

func ConvertToPartner(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Partner {
	dataProcessingPartner := psdc.Partner
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

	partners := make([]*Partner, 0)
	for range dataProcessingPartner {
		partners = append(partners, &Partner{
			InvoiceDocument: *dataConversionProcessingHeader.ConvertedInvoiceDocument,
		})
	}

	return partners
}

func getBoolPtr(b bool) *bool {
	return &b
}
