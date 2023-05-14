package models

import (
	"crypto/rand"
	// "database/sql"
    "encoding/hex"
    "fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
const codeLength = 10

func ConnectDatabase(){
	db, err := gorm.Open(mysql.Open("root:@tcp(host.docker.internal:3306)/wonderjack_web"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Users{}, &RedemptionCodes{}, &AvailableCodes{})

	// Check if the available_codes table is empty
    var count int64
    err = db.Model(&AvailableCodes{}).Count(&count).Error
    if err != nil {
        panic(err)
    }

	if count == 0 {
        // Generate 300 random redemption codes and insert them into the available_codes table
        for i := 0; i < 300; i++ {
            code := make([]byte, codeLength/2)
            _, err := rand.Read(code)
            if err != nil {
                panic(err)
            }
            availableCode := hex.EncodeToString(code)[:codeLength]

            err = db.Create(&AvailableCodes{Code: availableCode}).Error
            if err != nil {
                panic(err)
            }

            fmt.Printf("Redemption code generated: %s\n", availableCode)
        }

        // Print a message to indicate that the redemption codes have been generated and inserted
        fmt.Println("300 redemption codes generated and inserted successfully")
    } else {
        fmt.Println("Table available_codes already has data. No new redemption codes generated.")
    }
	

	DB = db
}
