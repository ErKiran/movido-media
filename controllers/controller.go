package controllers

import (
	"movido-media/controllers/billing"
	"movido-media/controllers/email"
	"movido-media/controllers/pdf"

	"gorm.io/gorm"
)

type controller struct {
	BillingController billing.BillingController
	PDFController     pdf.PDFController
	EmailController   email.EmailController
}

func NewController(db *gorm.DB) *controller {
	return &controller{
		BillingController: billing.NewBillingController(db),
		PDFController:     pdf.NewPDFController(),
		EmailController:   email.NewEmailController(),
	}
}
