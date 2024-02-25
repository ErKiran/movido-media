package billing

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type Contract struct {
	Email             string    `json:"email"`
	ProductCode       string    `json:"product_code"`
	ContractID        string    `json:"contract_id"`
	ContractStartDate time.Time `json:"contract_start_date"`
}

func (br billingRepo) SearchContracts(ctx context.Context) ([]Contract, error) {
	var contracts []Contract

	if err := br.db.WithContext(ctx).Raw(`select email, c2.contract_id, p.product_code, start_date as contract_start_date, 
	duration, duration_units, price, currency, 
	billing_frequency, billing_frequency_units
	from customers c 
	inner join contracts c2 on c.customer_id = c2.customer_id 
	inner join products p on p.product_code = c2.product_code`).
		Scan(&contracts).Scan(&contracts).Error; err != nil {
		log.Error().Msgf("unable to fetch contracts %s", err)
		return nil, err
	}
	return contracts, nil
}
