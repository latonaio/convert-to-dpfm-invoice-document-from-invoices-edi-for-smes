package convert_complementer

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-dpfm-invoice-document-from-invoices-edi-for-smes/DPFM_API_Processing_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ConvertComplementer struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewConvertComplementer(ctx context.Context, db *database.Mysql, l *logger.Logger) *ConvertComplementer {
	return &ConvertComplementer{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (c *ConvertComplementer) CreateSdc(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(7)

	// Header
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-0. データ連携基盤Invoice Document HeaderとInvoices EDI For SMEsとの項目マッピング変換
		psdc.MappingHeader, err = c.ComplementMappingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Item
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-1. データ連携基盤Invoice Document ItemとInvoices EDI For SMEsとの項目マッピング変換
		psdc.MappingItem, err = c.ComplementMappingItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// <1-1. 番号変換＞
		psdc.CodeConversionHeader, err = c.CodeConversionHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// Item
		// <5-1. 番号変換＞
		psdc.CodeConversionItem, err = c.CodeConversionItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		psdc.ConversionData = c.ConversionData(sdc, psdc)
	}(&wg)

	// ItemPricingElement
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 8-1. データ連携基盤Invoice Document Item Pricing ElementとInvoices EDI For SMEsとの項目マッピング変換
		psdc.MappingItemPricingElement, err = c.ComplementMappingItemPricingElement(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	//	go func(wg *sync.WaitGroup) {
	//		defer wg.Done()
	//		// <8-1. 番号変換＞
	//		psdc.CodeConversionItemPricingElement, err = c.CodeConversionItemPricingElement(sdc, psdc)
	//		if e != nil {
	//			err = e
	//			return
	//		}
	//	}(&wg)

	// Partner
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-1. データ連携基盤Invoice Document PartnerとInvoices EDI For SMEs との項目マッピング変換
		psdc.MappingPartner, err = c.ComplementMappingPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	//	go func(wg *sync.WaitGroup) {
	//		defer wg.Done()
	//		// <2-1. 番号変換＞
	//		psdc.CodeConversionPartner, err = c.CodeConversionPartner(sdc, psdc)
	//		if e != nil {
	//			err = e
	//			return
	//		}
	//	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	c.l.Info(psdc)
	osdc, err = c.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
