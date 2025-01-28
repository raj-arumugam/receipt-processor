package api

import (
	"net/http"

	"receipt-processor/internal/service"
)

func SetupRoutes(service *service.ReceiptService) {
	handler := NewReceiptHandler(service)

	http.HandleFunc("/receipts/process", handler.ProcessReceipt)
	http.HandleFunc("/receipts/", handler.GetPoints)
}
