package worker

import (
	"worker/internal/config"
	"worker/internal/providers"
	"worker/pkg/amqp"
	"worker/pkg/redis"

	"log"
)

func Consume(cfg config.Config) {
	consumer, err := amqp.NewConsumer(cfg.AMQPConnectionString, cfg.AMQPExchangeName, cfg.AMQPExchangeType, cfg.AMQPQueueName, cfg.AMQPQueueRouteKey)

	if err != nil {
		log.Fatal("failed to create consumer", err)
	}

	redisClient, err := redis.NewClient(cfg.RedisConnectionString)

	if err != nil {
		log.Fatal("failed to create redis provider", err)
	}

	redisProvider := &providers.RedisProvider{
		Client: redisClient,
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
	go handler(redisProvider, deliveries, consumer.Done)

	log.Printf("running until Consumer is done")
	<-consumer.Done

	err = consumer.Shutdown()
	if err != nil {
		log.Fatalf("error during shutdown: %s", err)
	}
}
