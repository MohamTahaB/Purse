package utils

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/m/server/transactions"
	"github.com/icrowley/fake"
)

func FakeAddress() *Address {
	return &Address{
		Line1:   fake.StreetAddress(),
		ZipCode: fake.Zip(),
		City:    fake.City(),
		Country: fake.Country(),
	}
}

func FakeTransaction(id string) *transactions.Transaction {
	return &transactions.Transaction{
		Id:          id,
		DisplayName: fmt.Sprintf("From me to %s.", fake.FullName()),
	}
}

func FakeTransactionInfo() *transactions.TransactionInfo {

	monthMap := map[string]time.Month{
		"January":   time.January,
		"February":  time.February,
		"March":     time.March,
		"April":     time.April,
		"May":       time.May,
		"June":      time.June,
		"July":      time.July,
		"August":    time.August,
		"September": time.September,
		"October":   time.October,
		"November":  time.November,
		"December":  time.December,
	}

	return &transactions.TransactionInfo{
		Id:     fake.CharactersN(7),
		Amount: (rand.Float64() - 0.5) * 50.0,
		Date:   time.Date(fake.Year(2020, 2023), monthMap[fake.Month()], fake.Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, &time.Location{}),
	}
}
