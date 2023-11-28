package controller

import (
	"basic-trade/data/request"
	"basic-trade/data/response"
	"basic-trade/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminController struct {
	adminService service.AdminService
	Validate     *validator.Validate
}

func NewAdminController(service service.AdminService) *AdminController {
	return &AdminController{adminService: service}
}

func (controller AdminController) Register(ctx *gin.Context) {
	createAdminRequest := request.CreateAdminRequest{}

	err := ctx.ShouldBind(&createAdminRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAdmin, err := controller.adminService.Register(createAdminRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:    201,
		Status:  true,
		Message: "successfully register admin",
		Data:    newAdmin,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller AdminController) Login(ctx *gin.Context) {
	loginRequest := request.LoginAdminRequest{}

	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := controller.adminService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  true,
		Message: "success login admin",
		Data:    data,
	}

	ctx.JSON(200, webResponse)
}
