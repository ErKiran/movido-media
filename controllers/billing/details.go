package billing

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

type ContractDetail struct {
	CustomerName     string
	Email            string
	Address          string
	ProductCode      string
	ProductName      string
	Price            float64
	Currency         string
	BillingFrequency int
	BillDate         time.Time
}

func (bc billingController) Details(ctx context.Context, data []CandidateData) ([]ContractDetail, error) {
	var contracts []string
	var contractDetail []ContractDetail

	contractMap := make(map[string]CandidateData)

	for _, d := range data {
		contracts = append(contracts, d.ContractID)
		contractMap[d.ContractID] = d
	}

	if len(contracts) == 0 {
		return nil, errors.New("no contracts to fetch details")
	}

	details, err := bc.billingRepo.GetContractsDetail(ctx, contracts)
	if err != nil {
		log.Error().Msgf("unable to get contracts detail %s", err)
		return nil, err
	}

	for _, det := range details {
		contractDetail = append(contractDetail, ContractDetail{
			CustomerName:     det.CustomerName,
			Email:            det.Email,
			Address:          det.Address,
			ProductCode:      det.ProductCode,
			ProductName:      det.ProductName,
			Price:            det.Price,
			Currency:         det.Currency,
			BillingFrequency: contractMap[det.ContractID].BillingFrequency,
			BillDate:         contractMap[det.ContractID].CurrentBillingDate,
		})
	}
	return contractDetail, nil
}
