package configs

import (
	"auth_service/models"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "daman12345"
	dbname   = "auth_service"
)

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed Connect to database: ", err)
		return
	}

	fmt.Println("DATABASE CONNECTED")

	DB.AutoMigrate(models.Admin{})
}
