package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Processing_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*Header, error) {
	mappingHeader := psdc.MappingHeader
	codeConversionHeader := psdc.CodeConversionHeader

	header := Header{}

	data, err := json.Marshal(mappingHeader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &header)
	if err != nil {
		return nil, err
	}

	header.InvoiceDocument = codeConversionHeader.InvoiceDocument
	header.BillToParty = codeConversionHeader.BillToParty
	header.BillFromParty = codeConversionHeader.BillFromParty
	header.Payer = codeConversionHeader.Payer
	header.Payee = codeConversionHeader.Payee
	header.PaymentTerms = codeConversionHeader.PaymentTerms
	header.PaymentMethod = codeConversionHeader.PaymentMethod

	return &header, nil
}

func ConvertToItem(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Item, error) {
	var items []Item
	mappingItem := psdc.MappingItem
	codeConversionItem := psdc.CodeConversionItem
	conversionData := psdc.ConversionData

	for i := range *mappingItem {
		item := Item{}

		data, err := json.Marshal((*mappingItem)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &item)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.ExchangedInvoiceDocumentIdentifier == (*mappingItem)[i].ExchangedInvoiceDocumentIdentifier {
				item.InvoiceDocument = v.InvoiceDocument
				break
			}
		}
		item.InvoiceDocumentItem = (*codeConversionItem)[i].InvoiceDocumentItem
		item.InvoiceDocumentItemCategory = (*codeConversionItem)[i].InvoiceDocumentItemCategory
		item.Buyer = (*codeConversionItem)[i].Buyer
		item.Seller = (*codeConversionItem)[i].Seller
		item.DeliverToParty = (*codeConversionItem)[i].DeliverToParty
		item.TransactionTaxClassification = (*codeConversionItem)[i].TransactionTaxClassification
		item.Project = (*codeConversionItem)[i].Project
		item.OrderID = (*codeConversionItem)[i].OrderID
		item.OrderItem = (*codeConversionItem)[i].OrderItem
		item.DeliveryDocument = (*codeConversionItem)[i].DeliveryDocument
		item.DeliveryDocumentItem = (*codeConversionItem)[i].DeliveryDocumentItem
		item.OriginDocument = (*codeConversionItem)[i].OriginDocument
		item.OriginDocumentItem = (*codeConversionItem)[i].OriginDocumentItem
		item.ReferenceDocument = (*codeConversionItem)[i].ReferenceDocument
		item.ReferenceDocumentItem = (*codeConversionItem)[i].ReferenceDocumentItem

		items = append(items, item)
	}

	return &items, nil
}

func ConvertToItemPricingElement(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]ItemPricingElement, error) {
	var itemPricingElements []ItemPricingElement
	mappingItemPricingElement := psdc.MappingItemPricingElement
	conversionData := psdc.ConversionData

	for i := range *mappingItemPricingElement {
		itemPricingElement := ItemPricingElement{}

		data, err := json.Marshal((*mappingItemPricingElement)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &itemPricingElement)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.ExchangedInvoiceDocumentIdentifier == (*mappingItemPricingElement)[i].ExchangedInvoiceDocumentIdentifier && v.InvoiceDocumentItemIdentifier == (*mappingItemPricingElement)[i].InvoiceDocumentItemIdentifier {
				itemPricingElement.InvoiceDocument = v.InvoiceDocument
				itemPricingElement.InvoiceDocumentItem = v.InvoiceDocumentItem
				break
			}
		}

		itemPricingElements = append(itemPricingElements, itemPricingElement)
	}

	return &itemPricingElements, nil
}

func ConvertToPartner(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Partner, error) {
	var partners []Partner
	mappingPartner := psdc.MappingPartner
	conversionData := psdc.ConversionData

	for i := range *mappingPartner {
		partner := Partner{}

		data, err := json.Marshal((*mappingPartner)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &partner)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.ExchangedInvoiceDocumentIdentifier == (*mappingPartner)[i].ExchangedInvoiceDocumentIdentifier {
				partner.InvoiceDocument = v.InvoiceDocument
				break
			}
		}

		partners = append(partners, partner)
	}

	return &partners, nil
}
