package models

// Receipt represents the structure of the receipt JSON payload
type Receipt struct {
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []Item `json:"items" validate:"required,min=1,dive,required"`
	Total        string `json:"total" validate:"required,currency"`
}

// Item represents each item in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,currency"`
}

// ProcessResponse represents the response after processing a receipt
type ProcessResponse struct {
	ID string `json:"id"`
}

// PointsResponse represents the response containing points
type PointsResponse struct {
	Points int `json:"points"`
}
