package database

import (
	"log"

	"github.com/adityasunny1189/jwt-demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) () {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("cannot connect to DB")
	}
	log.Println("connected to DB")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database migration completed")
}