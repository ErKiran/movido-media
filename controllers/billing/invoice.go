package billing

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

var frequencyMap = map[string]int{
	"MONTHS": 1,
	"YEAR":   12,
}

type CandidateData struct {
	ContractID         string
	BillingFrequency   int
	CurrentBillingDate time.Time
}

func (bc billingController) SearchCanditates(ctx context.Context) ([]CandidateData, error) {
	invoices, err := bc.billingRepo.SearchContracts(ctx)
	if err != nil {
		log.Error().Msgf("unable to fetch contracts %s", err)
		return nil, err
	}

	var canidates []CandidateData

	today := time.Now()

	for _, invoice := range invoices {
		startDate := invoice.StartDate
		billingFrequency := invoice.BillingFrequency
		billingDuration := invoice.BillingDuration
		billingDurationUnit := invoice.BillingDurationUnits

		// to calculate the expiry of the contract
		contractDuration := billingDuration * frequencyMap[billingDurationUnit]

		endDate := startDate.AddDate(0, contractDuration, 0)

		if today.After(endDate) {
			log.Info().Msgf("This contract %s has expired skipping it", invoice.ContractID)
			// expired contract skip it
			continue
		}

		// using 30 days month for simplicity
		contractRenewedFor := MonthsDiff(startDate, today)

		billingFrequencyUnit := invoice.BillingFrequencyUnits
		bfreq := billingFrequency * frequencyMap[billingFrequencyUnit]

		// for simplicity using billing date as initial subscription date instead of using Prorated or Calendar Billing
		nextBillingDate := startDate.AddDate(0, bfreq+int(contractRenewedFor), 0)

		if TimeToString(today) != TimeToString(nextBillingDate) {
			continue
		}

		canidates = append(canidates, CandidateData{
			ContractID:         invoice.ContractID,
			BillingFrequency:   bfreq,
			CurrentBillingDate: today,
		})
	}

	return canidates, err
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02")
}

func MonthsDiff(t1, t2 time.Time) int {
	diffYears := t2.Year() - t1.Year()
	diffMonths := int(t2.Month()) - int(t1.Month())

	if t2.Day() <= t1.Day() {
		diffMonths--
	}

	// Calculate the total difference in months
	totalMonths := diffYears*12 + diffMonths

	// Take the absolute value
	if totalMonths < 0 {
		return -totalMonths
	}
	return totalMonths
}
