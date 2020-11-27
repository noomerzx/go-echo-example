package models

// ProductResponse struct
type ProductResponse struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Amount string `json:"amount,omitempty"`
	Stock  string `json:"stock,omitempty"`
}

// Product struct
type Product struct {
	ID         int
	MerchantID int
	Name       string
	Amount     float64
	Stock      int
}
