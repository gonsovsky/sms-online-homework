package main

import (
	"consumer"
	"producer"
	"shared"
	"sync"
)

func main() {
	go startProducer()
	go startConsumer()
	go startEmulator()

	var wg = &sync.WaitGroup{}
	wg.Add(3)
	wg.Wait()
}

func startProducer() {
	//AMQP service to resend incoming requests from Web
	amqpSender := producer.AmqpSender{Config: shared.AppConfig()}

	//Start Web server to listen incoming requests
	webServer := producer.WebServer{Config: shared.AppConfig(), Sender: amqpSender}
	webServer.Start()
}

func startConsumer() {
	//Atomigrate DB
	db := consumer.DataBase{Config: shared.AppConfig()}
	db.Initalize()

	//Subsrcibe for Amqp messages to put them to databse
	amqpReceiver := consumer.AmqpReceiver{Config: shared.AppConfig(), Db: db}
	amqpReceiver.Subscribe()
}

func startEmulator() {
	//Start emulation by timer with 2 seconds
	emu := Emulator{Config: shared.AppConfig()}
	emu.Start()
}
