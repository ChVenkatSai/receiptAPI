package service

import (
    "math"
    "strconv"
    "strings"
    "time"

    "github.com/ChVenkatSai/receiptAPI/internal/storage"
    "github.com/ChVenkatSai/receiptAPI/pkg/models"
)

type Service struct {
    Store *storage.InMemoryStorage
}

func NewService(s *storage.InMemoryStorage) *Service {
    return &Service{Store: s}
}

func (s *Service) ProcessReceipt(receipt models.Receipt) string {
    id := s.Store.SaveReceipt(receipt)
    points := s.calculatePoints(receipt)
    s.Store.SavePoints(id, points)
    return id
}

func (s *Service) GetPoints(id string) (int, error) {
    return s.Store.GetPoints(id)
}

func (s *Service) calculatePoints(receipt models.Receipt) int {
    points := 0

    // Rule 1: One point for every alphanumeric character in the retailer name.
    for _, char := range receipt.Retailer {
        if strings.ContainsAny(string(char), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") {
            points++
        }
    }

    // Rule 2: 50 points if the total is a round dollar amount with no cents.
    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if total == float64(int(total)) {
            points += 50
        }
    }

    // Rule 3: 25 points if the total is a multiple of 0.25.
    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if math.Mod(total, 0.25) == 0 {
            points += 25
        }
    }

    // Rule 4: 5 points for every two items on the receipt.
    points += (len(receipt.Items) / 2) * 5

    // Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
    for _, item := range receipt.Items {
        if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
            if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
                points += int(math.Ceil(price * 0.2))
            }
        }
    }

    // Rule 6: 6 points if the day in the purchase date is odd.
    if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
        if date.Day()%2 != 0 {
            points += 6
        }
    }

    // Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
    if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
        if t.Hour() >= 14 && t.Hour() < 16 {
            points += 10
        }
    }

    return points
}
