package Fundsdto

type FundCreateRequest struct {
	Title         string `json:"title" validate:"required"`
	Thumbnail     string `json:"thumbnail" validate:"required"`
	Goal          string `json:"goal" validate:"required"`
	Description   string `json:"description" validate:"required"`
	TransactionID int    `json:"transaction_id" form:"transaction_id"`
}

type UpdateFundRequest struct {
	Title       string `json:"title" validate:"required"`
	Thumbnail   string `json:"thumbnail" validate:"required"`
	Goal        string `json:"goal" validate:"required"`
	Description string `json:"description" validate:"required"`
}
