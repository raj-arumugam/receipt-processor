package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"receipt-processor/internal/model"
	"receipt-processor/internal/service"
)

type ReceiptHandler struct {
	service *service.ReceiptService
}

func NewReceiptHandler(service *service.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{service: service}
}

func (h *ReceiptHandler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt model.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}

	id := h.service.ProcessReceipt(receipt)
	response := model.IDResponse{ID: id}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ReceiptHandler) GetPoints(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the path /receipts/{id}/points
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

	log.Printf("GetPoints request: ID=%s", id) // Log the extracted ID

	points, exists := h.service.GetPoints(id)
	if !exists {
		http.Error(w, "No receipt found for that ID", http.StatusNotFound)
		return
	}

	response := model.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
