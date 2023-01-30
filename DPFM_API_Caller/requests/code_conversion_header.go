package requests

type CodeConversionHeader struct {
	ExchangedInvoiceDocumentIdentifier string  `json:"ExchangedInvoiceDocumentIdentifier "`
	InvoiceDocument                    int     `json:"InvoiceDocument"`
	BillToParty                        *int    `json:"BillToParty"`
	BillFromParty                      *int    `json:"BillFromParty"`
	Payer                              *int    `json:"Payer"`
	Payee                              *int    `json:"Payee"`
	PaymentTerms                       *string `json:"PaymentTerms"`
	PaymentMethod                      *string `json:"PaymentMethod"`
}
