package billing

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (bc billingController) SearchCanditates(ctx context.Context) ([]string, error) {
	invoices, err := bc.billingRepo.SearchContracts(ctx)
	if err != nil {
		log.Error().Msgf("unable to fetch contracts %s", err)
		return nil, err
	}

	var canidates []string

	for _, invoice := range invoices {
		canidates = append(canidates, invoice.ContractID)
	}

	return canidates, err
}
