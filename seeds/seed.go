package seeds

import (
	"auth_service/configs"
	"auth_service/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {

	var existingAdmin models.Admin
	if err := configs.DB.Where("username = ?", "admin").First(&existingAdmin).Error; err == nil {
		fmt.Println("Admin already exists, skipping seed.")
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("admin12345"), bcrypt.DefaultCost)
	admin := models.Admin{
		Username: "admin",
		Password: string(password),
	}

	configs.DB.Create(&admin)
}
