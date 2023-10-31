package main

import (
	"basic-trade/config"
	"basic-trade/model"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := config.DBInit()

	db.Table("admins").AutoMigrate(&model.Admins{})
	db.Table("products").AutoMigrate(&model.Products{})
	db.Table("variants").AutoMigrate(&model.Variants{})
	fmt.Println("Migrasi dulu")
}
