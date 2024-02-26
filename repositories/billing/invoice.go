package billing

import (
	"context"
	"time"
)

type Contract struct {
	ContractID            string    `gorm:"column:contract_id"`
	StartDate             time.Time `gorm:"column:contract_start_date"`
	BillingFrequency      int       `gorm:"column:billing_frequency"`
	BillingDuration       int       `gorm:"column:duration"`
	BillingFrequencyUnits string    `gorm:"column:billing_frequency_units"`
	BillingDurationUnits  string    `gorm:"column:duration_units"`
}

type ContractDetails struct {
	ContractID   string  `gorm:"column:contract_id"`
	CustomerName string  `gorm:"column:full_name"`
	Email        string  `gorm:"column:email"`
	Address      string  `gorm:"column:address"`
	ProductCode  string  `gorm:"column:product_code"`
	ProductName  string  `gorm:"column:product_name"`
	Price        float64 `gorm:"column:price"`
	Currency     string  `gorm:"column:currency"`
}

func (br billingRepo) GetContractsDetail(ctx context.Context, ids []string) ([]ContractDetails, error) {
	var cd []ContractDetails

	if err := br.db.WithContext(ctx).Raw(`
	select  full_name, email, address,c2.contract_id,
	p.product_code, product_name, price, currency
	from customers c 
	inner join contracts c2 on c.customer_id = c2.customer_id
	inner join products p on p.product_code = c2.product_code
	where contract_id in (?)
	`, ids).Scan(&cd).Error; err != nil {
		return nil, err
	}
	return cd, nil
}

func (br billingRepo) SearchContracts(ctx context.Context) ([]Contract, error) {
	var c []Contract

	if err := br.db.WithContext(ctx).Raw(`
	select  c2.contract_id, 
	start_date as contract_start_date, 
	duration, duration_units, 
	billing_frequency, billing_frequency_units
	from customers c 
	inner join contracts c2 on c.customer_id = c2.customer_id
	`).Scan(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
