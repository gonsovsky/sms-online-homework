package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

//Configuration of Application
type Configuration struct {
	WebServer struct {
		Address string
		Port    string
	}
	Amqp struct {
		URL   string
		Queue string
	}
	Db struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

//ReadConfiguration - read config.json
func ReadConfiguration() Configuration {
	c := flag.String("c", "config.json", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	log.Println("Config.WebServer.Port: ", Config.WebServer.Port)
	log.Println("Config.Amqp.Url: ", Config.Amqp.URL)
	log.Println("Config.Db.Host: ", Config.Db.Host)
	return Config
}
