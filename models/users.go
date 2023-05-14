package models

type Users struct {
	UserId int64 `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"varchar(250);not null" json:"user_name"`
	UserEmail string `gorm:"varchar(250);unique;not null" json:"user_email" binding:"required,email"`
	Password string `gorm:"varchar(250);not null" json:"user_password"`
	RedemptionCodes []RedemptionCodes `gorm:"foreignKey:UserId"`
}