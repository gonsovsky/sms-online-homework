package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"shared"
	"time"
)

//Emulator ...
type Emulator struct {
	Config *shared.Config
}

//Start emulator
func (e *Emulator) Start() {
	var n = 0
	for true {
		n++
		time.Sleep(2 * time.Second)
		var msg = shared.Message{Item: n}
		e.post(&msg)
	}
}

func (e *Emulator) serviceURL() string {
	return "http://" + e.Config.WebServer.Address + ":" + e.Config.WebServer.Port + "/"
}

func (e *Emulator) post(msg *shared.Message) string {
	req, err := http.NewRequest("POST", e.serviceURL(), msg.ToReader())
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
