package broker

import (
	"bitcoin_checker_api/config"
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMQ struct {
	connection   *amqp.Connection
	Chanel       *amqp.Channel
	BrokerName   string
	RoutingKey   string
	LogsExchange string
	LogsQueue    string
}

type RabbitMQService struct {
	connection Connection
	routingKey string
	exchange   string
	queue      string
}

const (
	RabbitmqBrokerName   = "rabbitmq"
	ErrRoutingKeyPostfix = "-error"
)

func NewRabbitMQ(cfg config.RabbitMQ) (*RabbitMQ, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel %w", err)
	}
	return &RabbitMQ{
		connection:   conn,
		Chanel:       ch,
		BrokerName:   RabbitmqBrokerName,
		RoutingKey:   cfg.RoutingKey,
		LogsExchange: cfg.LogsExchange,
		LogsQueue:    cfg.LogsQueue}, nil
}

func NewRabbitMQService(rbmq Connection, e, q, r string) (*RabbitMQService, error) {
	if err := initExchangeAndQueue(rbmq.(*RabbitMQ), e, q, r); err != nil {
		return nil, err
	}
	return &RabbitMQService{
		connection: rbmq,
		exchange:   e,
		queue:      q,
		routingKey: r,
	}, nil
}

func (that *RabbitMQService) Send(message []byte) error {
	return that.connection.Send(message, that.exchange, that.routingKey)
}
func (that *RabbitMQService) SendErr(message []byte) error {
	return that.connection.Send(message, that.exchange, that.routingKey+ErrRoutingKeyPostfix)
}

func (that *RabbitMQService) Receive() error {
	return nil
}

func initExchangeAndQueue(rmbq *RabbitMQ, exchangeName, queueName, routingKey string) error {
	err := rmbq.Chanel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // kind
		false,        // durable
		false,        // autoDelete
		false,        // internal
		false,        // noWait,
		nil,          //args
	)
	if err != nil {
		return nil
	}

	_, err = rmbq.Chanel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	qerror, err := rmbq.Chanel.QueueDeclare(
		queueName+ErrRoutingKeyPostfix, // name
		false,                          // durable
		false,                          // delete when unused
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)
	if err != nil {
		return err
	}

	err = rmbq.Chanel.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if err != nil {
		return err
	}

	err = rmbq.Chanel.QueueBind(qerror.Name, routingKey+ErrRoutingKeyPostfix, exchangeName, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (that *RabbitMQ) Send(message []byte, exchangeName, routingKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := that.Chanel.PublishWithContext(ctx,
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})

	if err != nil {
		return err
	}
	return nil
}

func (that *RabbitMQ) Receive() error {
	return nil
}
