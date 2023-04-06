package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	db, err := gorm.Open(mysql.Open("root:@tcp(localhos:3306)/wonderjack-web"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Users{})

	DB = db
}