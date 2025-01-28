package service

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"receipt-processor/internal/model"

	"github.com/google/uuid"
)

type ReceiptService struct {
	receipts map[string]int
	mutex    *sync.Mutex
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{
		receipts: make(map[string]int),
		mutex:    &sync.Mutex{},
	}
}

func (s *ReceiptService) ProcessReceipt(receipt model.Receipt) string {
	id := uuid.New().String()
	points := calculatePoints(receipt)

	s.mutex.Lock()
	s.receipts[id] = points
	s.mutex.Unlock()

	log.Printf("Receipt processed: ID=%s, Points=%d", id, points) // Add logging
	return id
}

func (s *ReceiptService) GetPoints(id string) (int, bool) {
	s.mutex.Lock()
	points, exists := s.receipts[id]
	s.mutex.Unlock()

	log.Printf("GetPoints: ID=%s, Exists=%v, Points=%d", id, exists, points) // Add logging
	return points, exists
}

func calculatePoints(receipt model.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	alphanumeric := reg.ReplaceAllString(receipt.Retailer, "")
	points += len(alphanumeric)

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
