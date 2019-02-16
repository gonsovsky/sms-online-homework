package main

import (
	"time"
)

//Message Сообщение веб-серверу...
type Message struct {
	Item      int       `json:"item"`
	TimeStamp time.Time `json:"timestamp"`
}
