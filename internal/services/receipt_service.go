package services

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/internal/models"
	"receipt-processor/internal/repository"
)

// ErrReceiptNotFound is returned when a receipt ID does not exist.
var ErrReceiptNotFound = errors.New("receipt not found")

// ReceiptService defines the methods for processing receipts.
type ReceiptService interface {
	ProcessReceipt(receipt models.Receipt) (string, error)
	GetPoints(id string) (int, error)
}

// receiptService is the concrete implementation of ReceiptService.
type receiptService struct {
	repo repository.ReceiptRepository
}

// NewReceiptService creates a new ReceiptService.
func NewReceiptService(repo repository.ReceiptRepository) ReceiptService {
	return &receiptService{
		repo: repo,
	}
}

// ProcessReceipt processes the receipt and stores it with a unique ID.
func (s *receiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	id, err := s.repo.Create(receipt)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetPoints retrieves the points for a given receipt ID.
func (s *receiptService) GetPoints(id string) (int, error) {
	receipt, err := s.repo.Get(id)
	if err != nil {
		return 0, ErrReceiptNotFound
	}

	points, err := calculatePoints(receipt)
	if err != nil {
		return 0, err
	}

	return points, nil
}

// calculatePoints remains the same as previously defined.
func calculatePoints(receipt models.Receipt) (int, error) {
	totalPoints := 0

	// 1. One point for every alphanumeric character in the retailer name.
	alphanumeric := regexp.MustCompile(`[A-Za-z0-9]`)
	alphanumericMatches := alphanumeric.FindAllString(receipt.Retailer, -1)
	totalPoints += len(alphanumericMatches)

	// 2. 50 points if the total is a round dollar amount with no cents.
	totalAmount, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, err
	}
	if totalAmount == math.Trunc(totalAmount) {
		totalPoints += 50
	}

	// 3. 25 points if the total is a multiple of 0.25.
	if math.Mod(totalAmount, 0.25) == 0 {
		totalPoints += 25
	}

	// 4. 5 points for every two items on the receipt.
	itemCount := len(receipt.Items)
	totalPoints += (itemCount / 2) * 5

	// 5. If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, err
			}
			points := int(math.Ceil(price * 0.2))
			totalPoints += points
		}
	}

	// 6. 6 points if the day in the purchase date is odd.
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	if purchaseDate.Day()%2 != 0 {
		totalPoints += 6
	}

	// 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	hour := purchaseTime.Hour()
	minute := purchaseTime.Minute()
	totalMinutes := hour*60 + minute
	if totalMinutes > 14*60 && totalMinutes < 16*60 {
		totalPoints += 10
	}

	return totalPoints, nil
}
