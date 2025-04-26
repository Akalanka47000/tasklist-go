package v1

import (
	m "tasklist/src/middleware"
	"tasklist/src/modules/users/api/v1/dto"
	v1 "tasklist/src/modules/users/api/v1/models"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1.SyncIndexes()
	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreateUserReq](m.Body), Create)
	return v1
}
