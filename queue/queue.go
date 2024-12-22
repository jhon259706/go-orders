package queue

import (
	"context"
	"encoding/json"

	"go-orders/pubsub"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "redis-queue:6379",
})

func Enqueue (data interface{}, queueName string) {
	jsonData, _ := json.Marshal(data)
	rdb.LPush(ctx, queueName, jsonData)
}

func ProcessQueue(queueName string, pubSubChannel string) {
    for {
		result, err := rdb.BRPop(ctx, 0, queueName).Result()
        if err != nil {
            log.Println("Error popping from orderQueue:", err)
            continue
        }

		var data map[string]interface{}
		json.Unmarshal([]byte(result[1]), &data)
		log.Println("Enqueuing order/shipment:", data)

		pubsub.Publish(pubSubChannel, data)
    }
}
