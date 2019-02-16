package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DataBase - DataBase
type DataBase struct {
	config Configuration
	db     *gorm.DB
}

//Initalize - Initialize DB
func (database *DataBase) Initalize() {
	constr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		database.config.Db.Host, database.config.Db.Port,
		database.config.Db.User, database.config.Db.Database,
		database.config.Db.Password)
	db, err := gorm.Open("postgres", constr)
	if err != nil {
		panic(err)
	}
	db.DropTableIfExists(&Message{})
	db.AutoMigrate(&Message{})
	database.db = db
	log.Println("Databse initialized and truncated")
}

//Post - Post message to databse
func (database *DataBase) Post(msg Message) {
	database.db.Create(msg)
	log.Println("Databse wrote message: ", msg.Item)
}
