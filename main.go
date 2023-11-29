package main

import (
	"basic-trade/config"
	"basic-trade/controller"
	"basic-trade/model"
	"basic-trade/repository"
	"basic-trade/router"
	"basic-trade/service"
	"fmt"
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

	db.Table("admins").Debug().AutoMigrate(&model.Admins{})
	db.Table("products").Debug().AutoMigrate(&model.Products{})
	db.Table("variants").Debug().AutoMigrate(&model.Variants{})

	// Init Repository
	adminRepository := repository.NewAdminRepository(db)
	productRepository := repository.NewProductRepository(db)
	variantRepository := repository.NewVariantRepository(db)

	// Init Service
	adminService := service.NewAdminServiceImpl(adminRepository, validate)
	productService := service.NewProductService(productRepository, variantRepository, validate)
	variantService := service.NewVariantService(variantRepository, validate)

	// Init Controller
	adminController := controller.NewAdminController(adminService)
	productController := controller.NewProductController(productService, validate)
	variantController := controller.NewVariantController(variantService, validate)

	// Router
	routes := router.NewRouter(adminController, productController, variantController)
	port := os.Getenv("PORT")
	address := fmt.Sprintf(":%s", port)
	err := routes.Run(address)
	if err != nil {
		panic(err)
	}
}
