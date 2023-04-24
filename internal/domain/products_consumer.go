package domain

type ProductsConsumer struct {
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
}
