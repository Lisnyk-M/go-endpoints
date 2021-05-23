package api

import (
	"fmt"
	"os"

	"github.com/Lisnyk-M/go-endpoints/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

var DB *gorm.DB

func Db_connection() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	postgresql_string, exists := os.LookupEnv("POSTGRESQL_CONNECTION")

	if exists {
		fmt.Println(postgresql_string)
	} else {
		fmt.Println(".env not exist")
	}

	pgUrl, err := pq.ParseURL(postgresql_string)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db, err := gorm.Open("postgres", pgUrl)

	if err != nil {
		panic("Failed to connect to database! 222222222222222")
	}
	// return db
	db.AutoMigrate(&models.User{})
	DB = db
}
