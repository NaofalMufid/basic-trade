package router

import (
	"basic-trade/controller"
	"basic-trade/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(AdminController *controller.AdminController, ProductController *controller.ProductController, VariantController *controller.VarianController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "BASIC TRADE API")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page Not Found"})
	})

	router := service.Group("/api")
	adminRouter := router.Group("/auth")
	{
		adminRouter.POST("/register", AdminController.Register)
		adminRouter.POST("/login", AdminController.Login)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", ProductController.GetAll)
		productRouter.GET("/:uuid", ProductController.GetById)

		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", ProductController.Create)
		productRouter.PUT("/:uuid", middleware.ProductAuthorization(), ProductController.Edit)
		productRouter.DELETE("/:uuid", middleware.ProductAuthorization(), ProductController.Delete)
	}

	variantRouter := router.Group("/variants")
	{
		variantRouter.GET("/", VariantController.GetAll)
		variantRouter.GET("/:uuid", VariantController.GetById)

		variantRouter.Use(middleware.Authentication())
		variantRouter.POST("/", middleware.ProductAuthorization(), VariantController.Create)
		variantRouter.PUT("/:uuid", middleware.VariantAuthorization(), VariantController.Edit)
		variantRouter.DELETE("/:uuid", middleware.VariantAuthorization(), VariantController.Delete)
	}

	return service
}
