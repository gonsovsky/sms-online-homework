package main

import (
	"sync"
)

func main() {
	//Let's read configuration
	var config = ReadConfiguration()

	//Connect to AMQP service to resend incoming requests from Web
	amqpSender := AmqpSender{config: config}
	amqpSender.Connect()

	//Let's start Web server to listen incoming requests
	webServer := WebServer{config: config, sender: amqpSender}
	go webServer.Start()

	//Let's start Amqp Receiver to listen incoming requests
	go AmqpReceiver(config)

	var wg = &sync.WaitGroup{}
	wg.Add(2)
	wg.Wait()
}
