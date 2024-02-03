package database

import (
	"fmt"
	"log"
	"newsblog-users-ms/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.User{})

}
