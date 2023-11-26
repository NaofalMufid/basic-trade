package request

type UpdateVariantRequest struct {
	Variant_Name string `validate:"required" form:"variant_name" json:"variant_name"`
	Quantity     int    `validate:"required" form:"quantity" json:"quantity"`
}
