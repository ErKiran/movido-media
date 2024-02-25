package billing

import (
	"context"

	repo "movido-media/repositories/billing"

	"gorm.io/gorm"
)

type billingController struct {
	billingRepo repo.BillingRepo
}

type BillingController interface {
	SearchCanditates(ctx context.Context) ([]string, error)
}

func NewBillingController(db *gorm.DB) BillingController {
	return billingController{
		billingRepo: repo.NewBillingRepo(db),
	}
}
