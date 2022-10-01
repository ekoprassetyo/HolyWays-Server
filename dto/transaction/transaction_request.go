package Transactionsdto

import "time"

type TransactionCreateRequest struct {
	DonateAmount string    `json:"donateAmount" form:"donateAmount"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"startdate"`
	UserID       int       `json:"user_id" form:"user_id"`
}
type UpdateTransactionRequest struct {
	DonateAmount string `json:"donateAmount" validate:"required"`
	Status       string `json:"status" validate:"required"`
	UserID       int    `json:"user_id" form:"user_id"`
}
