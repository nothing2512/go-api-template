package models

import (
	"fmt"
	"main/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetSqlConnection() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s database=%s port=%s TimeZone=Asia/Jakarta",
		constants.DB_HOST,
		constants.DB_USER,
		constants.DB_PASS,
		constants.DB_NAME,
		constants.DB_NAME,
		constants.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
