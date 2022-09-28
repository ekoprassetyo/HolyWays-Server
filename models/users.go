package models

type User struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"type:varchar(300)" json:"fullName"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
	Phone    string `gorm:"type:varchar(300)" json:"phone"`
}
