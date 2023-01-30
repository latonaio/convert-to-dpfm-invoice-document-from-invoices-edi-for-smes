package requests

type ConversionData struct {
	ExchangedInvoiceDocumentIdentifier string `json:"ExchangedInvoiceDocumentIdentifier "`
	InvoiceDocument                    int    `json:"InvoiceDocument"`
	InvoiceDocumentItemIdentifier      string `json:"InvoiceDocumentItemIdentifier "`
	InvoiceDocumentItem                int    `json:"InvoiceDocumentItem"`
}
