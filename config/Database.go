package config

import (
	"BlogAPI-GO/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectToDB(){
	fmt.Print("running")
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println("err")
	}
	User := models.User{}
	DB.AutoMigrate(&User)
	
}