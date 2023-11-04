package users

import "example.com/m/server/utils"

// Define the user struct.
type User struct {
	Id              string
	TransactionsLog []string
	Balance         float64
}

// Define the user info: name, email, phone, address ...
type UserInfo struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   utils.Address
}
