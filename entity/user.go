package entity

import "time"

type User struct {
	ID              int        `json:"id" gorm:"column:id;primary_key"`
	FullName        string     `json:"full_name" gorm:"column:full_name"`
	PhoneNumber     string     `json:"phone_number" gorm:"column:phone_number"`
	Password        string     `json:"password" gorm:"column:password"`
	AccountSalt     string     `json:"account_salt" gorm:"column:account_salt"`
	SuccessfulLogin int        `json:"successfuul_login" gorm:"column:successful_login"`
	CreatedAt       time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (e *User) TableName() string {
	return `user`
}
