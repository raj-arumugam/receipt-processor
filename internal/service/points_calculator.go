package service

import (
	"math"
	"receipt-processor/internal/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type pointsCalculator struct{}

func NewPointsCalculator() PointsCalculator {
	return &pointsCalculator{}
}

func (pc *pointsCalculator) Calculate(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}

	// Rule 2: 50 points if the total is a round dollar amount
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(total*100, 25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for items with descriptions of length multiple of 3
	for _, item := range receipt.Items {
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: Award 5 points if the total is greater than 10.00
	// This is included because the solution was partially assisted by an AI tool.
	if total > 10.00 {
		points += 5
	}

	// Rule 7: 6 points if the day is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Rule 8: 10 points if purchase time is between 2:00 PM and 4:00 PM
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	targetStart, _ := time.Parse("15:04", "14:00")
	targetEnd, _ := time.Parse("15:04", "16:00")

	if purchaseTime.After(targetStart) && purchaseTime.Before(targetEnd) {
		points += 10
	}

	return points
}
