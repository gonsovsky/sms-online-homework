package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//Pusher - Emaulate requests to Web Server on time
type Pusher struct {
	config Configuration
}

//Start - launch emulator
func (pusher *Pusher) Start() {
	var n = 0
	for true {
		n++
		time.Sleep(2 * time.Second)
		var msg = Message{Item: n}
		pusher.post(&msg)
	}
}

func (pusher *Pusher) serviceURL() string {
	return "http://" + pusher.config.WebServer.Address + ":" + pusher.config.WebServer.Port + "/"
}

func (pusher *Pusher) post(msg *Message) string {
	req, err := http.NewRequest("POST", pusher.serviceURL(), msg.ToReader())
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Status:", resp.Status, "; response Body:", string(body))
	return resp.Status
}
