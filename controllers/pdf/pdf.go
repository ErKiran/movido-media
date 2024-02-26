package pdf

import "movido-media/controllers/billing"

type invoicePDF struct{}

type PDFController interface {
	Generate(billing.ContractDetail) (string, error)
}

func NewPDFController() PDFController {
	return invoicePDF{}
}
