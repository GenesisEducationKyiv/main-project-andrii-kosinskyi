package main

import (
	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/pkg/broker"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatalf("cfg.Load: %s\n", err)
	}

	cfg.RabbitMQ.Host = "localhost"
	rabbitmqConn, err := broker.NewRabbitMQ(cfg.RabbitMQ)
	if err != nil {
		failOnError(err, "Failed to create connection")
	}

	q, err := rabbitmqConn.Chanel.QueueDeclare(
		cfg.RabbitMQ.LogsQueue, // name
		false,                  // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := rabbitmqConn.Chanel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
