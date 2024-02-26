package controllers

import (
	"movido-media/controllers/billing"
	"movido-media/controllers/pdf"

	"gorm.io/gorm"
)

type controller struct {
	BillingController billing.BillingController
	PDFController     pdf.PDFController
}

func NewController(db *gorm.DB) *controller {
	return &controller{
		BillingController: billing.NewBillingController(db),
		PDFController:     pdf.NewPDFController(),
	}
}
