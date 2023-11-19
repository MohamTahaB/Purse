package api

import (
	"net/http"

	"example.com/m/server/transactions"
	"example.com/m/server/users"
	"github.com/gin-gonic/gin"
)

// Gets the user info from the user id.
func GetUserInfoFromUser(c *gin.Context) {
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

	user := users.User{
		Id: id,
	}
	userInfoRes, err := user.FetchUserInfoFromUser(userInfoMap)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, userInfoRes)
	}
}

func PostTransactionToUser(c *gin.Context) {
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

	if _, ok := userMap[id]; !ok {
		c.String(http.StatusNotFound, "user of id %s not found", id)
	} else {
		var newTransaction transactions.RawTransaction
		if err := c.BindJSON(newTransaction); err != nil {
			return
		}

		userMap[id].TransactionsLog = append(userMap[id].TransactionsLog, newTransaction.Transaction.Id)
		transactionInfoMap[newTransaction.Transaction.Id] = newTransaction.TransactionInfo
		c.IndentedJSON(http.StatusCreated, newTransaction)
	}
}
