package databases

import (
	"fmt"
	"strconv"

	"github.com/darot-chen/go-to-do/config"
	"github.com/darot-chen/go-to-do/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Error parsing string to int")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println("Failed to connect database")
	}

	fmt.Println("Connected to database successful")

	DB.AutoMigrate(&models.Todo{})
	fmt.Println("Database migrated")
}
