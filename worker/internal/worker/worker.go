package worker

import (
	"fmt"
	amqp091 "github.com/rabbitmq/amqp091-go"
	"worker/internal/config"
	"worker/pkg/amqp"

	"log"
)

func handler(deliveries <-chan amqp091.Delivery, done chan error) {
	cleanup := func() {
		log.Printf("handle: deliveries channel closed")
		done <- nil
	}

	defer cleanup()

	for d := range deliveries {
		fmt.Printf(
			"got new message: [%v] %q\n",
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
}

func Consume(cfg config.Config) {
	consumer, err := amqp.NewConsumer(cfg.AMQPConnectionString, cfg.AMQPExchangeName, cfg.AMQPExchangeType, cfg.AMQPQueueName, cfg.AMQPQueueRouteKey)

	if err != nil {
		log.Fatal("failed to create consumer", err)
	}

	deliveries, err := consumer.Channel.Consume(
		cfg.AMQPQueueName,
		"tag",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("failed to start consuming", err)
	}

	consumer.SetupCloseHandler()
	go handler(deliveries, consumer.Done)

	log.Printf("running until Consumer is done")
	<-consumer.Done

	err = consumer.Shutdown()
	if err != nil {
		log.Fatalf("error during shutdown: %s", err)
	}
}
