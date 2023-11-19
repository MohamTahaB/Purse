package transactions

import "time"

// Define the transaction type.
type Transaction struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

// Define the transaction info type.
type TransactionInfo struct {
	Id     string    `json:"id"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
	UserID string    `json:"userId"`
}
