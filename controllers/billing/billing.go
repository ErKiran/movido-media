package billing

import (
	"context"

	"movido-media/repositories/billing"
	repo "movido-media/repositories/billing"

	"gorm.io/gorm"
)

type billingController struct {
	billingRepo repo.BillingRepo
}

type BillingController interface {
	SearchCanditates(ctx context.Context) ([]CandidateData, error)
	Details(ctx context.Context, data []CandidateData) ([]billing.ContractDetails, error)
}

func NewBillingController(db *gorm.DB) BillingController {
	return billingController{
		billingRepo: repo.NewBillingRepo(db),
	}
}
