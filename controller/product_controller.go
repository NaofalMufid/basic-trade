package controller

import (
	"basic-trade/data/request"
	"basic-trade/data/response"
	"basic-trade/helper"
	"basic-trade/model"
	"basic-trade/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ProductController struct {
	productService service.ProductService
	Validate       *validator.Validate
}

func NewProductController(service service.ProductService, validate *validator.Validate) *ProductController {
	return &ProductController{
		productService: service,
		Validate:       validate,
	}
}

func (c ProductController) Create(ctx *gin.Context) {
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))

	productRequest := request.CreateProductRequest{}

	if err := ctx.ShouldBind(&productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Validate.Struct(productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := helper.RemoveExtension(productRequest.Image.Filename)
	uploadResult, err := helper.UploadFile(productRequest.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUUID := uuid.New()
	product := model.Products{
		UUID:      newUUID,
		Name:      productRequest.Name,
		Image_URL: uploadResult,
		AdminID:   adminID,
	}
	if err := c.productService.Create(product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    201,
		Status:  true,
		Message: "successfully create product",
		Data:    nil,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

func (c ProductController) Edit(ctx *gin.Context) {
	productUUID := ctx.Param("uuid")

	productRequest := request.UpdateProductRequest{}
	if err := ctx.ShouldBind(&productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Validate.Struct(productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := helper.RemoveExtension(productRequest.Image.Filename)
	uploadResult, err := helper.UploadFile(productRequest.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProduct := model.Products{
		Name:      productRequest.Name,
		Image_URL: uploadResult,
	}
	err = c.productService.Update(productUUID, updateProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully update product",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c ProductController) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	search := ctx.Query("search")

	productResponse, err := c.productService.GetAll(page, size, search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully get all product",
		Data:    productResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c ProductController) GetById(ctx *gin.Context) {
	productUUID := ctx.Param("uuid")

	productResponse := c.productService.GetById(productUUID)
	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully get product",
		Data:    productResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c ProductController) Delete(ctx *gin.Context) {
	productUUID := ctx.Param("uuid")
	c.productService.Delete(productUUID)
	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully delete product",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
