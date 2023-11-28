package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	sqlInfo := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True&loc-Local", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
