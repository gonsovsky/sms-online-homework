package tester

import (
	"consumer"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"shared"
	"time"
)

const (
	toSend int = 1000
)

//Tester ...
type Tester struct {
	Config *shared.Config
	Db     *consumer.DataBase
}

//Start send fake messages to web server
func (e *Tester) Start() {
	time.Sleep(1 * time.Second)
	for i := 1; i <= toSend; i++ {
		var msg = shared.Message{Item: i}
		e.post(&msg)
	}
}

func (e *Tester) serviceURL() string {
	url := url.URL{
		Scheme: "http",
		Host:   e.Config.WebServer.Address + ":" + e.Config.WebServer.Port,
	}
	return url.String()
}

// Send fake message to web host
func (e *Tester) post(msg *shared.Message) string {
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

//WaitReady - wait finish
func (e *Tester) WaitReady() {
	for {
		time.Sleep(1 * time.Second)
		if e.Db.Count() == toSend {
			log.Println("Ready.")
			os.Exit(0)
		}
	}
}
