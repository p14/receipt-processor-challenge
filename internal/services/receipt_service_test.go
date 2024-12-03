package services

import (
	"receipt-processor/internal/models"
	"testing"
)

// TestCalculatePoints uses table-driven tests to verify the calculatePoints function.
func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Valid Receipt - Target",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name: "Valid Receipt - M&M Corner Market",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := calculatePoints(tt.receipt)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
