package main

import (
	"basic-trade/config"
	"basic-trade/controller"
	"basic-trade/model"
	"basic-trade/repository"
	"basic-trade/router"
	"basic-trade/service"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	db := config.DBInit()
	validate := validator.New()

	db.Table("admins").AutoMigrate(&model.Admins{})
	db.Table("products").AutoMigrate(&model.Products{})
	db.Table("variants").AutoMigrate(&model.Variants{})

	// Init Repository
	adminRepository := repository.NewAdminRepository(db)
	productRepository := repository.NewProductRepository(db)

	// Init Service
	adminService := service.NewAdminServiceImpl(adminRepository, validate)
	productService := service.NewProductService(productRepository, validate)

	// Init Controller
	adminController := controller.NewAdminController(adminService)
	productController := controller.NewProductController(productService, validate)

	// Router
	routes := router.NewRouter(adminController, productController)

	port := os.Getenv("API_PORT")
	server := &http.Server{
		Addr:    port,
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
