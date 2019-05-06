package shared

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"time"
)

//Message Сообщение веб-серверу...
type Message struct {
	Item            int       `json:"item"`
	Consumer        int       `json:"consumer"`
	RequestTime     time.Time `json:"requestTime"`
	ResponseTime    time.Time
	AcknowledgeTime time.Time
	Message         string
}

//ToJSON convert struct to Json byte array
func (msg *Message) ToJSON() []byte {
	output, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return output
}

//FromJSON Fill in with data from byte[]
func (msg *Message) FromJSON(bytes []byte) {
	json.Unmarshal(bytes, msg)
}

//FromReadCloser Fill in with data from Stream
func (msg *Message) FromReadCloser(r io.ReadCloser) {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(msg)
	if err != nil {
		panic(err)
	}
	log.Println("WebServer received a message: ", msg.Item)
}

//ToReader return ByteReader for this structure
func (msg *Message) ToReader() io.Reader {
	return bytes.NewBuffer(msg.ToJSON())
}
