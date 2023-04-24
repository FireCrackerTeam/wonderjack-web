package models

type RedemptionCodes struct {
	RedemptionId int64 `gorm:"primaryKey" json:"id"`
	UserId int64 `gorm:"not null" json:"user_id"`
	Code string `gorm:"varchar(250);not null" json:"code"`
	// Users Users `gorm:"foreignKey:UserId;references:UserId"`
}