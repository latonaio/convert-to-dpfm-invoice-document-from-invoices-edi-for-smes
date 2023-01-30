package requests

type CodeConversionItem struct {
	InvoiceDocumentItemIdentifier string  `json:"InvoiceDocumentItemIdentifier "`
	InvoiceDocumentItem           int     `json:"InvoiceDocumentItem"`
	InvoiceDocumentItemCategory   *string `json:"InvoiceDocumentItemCategory"`
	Buyer                         *int    `json:"Buyer"`
	Seller                        *int    `json:"Seller"`
	DeliverToParty                *int    `json:"DeliverToParty"`
	TransactionTaxClassification  *string `json:"TransactionTaxClassification"`
	Project                       *string `json:"Project"`
	OrderID                       *int    `json:"OrderID"`
	OrderItem                     *int    `json:"OrderItem"`
	DeliveryDocument              *int    `json:"DeliveryDocument"`
	DeliveryDocumentItem          *int    `json:"DeliveryDocumentItem"`
	OriginDocument                *int    `json:"OriginDocument"`
	OriginDocumentItem            *int    `json:"OriginDocumentItem"`
	ReferenceDocument             *int    `json:"ReferenceDocument"`
	ReferenceDocumentItem         *int    `json:"ReferenceDocumentItem"`
}
