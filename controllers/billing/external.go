package billing

import (
	"context"
	"net/http"
	"time"
)

type BillingResponse struct {
	ID        string    `json:"id"`
	ToAddress string    `json:"to_address"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	Timestamp time.Time `json:"timestamp"`
}

type BillingRequest struct {
	RecipientAddress string  `json:"recipient_address"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
}

func (bc billingController) ExternalBillingCall(ctx context.Context) (*BillingResponse, error) {
	br := BillingRequest{
		RecipientAddress: "mock-address",
		Amount:           700,
		Currency:         "EUR",
	}
	req, err := bc.client.NewRequest(http.MethodPost, "api/v1/transactions", br)
	if err != nil {
		return nil, err
	}

	var res *BillingResponse
	if _, err := bc.client.Do(ctx, req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
