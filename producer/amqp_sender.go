package producer

import (
	"log"
	"shared"

	"github.com/streadway/amqp"
)

//AmqpSender - publish requests to Rabbit service
type AmqpSender struct {
	Config  *shared.Config
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func failOnSend(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Open channel to RabbitMQ
func (sender *AmqpSender) Open() {
	conn, err := amqp.Dial(sender.Config.Amqp.URL)
	failOnSend(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	ch.Confirm(false)
	failOnSend(err, "Failed to open a channel")

	queue, err := ch.QueueDeclare(
		sender.Config.Amqp.Queue, // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	failOnSend(err, "Failed to declare a queue")

	sender.Channel = ch
	sender.Queue = &queue
}

//Publish - Send Message to RabbitMQ
func (sender *AmqpSender) Publish(msg shared.Message) {
	ch := sender.Channel
	queue := sender.Queue

	err := ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate

		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         msg.ToJSON(),
			DeliveryMode: 2, //persistent
		})
	failOnSend(err, "Failed to publish a message")

	if err != nil {
		msg.Message = "Your sms has been NOT delivered to our SMC Center. Try later."
		log.Printf("Failed to publish message to queue %s: %s", queue.Name, err)
		return
	}

	log.Println("AMQP published message: ", msg.Item)
}
