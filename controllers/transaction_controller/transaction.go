package transaction_controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/adrian7123/rinha-backend-2024-q1-gin/models"
	"github.com/adrian7123/rinha-backend-2024-q1-gin/repositories/customer_repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Transaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var transaction models.Transaction

	c.BindJSON(&transaction)

	descriptionLen := len(strings.Split(transaction.Description, ""))

	if descriptionLen <= 0 || descriptionLen > 10 {
		c.JSON(http.StatusUnprocessableEntity, "Descricao invalida")
		return
	}

	customer, err := customer_repository.GetOne(
		bson.D{{Key: "id", Value: id}},
		bson.D{{Key: "transactions", Value: bson.D{{Key: "$slice", Value: 9}}}},
	)

	if err != nil {
		c.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}

	if err := customer.Transact(transaction); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	customer.ReverseTransactions()

	customer_repository.Update(*customer)

	c.JSON(http.StatusOK, gin.H{
		"saldo":  customer.Balance,
		"limite": customer.Limit,
	})
}
