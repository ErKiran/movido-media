package billing

import (
	"context"

	"gorm.io/gorm"
)

type BillingRepo interface {
	SearchContracts(ctx context.Context) ([]Contract, error)
}

type billingRepo struct {
	db *gorm.DB
}

func NewBillingRepo(db *gorm.DB) BillingRepo {
	return billingRepo{
		db: db,
	}
}
