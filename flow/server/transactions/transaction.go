package transactions

import "time"

// Define the transaction type.
type Transaction struct {
	Id          string
	DisplayName string
}

// Define the transaction info type.
type TransactionInfo struct {
	Id     string
	Amount float64
	Date   time.Time
	UserID string
}
