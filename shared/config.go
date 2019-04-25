package shared

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

//Config - настройки программы
type Cfg struct {
	Participants int    //кол-во участников
	Redis RedisCfg
}

//RedisConfig - настройки программы, Redis
type RedisCfg struct {
	Host     string
	Port     string
	Db       int
	Key      string
}

var instantiated *Cfg
var once sync.Once

//AppConfig - настройки программы
func AppConfig() *Cfg {
	once.Do(func() {
		c := flag.String("c", "config.json", "Specify the configuration file.")
		flag.Parse()
		file, err := os.Open(*c)
		if err != nil {
			log.Fatal("can't open config file: ", err)
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		Config := Cfg{}
		err = decoder.Decode(&Config)
		if err != nil {
			log.Fatal("can't decode config JSON: ", err)
		}
		Pretty(Config)
		instantiated = &Config
	})
	return instantiated
}

func RedisConfig() *RedisCfg {
	return &AppConfig().Redis
}

//HostAndPort - Узел и порт Редис
func (c *RedisCfg) HostAndPort() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)

}
