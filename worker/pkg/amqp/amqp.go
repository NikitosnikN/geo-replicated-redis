package amqp

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Consumer struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	tag     string
	Done    chan error
}

func NewConsumer(connectionString string, exchangeName string, exchangeType string, queueName string, queueRouteKey string) (*Consumer, error) {
	c := &Consumer{
		Conn:    nil,
		Channel: nil,
		Done:    make(chan error),
	}
	var err error

	c.Conn, err = newConnection(connectionString)

	if err != nil {
		return nil, fmt.Errorf("newConnection: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.Conn.NotifyClose(make(chan *amqp.Error)))
	}()

	c.Channel, err = newChannel(c.Conn)

	if err != nil {
		return nil, fmt.Errorf("newChannel: %s", err)
	}

	err = initExchangeAndQueue(c.Channel, exchangeName, exchangeType, queueName, queueRouteKey)

	if err != nil {
		return nil, fmt.Errorf("initExchangeAndQueue: %s", err)
	}

	return c, nil
}

func newConnection(connectionString string) (*amqp.Connection, error) {
	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName("consumer")
	return amqp.DialConfig(connectionString, config)
}

func newChannel(connection *amqp.Connection) (*amqp.Channel, error) {
	ch, err := connection.Channel()

	if err != nil {
		return nil, fmt.Errorf("channel: %s", err)
	}

	return ch, nil
}

func initExchangeAndQueue(channel *amqp.Channel, exchangeName string, exchangeType string, queueName string, queueRouteKey string) error {
	err := channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("exchange Declare: %s", err)
	}

	queue, err := channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)

	if err != nil {
		return fmt.Errorf("queue Declare: %s", err)
	}

	err = channel.QueueBind(
		queue.Name,
		queueRouteKey,
		exchangeName,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("queue Bind %s", err)
	}

	return nil
}

func (c *Consumer) Shutdown() error {
	if err := c.Channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("consumer cancel failed: %s", err)
	}

	if err := c.Conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer fmt.Printf("AMQP shutdown OK")

	return <-c.Done
}

func (c *Consumer) SetupCloseHandler() {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		log.Printf("Ctrl+C pressed in Terminal")
		if err := c.Shutdown(); err != nil {
			log.Fatalf("error during shutdown: %s", err)
		}
		os.Exit(0)
	}()
}
