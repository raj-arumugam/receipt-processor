package service

import (
	"receipt-processor/internal/models"

	"github.com/google/uuid"
)

type receiptService struct {
	calculator PointsCalculator
	storage    Storage
}

func NewReceiptService(calculator PointsCalculator, storage Storage) ReceiptProcessor {
	return &receiptService{
		calculator: calculator,
		storage:    storage,
	}
}

func (s *receiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	points := s.calculator.Calculate(receipt)
	id := uuid.New().String()

	processedReceipt := models.ProcessedReceipt{
		ID:     id,
		Points: points,
	}

	if err := s.storage.Save(id, processedReceipt); err != nil {
		return "", err
	}

	return id, nil
}

func (s *receiptService) GetPoints(id string) (int, error) {
	receipt, err := s.storage.Get(id)
	if err != nil {
		return 0, err
	}
	return receipt.Points, nil
}
