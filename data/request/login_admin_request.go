package request

type LoginAdminRequest struct {
	Email    string `validate:"required" json:"email" form:"email"`
	Password string `validate:"required" json:"password" form:"password"`
}
