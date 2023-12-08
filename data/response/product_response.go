package response

import (
	"time"
)

type ProductResponse struct {
	ID        int               `json:"id"`
	UUID      string            `json:"uuid"`
	Name      string            `json:"name"`
	Image_URL string            `json:"image_url"`
	AdminID   uint              `json:"admin_id"`
	CreatedAt *time.Time        `json:"created_at"`
	UpdatedAt *time.Time        `json:"updated_at"`
	Variants  []VariantResponse `json:"variants"`
}

type PaginatedProductResponse struct {
	Page      int               `json:"page"`
	PageSize  int               `json:"pageSize"`
	TotalPage int               `json:"totalPage"`
	TotalData int64             `json:"totalCount"`
	Data      []ProductResponse `json:"data"`
}
