package router

import (
	"receipt-processor/internal/api/handlers"
	"receipt-processor/internal/service"

	"github.com/gorilla/mux"
)

func SetupRouter(processor service.ReceiptProcessor) *mux.Router {
	r := mux.NewRouter()

	receiptHandler := handlers.NewReceiptHandler(processor)

	r.HandleFunc("/receipts/process", receiptHandler.ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", receiptHandler.GetPoints).Methods("GET")

	return r
}
