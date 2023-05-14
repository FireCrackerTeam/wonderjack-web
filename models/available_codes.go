package models

type AvailableCodes struct {
	CodeId int64 `gorm:"primaryKey" json:"id"`
	Code string `gorm:"varchar(250);not null" json:"code"`
}
