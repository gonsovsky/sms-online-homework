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

	//Atomigrate DB
	db := DataBase{config: config}
	db.Initalize()

	//Subsrcibe for Amqp messages to put them to databse
	amqpReceiver := AmqpReceiver{config: config, db: db}
	go amqpReceiver.Subscribe()

	//Start emulation by timer with 2 seconds
	pusher := Pusher{config: config}
	go pusher.Start()

	//Don't terminate main thread
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	wg.Wait()
}
