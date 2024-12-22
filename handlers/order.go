package handlers

import (
	"go-orders/queue"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID           string `json:"id"`
	CustomerName string `json:"customer_name"`
	OrderAmount  int    `json:"order_amount"`
}

func CreateOrder(c *gin.Context) {
	var order Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	queue.Enqueue(order, "orderQueue")
	c.JSON(200, gin.H{
		"message":  "Order created",
		"order_id": order.ID,
	})
}
