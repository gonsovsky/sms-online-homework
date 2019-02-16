package main

import (
	"sync"
)

func main() {
	//Rread configuration from config.json
	var config = ReadConfiguration()

	//AMQP service to resend incoming requests from Web
	amqpSender := AmqpSender{config: config}

	//Start Web server to listen incoming requests
	webServer := WebServer{config: config, sender: amqpSender}
	go webServer.Start()

	//Subsrcibe for Amqp messages to put them to databse
	amqpReceiver := AmqpReceiver{config: config}
	go amqpReceiver.Subscribe()

	//Don't terminate main thread
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	wg.Wait()
}
