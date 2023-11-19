package api

import (
	"time"

	"example.com/m/server/transactions"
	"github.com/gin-gonic/gin"
)

func TransactionInfoMap() map[string]*transactions.TransactionInfo {
	return map[string]*transactions.TransactionInfo{
		"12345": {
			Id:     "12345",
			Amount: 5.0,
			Date:   time.Now(),
			UserID: "11235813",
		},
	}
}

func RunAPI() {
	router := gin.Default()
	router.GET("trans/:id", GetTransactionInfoFromTransaction)
	router.Run(":1234")
}
