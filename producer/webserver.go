package producer

import (
	"log"
	"net/http"
	"shared"
	"time"

	"github.com/gorilla/mux"
)

//WebServer - handle incoming web requests
type WebServer struct {
	Config *shared.Config
	Sender AmqpSender
}

//Start - launch web server
func (web *WebServer) Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", web.index)
	log.Fatal(http.ListenAndServe(web.Config.WebServer.Address+":"+web.Config.WebServer.Port, router))
}

//Index - handle POST request with JSON
func (web *WebServer) index(w http.ResponseWriter, r *http.Request) {
	var msg shared.Message
	msg.FromReadCloser(r.Body)
	msg.TimeStamp = time.Now()
	web.Sender.Publish(msg)
	w.Header().Set("content-type", "application/json")
	w.Write(msg.ToJSON())

}
