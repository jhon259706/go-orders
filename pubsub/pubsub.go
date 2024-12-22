package pubsub

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "redis-pubsub:6379",
})

// Publish publishes data to the specified channel
func Publish(channel string, data interface{}) {
	jsonData, _ := json.Marshal(data)
	rdb.Publish(ctx, channel, jsonData)
}

// Subscribe subscribes to the specified channel and processes messages using the provided handler
func Subscribe(channel string, handler func(data map[string]interface{})) {
	sub := rdb.Subscribe(ctx, channel)
	ch := sub.Channel()

	for msg := range ch {
		var data map[string]interface{}
		json.Unmarshal([]byte(msg.Payload), &data)
		handler(data)
	}
}
