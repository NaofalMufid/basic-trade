package router

import (
	"basic-trade/controller"
	"basic-trade/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(AdminController *controller.AdminController, ProductController *controller.ProductController) *gin.Engine {
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
		productRouter.GET("/:productUUID", middleware.ProductAuthorization(), ProductController.GetById)
		productRouter.PUT("/:productUUID", middleware.ProductAuthorization(), ProductController.Edit)
		productRouter.DELETE("/:productUUID", middleware.ProductAuthorization(), ProductController.Delete)
	}

	return service
}