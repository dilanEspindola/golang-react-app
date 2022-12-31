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
	users := []models.User{}

	if err := config.Database.Db.Find(&users); err == nil {
		fmt.Println(err.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "INTERNAL_SERVER_ERROR"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}
	errors := helpers.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if result := config.Database.Db.Create(&user); result.Error != nil {
		var errorKind = result.Error.Error()
		value := helpers.ErrorValidations(errorKind, user.Email)
		if len(value) > 0 {
			return c.Status(fiber.StatusConflict).JSON(map[string]string{"message": value})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "INTERNAL_SERVER_ERROR"})
	}

	response := Response{
		Message: "USER_CREATED",
		Data:    user,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
