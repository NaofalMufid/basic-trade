package controller

import (
	"basic-trade/data/request"
	"basic-trade/data/response"
	"basic-trade/model"
	"basic-trade/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type VarianController struct {
	variantService service.VariantService
	Validate       *validator.Validate
}

func NewVariantController(service service.VariantService, validate *validator.Validate) *VarianController {
	return &VarianController{
		variantService: service,
		Validate:       validate,
	}
}

func (c VarianController) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	search := ctx.Query("search")
	variantResponse, err := c.variantService.GetAll(page, size, search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully get all variant",
		Data:    variantResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c VarianController) GetById(ctx *gin.Context) {
	variantID := ctx.Param("uuid")

	variantResponse, err := c.variantService.GetById(variantID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully get variant by uuid",
		Data:    variantResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c VarianController) Create(ctx *gin.Context) {
	variantRequest := request.CreateVariantRequest{}
	if err := ctx.ShouldBind(&variantRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Validate.Struct(&variantRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUUID := uuid.NewString()
	variant := model.Variants{
		UUID:         newUUID,
		Variant_Name: variantRequest.Variant_Name,
		Quantity:     variantRequest.Quantity,
		ProductID:    variantRequest.ProductID,
	}
	new_variant, err := c.variantService.Create(variant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    201,
		Status:  true,
		Message: "successfully create variant",
		Data:    new_variant,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

func (c VarianController) Edit(ctx *gin.Context) {
	variantID := ctx.Param("uuid")

	variantRequest := request.UpdateVariantRequest{}
	if err := ctx.ShouldBind(&variantRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Validate.Struct(&variantRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateVariant := model.Variants{
		Variant_Name: variantRequest.Variant_Name,
		Quantity:     variantRequest.Quantity,
	}

	variantUpdate, err := c.variantService.Update(variantID, updateVariant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully update variant",
		Data:    variantUpdate,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (c VarianController) Delete(ctx *gin.Context) {
	variantID := ctx.Param("uuid")
	err := c.variantService.Delete(variantID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "successfully delete variant",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
