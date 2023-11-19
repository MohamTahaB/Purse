package users

import "example.com/m/server/utils"

// Define the user struct.
type User struct {
	Id              string   `json:"id"`
	TransactionsLog []string `json:"transactionLog"`
	Balance         float64  `json:"balance"`
}

// Define the user info: name, email, phone, address ...
type UserInfo struct {
	Id        string        `json:"id"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Address   utils.Address `json:"address"`
}
