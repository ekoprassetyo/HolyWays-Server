package authdto

type RegisterResponse struct {
	Fullname string `gorm:"type:varchar(300)" json:"fullName"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Phone    string `gorm:"type:varchar(300)" json:"phone"`
}

type LoginResponse struct {
	Fullname string `gorm:"type:varchar(300)" json:"fullName"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Token    string `gorm:"type:varchar(300)" json:"token"`
	Phone    string `gorm:"type:varchar(300)" json:"phone"`
}
