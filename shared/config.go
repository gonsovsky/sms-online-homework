package shared

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sync"
)

//Config of Application
type Config struct {
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

var instantiated *Config
var once sync.Once

//AppConfig config instnce
func AppConfig() *Config {
	once.Do(func() {
		c := flag.String("c", "config.json", "Specify the configuration file.")
		flag.Parse()
		file, err := os.Open(*c)
		if err != nil {
			log.Fatal("can't open config file: ", err)
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		Config := Config{}
		err = decoder.Decode(&Config)
		if err != nil {
			log.Fatal("can't decode config JSON: ", err)
		}
		log.Println("Config.WebServer.Port: ", Config.WebServer.Port)
		log.Println("Config.Amqp.Url: ", Config.Amqp.URL)
		log.Println("Config.Db.Host: ", Config.Db.Host)
		instantiated = &Config
	})
	return instantiated
}
