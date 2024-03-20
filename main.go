package main

import (
	"fmt"

	"github.com/adrian7123/rinha-backend-2024-q1-gin/configs"
	"github.com/adrian7123/rinha-backend-2024-q1-gin/models"
	"github.com/adrian7123/rinha-backend-2024-q1-gin/repositories/customer_repository"
	"github.com/adrian7123/rinha-backend-2024-q1-gin/routes"
	"github.com/gin-gonic/gin"
)

// id   limite     saldo inicial
// 1    100000     0
// 2    80000      0
// 3    1000000    0
// 4    10000000   0
// 5    500000     0
func initializeDB() {
	customer_repository.DeleteAll()

	customers := []interface{}{
		models.Customer{
			Id:           1,
			Limit:        100_000,
			Balance:      0,
			Transactions: []models.Transaction{},
		},
		models.Customer{
			Id:           2,
			Limit:        80_000,
			Balance:      0,
			Transactions: []models.Transaction{},
		},
		models.Customer{
			Id:           3,
			Limit:        1_000_000,
			Balance:      0,
			Transactions: []models.Transaction{},
		},
		models.Customer{
			Id:           4,
			Limit:        10_000_000,
			Balance:      0,
			Transactions: []models.Transaction{},
		},
		models.Customer{
			Id:           5,
			Limit:        500_000,
			Balance:      0,
			Transactions: []models.Transaction{},
		},
	}

	customer_repository.CreateMany(customers)
}

func main() {

	port, err := configs.Env("PORT")

	if err != nil {
		port = "3003"
	}

	initializeDB()

	r := gin.Default()

	routes.Routes(r)

	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
