package request

type CreateAdminRequest struct {
	Name     string `validate:"required" form:"name"`
	Email    string `validate:"required,email" form:"email"`
	Password string `validate:"required,min=8" form:"password"`
}
