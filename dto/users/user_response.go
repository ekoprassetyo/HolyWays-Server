package Usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}
