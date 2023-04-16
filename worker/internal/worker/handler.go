package worker

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"worker/internal/common"
	"worker/internal/providers"
)

func handler(redisProvider *providers.RedisProvider, deliveries <-chan amqp091.Delivery, done chan error) {
	cleanup := func() {
		log.Printf("handle: deliveries channel closed")
		done <- nil
	}

	defer cleanup()

	for d := range deliveries {
		var command common.CommandV1
		err := json.Unmarshal(d.Body, &command)

		if err != nil {
			_ = d.Ack(false)
			fmt.Printf("got poison message: %s; %s", err, d.Body)
			continue
		}

		redisProvider.ProceedCommand(command)

		_ = d.Ack(false)
	}
}
