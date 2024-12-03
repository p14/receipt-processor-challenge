package repository

import (
	"errors"
	"receipt-processor/internal/models"
	"sync"

	"github.com/google/uuid"
)

// ReceiptRepository defines methods for interacting with receipt data.
type ReceiptRepository interface {
	Create(receipt models.Receipt) (string, error)
	Get(id string) (models.Receipt, error)
}

// receiptRepository is the in-memory implementation of ReceiptRepository.
type receiptRepository struct {
	receipts   map[string]models.Receipt
	receiptsMu sync.RWMutex
}

// NewReceiptRepository creates a new ReceiptRepository.
func NewReceiptRepository() ReceiptRepository {
	return &receiptRepository{
		receipts: make(map[string]models.Receipt),
	}
}

// Create adds a new receipt and returns its ID.
func (r *receiptRepository) Create(receipt models.Receipt) (string, error) {
	id := uuid.New().String()

	r.receiptsMu.Lock()
	r.receipts[id] = receipt
	r.receiptsMu.Unlock()

	return id, nil
}

// Get retrieves a receipt by ID.
func (r *receiptRepository) Get(id string) (models.Receipt, error) {
	r.receiptsMu.RLock()
	receipt, exists := r.receipts[id]
	r.receiptsMu.RUnlock()

	if !exists {
		return models.Receipt{}, errors.New("receipt not found")
	}

	return receipt, nil
}
