package extract_controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adrian7123/rinha-backend-2024-q1-gin/repositories/customer_repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Extract(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := customer_repository.GetOne(
		bson.D{{Key: "id", Value: id}},
		bson.D{{Key: "transactions", Value: bson.D{{Key: "$slice", Value: 10}}}},
	)

	if err != nil {
		c.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"saldo": gin.H{
			"total":        customer.Balance,
			"data_extrato": time.Now(),
			"limite":       customer.Limit,
		},
		"ultimas_transacoes": customer.Transactions,
	})
}
