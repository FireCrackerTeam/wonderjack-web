package models

type Users struct {
	UserId int64 `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"varchar(250)" json:"user_name"`
	UserEmail string `gorm:"varchar(250)" json:"user_email"`
	Password string `gorm:"varchar(250)" json:"user_password"`
}