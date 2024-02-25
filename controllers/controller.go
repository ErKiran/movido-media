package controllers

import (
	"movido-media/controllers/billing"

	"gorm.io/gorm"
)

type controller struct {
	BillingController billing.BillingController
}

func NewController(db *gorm.DB) *controller {
	return &controller{
		BillingController: billing.NewBillingController(db),
	}
}
