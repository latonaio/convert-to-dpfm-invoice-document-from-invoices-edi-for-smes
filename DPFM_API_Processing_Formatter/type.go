package dpfm_api_processing_formatter

type ProcessingFormatterSDC struct {
	Header                     *Header                     `json:"Header"`
	ConversionProcessingHeader *ConversionProcessingHeader `json:"ConversionProcessingHeader"`
	Item                       []*Item                     `json:"Item"`
	ConversionProcessingItem   []*ConversionProcessingItem `json:"ConversionProcessingItem"`
	ItemPricingElement         []*ItemPricingElement       `json:"ItemPricingElement"`
	Address                    []*Address                  `json:"Address"`
	Partner                    []*Partner                  `json:"Partner"`
}

type ConversionProcessingKey struct {
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type ConversionProcessingCommonQueryGets struct {
	CodeConversionID      int      `json:"CodeConversionID"`
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	CodeConvertToInt      *int     `json:"CodeConvertToInt"`
	CodeConvertToFloat    *float32 `json:"CodeConvertToFloat"`
	CodeConvertToString   *string  `json:"CodeConvertToString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type Header struct {
	ConvertingInvoiceDocument         string   `json:"ConvertingInvoiceDocument"`
	CreationDate                      *string  `json:"CreationDate"`
	CreationTime                      *string  `json:"CreationTime"`
	LastChangeDate                    *string  `json:"LastChangeDate"`
	LastChangeTime                    *string  `json:"LastChangeTime"`
	ConvertingBillToParty             *string  `json:"ConvertingBillToParty"`
	ConvertingBillFromParty           *string  `json:"ConvertingBillFromParty"`
	ConvertingPayer                   *string  `json:"ConvertingPayer"`
	ConvertingPayee                   *string  `json:"ConvertingPayee"`
	InvoiceDocumentDate               *string  `json:"InvoiceDocumentDate"`
	InvoicePeriodStartDate            *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              *string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string  `json:"AccountingPostingDate"`
	TotalNetAmount                    *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32 `json:"TotalGrossAmount"`
	TransactionCurrency               *string  `json:"TransactionCurrency"`
	PaymentTerms                      *string  `json:"PaymentTerms"`
	PaymentDueDate                    *string  `json:"PaymentDueDate"`
	ConvertingPaymentMethod           *string  `json:"ConvertingPaymentMethod"`
	DocumentHeaderText                *string  `json:"DocumentHeaderText"`
	HeaderIsCleared                   *bool    `json:"HeaderIsCleared"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
	IsCancelled                       *bool    `json:"IsCancelled"`
}

type ConversionProcessingHeader struct {
	ConvertingInvoiceDocument *string `json:"ConvertingInvoiceDocument"`
	ConvertedInvoiceDocument  *int    `json:"ConvertedInvoiceDocument"`
	ConvertingBillToParty     *string `json:"ConvertingBillToParty"`
	ConvertedBillToParty      *int    `json:"ConvertedBillToParty"`
	ConvertingBillFromParty   *string `json:"ConvertingBillFromParty"`
	ConvertedBillFromParty    *int    `json:"ConvertedBillFromParty"`
	ConvertingPayer           *string `json:"ConvertingPayer"`
	ConvertedPayer            *int    `json:"ConvertedPayer"`
	ConvertingPayee           *string `json:"ConvertingPayee"`
	ConvertedPayee            *int    `json:"ConvertedPayee"`
	ConvertingPaymentMethod   *string `json:"ConvertingPaymentMethod"`
	ConvertedPaymentMethod    *string `json:"ConvertedPaymentMethod"`
}

type Item struct {
	ConvertingInvoiceDocument              string   `json:"ConvertingInvoiceDocument"`
	ConvertingInvoiceDocumentItem          string   `json:"ConvertingInvoiceDocumentItem"`
	ConvertingInvoiceDocumentItemCategory  *string  `json:"ConvertingInvoiceDocumentItemCategory"`
	InvoiceDocumentItemText                *string  `json:"InvoiceDocumentItemText"`
	ConvertingProduct                      *string  `json:"ConvertingProduct"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
	ConvertingBuyer                        *string  `json:"ConvertingBuyer"`
	ConvertingSeller                       *string  `json:"ConvertingSeller"`
	ConvertingDeliverToParty               *string  `json:"ConvertingDeliverToParty"`
	InvoiceQuantity                        *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit                    *string  `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit              *float32 `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                               *string  `json:"BaseUnit"`
	NetAmount                              *float32 `json:"NetAmount"`
	GrossAmount                            *float32 `json:"GrossAmount"`
	TransactionCurrency                    *string  `json:"TransactionCurrency"`
	ConvertingTransactionTaxClassification *string  `json:"ConvertingTransactionTaxClassification"`
	ConvertingProject                      *string  `json:"ConvertingProject"`
	ConvertingOrderID                      *string  `json:"ConvertingOrderID"`
	ConvertingOrderItem                    *string  `json:"ConvertingOrderItem"`
	InvoicePeriodStartDate                 *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                   *string  `json:"InvoicePeriodEndDate"`
	ConvertingDeliveryDocument             *string  `json:"ConvertingDeliveryDocument"`
	ConvertingDeliveryDocumentItem         *string  `json:"ConvertingDeliveryDocumentItem"`
	ConvertingOriginDocument               *string  `json:"ConvertingOriginDocument"`
	ConvertingOriginDocumentItem           *string  `json:"ConvertingOriginDocumentItem"`
	ConvertingReferenceDocument            *string  `json:"ConvertingReferenceDocument"`
	ConvertingReferenceDocumentItem        *string  `json:"ConvertingReferenceDocumentItem"`
	ItemPaymentRequisitionIsCreated        *bool    `json:"ItemPaymentRequisitionIsCreated"`
	ItemIsCleared                          *bool    `json:"ItemIsCleared"`
	ItemPaymentBlockStatus                 *bool    `json:"ItemPaymentBlockStatus"`
	IsCancelled                            *bool    `json:"IsCancelled"`
}

type ConversionProcessingItem struct {
	ConvertingInvoiceDocumentItem          *string `json:"ConvertingInvoiceDocumentItem"`
	ConvertedInvoiceDocumentItem           *int    `json:"ConvertedInvoiceDocumentItem"`
	ConvertingProduct                      *string `json:"ConvertingProduct"`
	ConvertedProduct                       *string `json:"ConvertedProduct"`
	ConvertingBuyer                        *string `json:"ConvertingBuyer"`
	ConvertedBuyer                         *int    `json:"ConvertedBuyer"`
	ConvertingSeller                       *string `json:"ConvertingSeller"`
	ConvertedSeller                        *int    `json:"ConvertedSeller"`
	ConvertingDeliverToParty               *string `json:"ConvertingDeliverToParty"`
	ConvertedDeliverToParty                *int    `json:"ConvertedDeliverToParty"`
	ConvertingTransactionTaxClassification *string `json:"ConvertingTransactionTaxClassification"`
	ConvertedTransactionTaxClassification  *string `json:"ConvertedTransactionTaxClassification"`
	ConvertingProject                      *string `json:"ConvertingProject"`
	ConvertedProject                       *string `json:"ConvertedProject"`
	ConvertingOrderID                      *string `json:"ConvertingOrderID"`
	ConvertedOrderID                       *int    `json:"ConvertedOrderID"`
	ConvertingOrderItem                    *string `json:"ConvertingOrderItem"`
	ConvertedOrderItem                     *int    `json:"ConvertedOrderItem"`
	ConvertingDeliveryDocument             *string `json:"ConvertingDeliveryDocument"`
	ConvertedDeliveryDocument              *int    `json:"ConvertedDeliveryDocument"`
	ConvertingDeliveryDocumentItem         *string `json:"ConvertingDeliveryDocumentItem"`
	ConvertedDeliveryDocumentItem          *int    `json:"ConvertedDeliveryDocumentItem"`
	ConvertingOriginDocument               *string `json:"ConvertingOriginDocument"`
	ConvertedOriginDocument                *int    `json:"ConvertedOriginDocument"`
	ConvertingOriginDocumentItem           *string `json:"ConvertingOriginDocumentItem"`
	ConvertedOriginDocumentItem            *int    `json:"ConvertedOriginDocumentItem"`
	ConvertingReferenceDocument            *string `json:"ConvertingReferenceDocument"`
	ConvertedReferenceDocument             *int    `json:"ConvertedReferenceDocument"`
	ConvertingReferenceDocumentItem        *string `json:"ConvertingReferenceDocumentItem"`
	ConvertedReferenceDocumentItem         *int    `json:"ConvertedReferenceDocumentItem"`
}

type ItemPricingElement struct {
	ConvertingInvoiceDocument     string   `json:"ConvertingInvoiceDocument"`
	ConvertingInvoiceDocumentItem string   `json:"ConvertingInvoiceDocumentItem"`
	ConditionRateValue            *float32 `json:"ConditionRateValue"`
	ConditionCurrency             *string  `json:"ConditionCurrency"`
	ConditionQuantity             *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit         *string  `json:"ConditionQuantityUnit"`
	ConditionAmount               *float32 `json:"ConditionAmount"`
	TransactionCurrency           *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged    *bool    `json:"ConditionIsManuallyChanged"`
}

type Address struct {
	ConvertingInvoiceDocument string  `json:"ConvertingInvoiceDocument"`
	PostalCode                *string `json:"PostalCode"`
}

type Partner struct {
	ConvertingInvoiceDocument string `json:"ConvertingInvoiceDocument"`
}
