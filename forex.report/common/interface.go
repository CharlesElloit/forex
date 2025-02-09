package common

import "github.com/google/uuid"

type GetRateInput struct {
	From       string    `json:"from"`
	To         string    `json:"to"`
	CurrencyID uuid.UUID `json:"currency_id"`
}
