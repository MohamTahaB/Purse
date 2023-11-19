package api

import (
	"time"

	"example.com/m/server/transactions"
	"example.com/m/server/users"
	"example.com/m/server/utils"
	"github.com/gin-gonic/gin"
)

var transactionInfoMap = map[string]*transactions.TransactionInfo{
	"12345": {
		Id:     "12345",
		Amount: 5.0,
		Date:   time.Now(),
		UserID: "11235813",
	},
}

var userInfoMap = map[string]*users.UserInfo{
	"11235813": {
		Id:        "11235813",
		FirstName: "Mohamed Taha",
		LastName:  "Bayoumi",
		Email:     "bayoumimt@xyz.io",
		Phone:     "0123456789",
		Address:   *utils.FakeAddress(),
	},
}

var userMap = map[string]*users.User{
	"11235813": {
		Id: "11235813",
		TransactionsLog: []string{
			"12345",
		},
		Balance: 100.0,
	},
}

func RunAPI() {
	router := gin.Default()
	router.GET("trans/:id", GetTransactionInfoFromTransaction)
	router.GET("user/:id", GetUserInfoFromUser)
	router.POST("user/:id", PostTransactionToUser)
	router.Run(":1234")
}
