package api

import (
	"net/http"

	"example.com/m/server/transactions"
	"github.com/gin-gonic/gin"
)

func GetTransactionInfoFromTransaction(c *gin.Context) {
	id := c.Param("id")

	// Check if there are leading or remaining slashes in the id, to check if can be done in a better way.
	start, end := 0, len(id)
	if id[0] == '/' {
		start++
	}
	if id[len(id)-1] == '/' {
		end--
	}
	id = id[start:end]

	transaction := &transactions.Transaction{
		Id: id,
	}
	transactionInfoRes, err := transaction.FetchTransactionInfoFromTransaction(transactionInfoMap)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, transactionInfoRes)
	}
}
