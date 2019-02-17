package main

import (
	"consumer"
	"producer"
	"shared"
	"sync"
	"tester"
)

func main() {
	startProducer()
	startConsumer(10)

	startTester()
	dontExit()
}

func startProducer() {
	//AMQP service to resend incoming requests from Web
	amqpSender := producer.AmqpSender{Config: shared.AppConfig()}
	amqpSender.Open()

	//Start Web server to listen incoming requests
	webServer := producer.WebServer{Config: shared.AppConfig(), Sender: amqpSender}
	go webServer.Start()
}

func startConsumer(count int) {
	//Atomigrate DB
	db := consumer.DataBase{Config: shared.AppConfig()}
	db.Initalize()

	for i := 1; i <= count; i++ {
		//Subsrcibe for Amqp messages to put them to databse
		amqpReceiver := consumer.AmqpReceiver{Config: shared.AppConfig(), Db: &db, No: i}
		go amqpReceiver.Subscribe()
	}
}

func startTester() {
	//Start emulation 1000 messages
	db := consumer.DataBase{Config: shared.AppConfig()}
	db.Open()
	tester := tester.Tester{Config: shared.AppConfig(), Db: &db}
	go tester.Start()
	go tester.WaitReady()
}

func dontExit() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
