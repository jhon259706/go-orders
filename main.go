package main

import (
	"go-orders/handlers"
	"go-orders/pubsub"
	"go-orders/queue"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/order/create", handlers.CreateOrder)
	r.POST("/order/ship", handlers.ShipOrder)

	go queue.ProcessQueue("orderQueue", "orderChannel")
	go queue.ProcessQueue("shipmentQueue", "shipmentChannel")
	go pubsub.Subscribe("orderChannel", handleMessage)
	go pubsub.Subscribe("shipmentChannel", handleMessage)

	log.Println("Server started on :8080")
	r.Run(":8080")
}

func handleMessage(data map[string]interface{}) {
	log.Println("Received order message:", data)
}
