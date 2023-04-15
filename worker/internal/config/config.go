package config

type Config struct {
	AMQPConnectionString string `envconfig:"AMQP_CONNECTION_STRING" required:"true"`
	AMQPExchangeName     string `envconfig:"AMQP_EXCHANGE_NAME" required:"true"`
	AMQPExchangeType     string `envconfig:"AMQP_EXCHANGE_TYPE" required:"true"`
	AMQPQueueName        string `envconfig:"AMQP_QUEUE_NAME" required:"true"`
	AMQPQueueRouteKey    string `envconfig:"AMQP_QUEUE_ROUTE_KEY" required:"true"`
}
