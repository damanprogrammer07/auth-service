package controllers

import (
	"auth_service/configs"
	"auth_service/models"
	"auth_service/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {
	var input Request
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input format")
	}

	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Username and password are required")
	}

	var admin models.Admin
	if err := configs.DB.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := utils.GenerateJWT(admin.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error generating token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":    token,
		"admin_id": admin.Id,
		"username": admin.Username,
	})
}
