package errors

import "github.com/gofiber/fiber/v2"

var (
	UserNotFound = fiber.NewError(fiber.StatusNotFound, "User not found")
)
