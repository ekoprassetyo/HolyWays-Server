package models

type Fund struct {
	ID            int                     `gorm:"primary_key:auto_increment" json:"id"`
	Title         string                  `gorm:"type:varchar(300)" json:"title" form:"title"`
	Thumbnail     string                  `gorm:"type:varchar(300)" json:"thumbnail" form:"image"`
	Goal          string                  `gorm:"type:varchar(300)" json:"goal" form:"goal"`
	Description   string                  `gorm:"type:text" json:"description" form:"description"`
	TransactionID int                     `json:"transaction_id" form:"transaction_id"`
	Transaction   TransactionResponseUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaction"`
	UserID        int                     `json:"user_id" form:"user_id"`
	User          UserTransactionResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

type FundRespon struct {
	Title         string                  `json:"title"`
	Thumbnail     string                  `json:"thumbnail"`
	Goal          string                  `json:"goal"`
	Description   string                  `json:"description"`
	TransactionID int                     `json:"-"`
	Transaction   TransactionResponseUser `json:"transaction"`
	UserID        int                     `json:"user_id"`
	User          UserTransactionResponse `json:"user"`
}

func (FundRespon) TableName() string {
	return "funds"
}
