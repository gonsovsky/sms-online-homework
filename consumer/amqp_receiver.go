package consumer

import (
	"log"
	"shared"

	"github.com/streadway/amqp"
)

//AmqpReceiver - subsribes for events
type AmqpReceiver struct {
	Config *shared.Config
	Db     DataBase
}

//Subscribe - Let's launch web server
func (receiver *AmqpReceiver) Subscribe() {
	conn, err := amqp.Dial(receiver.Config.Amqp.URL)
	failOnReceive(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnReceive(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		receiver.Config.Amqp.Queue, // name
		false,                      // durable
		false,                      // delete when usused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
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
			msg := shared.Message{}
			msg.FromJSON(d.Body)
			log.Println("AMQP Received a message: ", msg.Item)
			receiver.Db.Post(msg)
		}
	}()

	log.Printf(" [*] AMQP Waiting for messages.")
	<-forever
}

func failOnReceive(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
