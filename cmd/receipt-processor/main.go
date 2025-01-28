package main

import (
	"log"
	"net/http"

	"receipt-processor/internal/api"
	"receipt-processor/internal/service"
)

func main() {
	receiptService := service.NewReceiptService()
	api.SetupRoutes(receiptService)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
