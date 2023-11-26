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
		productRouter.Use(middleware.Authentication())
		productRouter.GET("/", ProductController.GetByAdminID)
		productRouter.POST("/", ProductController.Create)
		productRouter.GET("/:uuid", middleware.ProductAuthorization(), ProductController.GetById)
		productRouter.PUT("/:uuid", middleware.ProductAuthorization(), ProductController.Edit)
		productRouter.DELETE("/:uuid", middleware.ProductAuthorization(), ProductController.Delete)
	}

	variantRouter := router.Group("/variants")
	{
		variantRouter.Use(middleware.Authentication())
		variantRouter.GET("/", VariantController.GetAll)
		variantRouter.POST("/", VariantController.Create)
		variantRouter.GET("/:uuid", VariantController.GetById)
		variantRouter.PUT("/:uuid", VariantController.Edit)
		variantRouter.DELETE("/:uuid", VariantController.Delete)
	}

	return service
}
