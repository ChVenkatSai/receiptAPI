// Package models contains the data models for the Receipt Points Calculator API.
package models

// Receipt represents a purchase receipt.
type Receipt struct {
    Retailer     string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
    Total        string `json:"total"`
}

// Item represents an item on a receipt.
type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price            string `json:"price"`
}
