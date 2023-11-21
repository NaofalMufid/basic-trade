package request

type UpdateProductRequest struct {
	Name      string `validate:"required" json:"name"`
	Image_URL string `validate:"required" json:"image_url"`
	Admin_ID  uint   `json:"admin_id"`
}
