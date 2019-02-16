package main

//Pusher - Emaulate requests to Web Server on time
type Pusher struct {
	config Configuration
}

//Start - launch emulator
func (pusher *Pusher) Start() {
	for true {
	{
		time.Sleep(1)
		pusher.getURL()
	}
}

func (pusher *Pusher) getURL() {
	return "http://" + pusher.config.WebServer.Address + ":" + pusher.config.WebServer.Port + "/"
}
