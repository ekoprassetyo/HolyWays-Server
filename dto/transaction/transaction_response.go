package Transactionsdto

import "time"

type TransactionResponse struct {
	ID           int       `json:"id"`
	DonateAmount string    `json:"donateAmount" validate:"required"`
	Status       string    `json:"status" validate:"required"`
	CreatedAt    time.Time `json:"startdate"`
}
