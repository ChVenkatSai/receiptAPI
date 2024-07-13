// Package storage contains the in-memory storage logic for the Receipt Points Calculator API.
package storage

import (
    "errors"
    "sync"

    "github.com/rs/xid"
    "github.com/ChVenkatSai/receiptAPI/pkg/models"
)

//In memory map of id vs receipts and id vs points
type InMemoryStorage struct {
    receipts map[string]models.Receipt
    points   map[string]int
    mu       sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
    return &InMemoryStorage{
        receipts: make(map[string]models.Receipt),
        points:   make(map[string]int),
    }
}

//Generates id and saves receipt for that id
func (s *InMemoryStorage) SaveReceipt(receipt models.Receipt) string {
    s.mu.Lock()
    defer s.mu.Unlock()
    id := xid.New().String()
    s.receipts[id] = receipt
    return id
}

//Saves Points for an id
func (s *InMemoryStorage) SavePoints(id string, points int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.points[id] = points
}

//Returns Point for a given id
func (s *InMemoryStorage) GetPoints(id string) (int, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    points, exists := s.points[id]
    if !exists {
        return 0, errors.New("receipt not found")
    }
    return points, nil
}
