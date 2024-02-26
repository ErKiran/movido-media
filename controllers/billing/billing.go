package billing

import (
	"context"
	"os"

	repo "movido-media/repositories/billing"
	"movido-media/utils"

	"gorm.io/gorm"
)

type billingController struct {
	billingRepo repo.BillingRepo
	client      *utils.Client
}

type BillingController interface {
	SearchCanditates(ctx context.Context) ([]CandidateData, error)
	Details(ctx context.Context, data []CandidateData) ([]ContractDetail, error)
	ExternalBillingCall(ctx context.Context) (*BillingResponse, error)
}

func NewBillingController(db *gorm.DB) BillingController {
	client := utils.NewClient(nil, os.Getenv("MOCK_API"), "")
	return billingController{
		billingRepo: repo.NewBillingRepo(db),
		client:      client,
	}
}
