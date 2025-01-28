package service

import "receipt-processor/internal/models"

type ReceiptProcessor interface {
	ProcessReceipt(receipt models.Receipt) (string, error)
	GetPoints(id string) (int, error)
}

type PointsCalculator interface {
	Calculate(receipt models.Receipt) int
}

type Storage interface {
	Save(id string, receipt models.ProcessedReceipt) error
	Get(id string) (models.ProcessedReceipt, error)
}
