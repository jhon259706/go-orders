package handlers

import (
	"go-orders/queue"

	"github.com/gin-gonic/gin"
)

type Shipment struct {
	ID      string `json:"id"`
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

func ShipOrder(c *gin.Context) {
	var shipment Shipment
	if err := c.BindJSON(&shipment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	queue.Enqueue(shipment, "shipmentQueue")
	c.JSON(200, gin.H{
		"message":     "Shipment created",
		"shipment_id": shipment.ID,
	})
}
