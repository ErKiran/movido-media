package billing

import (
	"context"
	"errors"

	"movido-media/repositories/billing"

	"github.com/rs/zerolog/log"
)

func (bc billingController) Details(ctx context.Context, data []CandidateData) ([]billing.ContractDetails, error) {
	var contracts []string

	for _, d := range data {
		contracts = append(contracts, d.ContractID)
	}

	if len(contracts) == 0 {
		return nil, errors.New("no contracts to fetch details")
	}

	details, err := bc.billingRepo.GetContractsDetail(ctx, contracts)
	if err != nil {
		log.Error().Msgf("unable to get contracts detail %s", err)
		return nil, err
	}
	return details, nil
}
