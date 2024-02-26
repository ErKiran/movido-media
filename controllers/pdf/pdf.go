package pdf

import "movido-media/controllers/billing"

type invoicePDF struct{}

type PDFController interface {
	Generate(billing.ContractDetail)
}

func NewPDFController() PDFController {
	return invoicePDF{}
}
