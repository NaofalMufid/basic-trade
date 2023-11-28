package response

import (
	"time"

	"github.com/google/uuid"
)

type VariantResponse struct {
	ID           int        `json:"id"`
	UUID         uuid.UUID  `json:"uuid"`
	Variant_Name string     `json:"variant_name"`
	Quantity     int        `json:"quantity"`
	ProductID    uint       `json:"product_id"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type PaginatedVariantResponse struct {
	Page      int               `json:"page"`
	PageSize  int               `json:"pageSize"`
	TotalPage int               `json:"totalPage"`
	TotalData int64             `json:"totalCount"`
	Data      []VariantResponse `json:"data"`
}
