package request

type CreateVariantRequest struct {
	Variant_Name string `validate:"required" json:"variant_name" form:"variant_name"`
	Quantity     int    `validate:"required" json:"quantity" form:"quantity"`
	Product_ID   uint   `validate:"required" json:"product_id" form:"product_id"`
}
