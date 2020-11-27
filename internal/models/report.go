package models

// ProductSell struct
type ProductSell struct {
	Name       string `json:"name,omitempty"`
	SellVolume int64  `json:"selling_volume,omitempty"`
}

// Report struct
type Report struct {
	Date       string        `json:"date,omitempty"`
	Products   []ProductSell `json:"products,omitempty"`
	Accumulate float64       `json:"accumulate,omitempty"`
}
