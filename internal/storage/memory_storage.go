package storage

import (
	"fmt"
	"receipt-processor/internal/models"
	"receipt-processor/internal/service"
	"sync"
)

type memoryStorage struct {
	receipts map[string]models.ProcessedReceipt
	mu       sync.RWMutex
}

func NewMemoryStorage() service.Storage {
	return &memoryStorage{
		receipts: make(map[string]models.ProcessedReceipt),
	}
}

func (s *memoryStorage) Save(id string, receipt models.ProcessedReceipt) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.receipts[id] = receipt
	return nil
}

func (s *memoryStorage) Get(id string) (models.ProcessedReceipt, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	receipt, exists := s.receipts[id]
	if !exists {
		return models.ProcessedReceipt{}, fmt.Errorf("receipt not found")
	}
	return receipt, nil
}
