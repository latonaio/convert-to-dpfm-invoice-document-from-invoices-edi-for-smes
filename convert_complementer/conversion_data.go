package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Processing_Formatter"
)

func (c *ConvertComplementer) ConversionData(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) *[]dpfm_api_processing_formatter.ConversionData {
	data := psdc.ConvertToConversionData()

	return data
}
