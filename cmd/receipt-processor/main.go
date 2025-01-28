package main

import (
	"log"
	"net/http"

	"receipt-processor/internal/api/router"
	"receipt-processor/internal/service"
	"receipt-processor/internal/storage"
)

func main() {
	// Initialize dependencies
	calculator := service.NewPointsCalculator()
	storage := storage.NewMemoryStorage()
	receiptService := service.NewReceiptService(calculator, storage)

	// Setup router with dependencies
	r := router.SetupRouter(receiptService)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
