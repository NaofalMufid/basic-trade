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
		var productUUID string
		if paramUUID, ok := ctx.Params.Get("uuid"); ok {
			productUUID = paramUUID
		} else {
			var request struct {
				ProductID string `form:"product_id" json:"product_id" binding:"required"`
			}

			if err := ctx.ShouldBind(&request); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "Invalid request payload",
				})
				return
			}
			productUUID = request.ProductID
		}

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

		if getProduct.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func VariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := config.DBInit()
		variantUUID := ctx.Param("uuid")

		adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
		adminID := uint(adminData["id"].(float64))

		// get product_id from variant
		var getVariant model.Variants
		err := db.Select("product_id").Where("uuid = ?", variantUUID).First(&getVariant).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data not found",
			})
			return
		}
		// get admin_id from product
		var getProduct model.Products
		err = db.Select("admin_id").Where("uuid = ?", getVariant.ProductID).First(&getProduct).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data not found",
			})
			return
		}

		if getProduct.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
