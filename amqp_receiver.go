package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnReceive(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//StartAmqpReceiver - Let's start AMQP Receiver
func AmqpReceiver(config Configuration) {
	conn, err := amqp.Dial(config.Amqp.URL)
	failOnReceive(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnReceive(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"myqueue", // name
		false,     // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnReceive(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnReceive(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("AMQP Received a message: ", d.Body)
		}
	}()

	log.Printf(" [*] AMQP Waiting for messages.")
	<-forever
}
