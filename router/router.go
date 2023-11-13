package router

import (
	"basic-trade/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(AdminController *controller.AdminController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "BASIC TRADE API")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page Not Found"})
	})

	router := service.Group("/api")
	adminRouter := router.Group("/auth")
	adminRouter.POST("/register", AdminController.Register)
	adminRouter.POST("/login", AdminController.Login)

	return service
}
