package main

import (
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

//Start - launch web server
func (web *WebServer) Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", web.index)
	log.Fatal(http.ListenAndServe(web.config.WebServer.Address+":"+web.config.WebServer.Port, router))
}

//Index - handle POST request with JSON
func (web *WebServer) index(w http.ResponseWriter, r *http.Request) {
	var msg Message
	msg.FromReadCloser(r.Body)
	msg.TimeStamp = time.Now()
	web.sender.Publish(msg)
	w.Header().Set("content-type", "application/json")
	w.Write(msg.ToJSON())

}
