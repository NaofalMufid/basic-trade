package request

import "mime/multipart"

type UpdateProductRequest struct {
	Name    string                `binding:"required" form:"name"`
	Image   *multipart.FileHeader `binding:"required" form:"image"`
	AdminID uint                  `json:"admin_id"`
}
