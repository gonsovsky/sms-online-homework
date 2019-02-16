package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//WebServer - handle incoming web requests
type WebServer struct {
	config Configuration
	sender AmqpSender
}

//Start - Let's launch web server
func (web *WebServer) Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", web.index)
	log.Fatal(http.ListenAndServe(web.config.WebServer.Address+":"+web.config.WebServer.Port, router))
}

//Index Let's handle POST request with JSON
func (web *WebServer) index(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var msg Message
	err := decoder.Decode(&msg)
	if err != nil {
		panic(err)
	}
	log.Println("WebServer received a message: ", msg.Item)

	//Return TimeStamp to client
	msg.TimeStamp = time.Now()
	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Send to Rabbit MQ
	web.sender.Publish(msg)

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}
