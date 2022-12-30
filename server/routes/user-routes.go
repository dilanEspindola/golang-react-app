package routes

import (
	"anime-hentai-backend/config"
	"anime-hentai-backend/helpers"
	"anime-hentai-backend/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "users"})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	errors := helpers.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := config.Database.Db.Create(&user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "INTERNAL_SERVER_ERROR"})
	}

	response := Response{
		Message: "USER_CREATED",
		Data:    user,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
