package consumer

import (
	"fmt"
	"log"
	"shared"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DataBase - DataBase
type DataBase struct {
	Config *shared.Config
	db     *gorm.DB
}

//Initalize - Initialize DB
func (database *DataBase) Initalize() {
	constr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		database.Config.Db.Host, database.Config.Db.Port,
		database.Config.Db.User, database.Config.Db.Database,
		database.Config.Db.Password)
	db, err := gorm.Open("postgres", constr)
	if err != nil {
		panic(err)
	}
	db.DropTableIfExists(&shared.Message{})
	db.AutoMigrate(&shared.Message{})
	database.db = db
	log.Println("Databse initialized and truncated")
}

//Post - Post message to databse
func (database *DataBase) Post(msg shared.Message) {
	database.db.Create(msg)
	log.Println("Databse wrote message: ", msg.Item)
}
