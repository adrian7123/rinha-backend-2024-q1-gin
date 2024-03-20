package routes

import (
	"github.com/adrian7123/rinha-backend-2024-q1-gin/controllers/extract_controller"
	"github.com/adrian7123/rinha-backend-2024-q1-gin/controllers/transaction_controller"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/clientes/:id/extrato", extract_controller.Extract)
	r.POST("/clientes/:id/transacoes", transaction_controller.Transaction)
}
