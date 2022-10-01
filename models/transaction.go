package models

import "time"

type Transaction struct {
	ID           int                     `gorm:"primary_key:auto_increment" json:"id"`
	DonateAmount string                  `gorm:"type:varchar(300)" json:"donateAmount" form:"donateAmount"`
	Status       string                  `gorm:"type:varchar(300)" json:"status" form:"status"`
	CreatedAt    time.Time               `json:"startdate"`
	UserID       int                     `json:"user_id" form:"user_id"`
	User         UserTransactionResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

type TransactionResponseUser struct {
	ID           int       `json:"id"`
	DonateAmount string    `json:"donateAmount"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"startdate"`
	UserID       int       `json:"-"`
	// User         UserTransactionResponse `json:"user"`
}

func (TransactionResponseUser) TableName() string {
	return "transactions"
}
