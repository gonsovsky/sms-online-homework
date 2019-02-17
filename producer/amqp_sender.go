package producer

import (
	"log"
	"shared"

	"github.com/streadway/amqp"
)

//AmqpSender - publish requests to Rabbit service from web
type AmqpSender struct {
	Config  *shared.Config
	channel *amqp.Channel
	conn    *amqp.Connection
	queue   amqp.Queue
	err     error
}

func failOnSend(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Publish - Send Message to RabbitMQ
func (sender *AmqpSender) Publish(msg shared.Message) {

	sender.conn, sender.err = amqp.Dial(sender.Config.Amqp.URL)
	failOnSend(sender.err, "Failed to connect to RabbitMQ")
	defer sender.conn.Close()

	sender.channel, sender.err = sender.conn.Channel()
	failOnSend(sender.err, "Failed to open a channel")
	defer sender.channel.Close()

	sender.queue, sender.err = sender.channel.QueueDeclare(
		sender.Config.Amqp.Queue, // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	failOnSend(sender.err, "Failed to declare a queue")

	sender.err = sender.channel.Publish(
		"",                // exchange
		sender.queue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg.ToJSON(),
		})
	failOnSend(sender.err, "Failed to publish a message")

	log.Println("AMQP published message: ", msg.Item)
}