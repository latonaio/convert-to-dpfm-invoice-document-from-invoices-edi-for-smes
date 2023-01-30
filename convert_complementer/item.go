package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Itemの処理
func (c *ConvertComplementer) ComplementMappingItem(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingItem, error) {
	res, err := psdc.ConvertToMappingItem(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionItem(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionItem, error) {
	var data []dpfm_api_processing_formatter.CodeConversionItem

	for _, item := range sdc.InvoicesEDIForSMEsHeader.InvoicesEDIForSMEsItem {
		var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
		var args []interface{}

		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "InvoiceDocumentItemIdentifier", "InvoiceDocumentItem", item.InvoiceDocumentItemIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "InvoiceDocumentItemCategoryCode", "InvoiceDocumentItemCategory", item.InvoiceDocumentItemCategoryCode))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeBuyerIdentifier", "Buyer", sdc.InvoicesEDIForSMEsHeader.TradeBuyerIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeSellerIdentifier", "Seller", sdc.InvoicesEDIForSMEsHeader.TradeSellerIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeShipToPartyIdentifier", "DeliverToParty", item.TradeShipToPartyIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ProjectIdentifier", "Project", sdc.InvoicesEDIForSMEsHeader.ProjectIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ReferencedOrdersDocumentIssuerAssignedIdentifier", "OrderID", item.ReferencedOrdersDocumentIssuerAssignedIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ReferencedSalesOrderDocumentItemLineIssuerAssignedIdentifier", "OrderItem", item.ReferencedSalesOrderDocumentItemLineIssuerAssignedIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ReferencedOrdersDocumentIssuerAssignedIdentifier", "DeliveryDocument", item.ReferencedOrdersDocumentIssuerAssignedIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ItemLineReferencedOrdersDocumentItemLineIssuerAssignedIdentifier", "DeliveryDocumentItem", item.ItemLineReferencedOrdersDocumentItemLineIssuerAssignedIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ReferencedOrdersDocumentIssuerAssignedIdentifier", "OriginDocument", item.ReferencedOrdersDocumentIssuerAssignedIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ReferencedSalesOrderDocumentItemLineIssuerAssignedIdentifier", "OriginDocumentItem", item.ReferencedSalesOrderDocumentItemLineIssuerAssignedIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ReferencedOrdersDocumentIssuerAssignedIdentifier", "ReferenceDocument", item.ReferencedOrdersDocumentIssuerAssignedIdentifier))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ReferencedSalesOrderDocumentItemLineIssuerAssignedIdentifier", "ReferenceDocumentItem", item.ReferencedSalesOrderDocumentItemLineIssuerAssignedIdentifier))

		repeat := strings.Repeat("(?,?,?,?,?,?),", len(dataKey)-1) + "(?,?,?,?,?,?)"
		for _, v := range dataKey {
			args = append(args, v.SystemConvertTo, v.SystemConvertFrom, v.LabelConvertTo, v.LabelConvertFrom, v.CodeConvertFrom, v.BusinessPartner)
		}

		rows, err := c.db.Query(
			`SELECT CodeConversionID, SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom,
			CodeConvertFrom, CodeConvertTo, BusinessPartner
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_code_conversion_code_conversion_data
			WHERE (SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom, CodeConvertFrom, BusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		dataQueryGets, err := psdc.ConvertToCodeConversionQueryGets(rows)
		if err != nil {
			return nil, err
		}

		datum, err := psdc.ConvertToCodeConversionItem(dataQueryGets)
		if err != nil {
			return nil, err
		}

		data = append(data, *datum)
	}

	return &data, nil
}
