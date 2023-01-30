package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Headerの処理
func (c *ConvertComplementer) ComplementMappingHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.MappingHeader, error) {
	res, err := psdc.ConvertToMappingHeader(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.CodeConversionHeader, error) {
	var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
	var args []interface{}

	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ExchangedInvoiceDocumentIdentifier", "InvoiceDocument", sdc.InvoicesEDIForSMEsHeader.ExchangedInvoiceDocumentIdentifier))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeBillToPartyRegisteredIdentifier", "BillToParty", sdc.InvoicesEDIForSMEsHeader.TradeBillToPartyRegisteredIdentifier))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeBillFromPartyIdentifier", "BillFromParty", sdc.InvoicesEDIForSMEsHeader.TradeBillFromPartyIdentifier))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeBillToPartyRegisteredIdentifier", "Payer", sdc.InvoicesEDIForSMEsHeader.TradeBillToPartyRegisteredIdentifier))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeBillFromPartyIdentifier", "Payer", sdc.InvoicesEDIForSMEsHeader.TradeBillFromPartyIdentifier))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradePaymentTermsTypeCode", "PaymentTerms", sdc.InvoicesEDIForSMEsHeader.TradePaymentTermsTypeCode))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "TradeSettlementPaymentMeansTypeCode", "PaymentMethod", sdc.InvoicesEDIForSMEsHeader.TradeSettlementPaymentMeansTypeCode))

	repeat := strings.Repeat("(?,?,?,?,?,?,?),", len(dataKey)-1) + "(?,?,?,?,?,?,?)"
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

	data, err := psdc.ConvertToCodeConversionHeader(dataQueryGets)
	if err != nil {
		return nil, err
	}

	return data, nil
}
