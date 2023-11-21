package middleware

import (
	"basic-trade/config"
	"basic-trade/model"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := config.DBInit()
		productUUID := ctx.Param("productUUID")

		adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
		adminID := uint(adminData["id"].(float64))

		var getProduct model.Products
		err := db.Select("admin_id").Where("uuid = ?", productUUID).First(&getProduct).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data not found",
			})
			return
		}

		if getProduct.Admin_ID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
