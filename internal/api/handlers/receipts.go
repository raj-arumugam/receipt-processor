package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/internal/models"
	"receipt-processor/internal/service"

	"github.com/gorilla/mux"
)

type ReceiptHandler struct {
	processor service.ReceiptProcessor
}

func NewReceiptHandler(processor service.ReceiptProcessor) *ReceiptHandler {
	return &ReceiptHandler{processor: processor}
}

func (h *ReceiptHandler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Please verify input.", http.StatusBadRequest)
		return
	}

	id, err := h.processor.ProcessReceipt(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *ReceiptHandler) GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, err := h.processor.GetPoints(id)
	if err != nil {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
