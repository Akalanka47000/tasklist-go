package v1

import (
	"github.com/gofiber/fiber/v2"
	m "tasklist/src/middleware"
	sdto "tasklist/src/modules/_shared/dto"
	"tasklist/src/modules/users/api/v1/dto"
)

func New() *fiber.App {
	v1 := fiber.New()
	v1.Post("/", m.Zelebrate[dto.CreateUserRequest](m.ZelebrateSegmentBody), CreateUserHandler)
	v1.Get("/", m.Zelebrate[sdto.PaginatedQuery](m.ZelebrateSegmentQuery),
		m.FilterQuery, GetUsersHandler)
	v1.Get("/:id", m.Zelebrate[dto.GetUserRequest](m.ZelebrateSegmentParams), GetUserHandler)
	v1.Patch("/:id", m.Zelebrate[dto.UpdateUserRequest](m.ZelebrateSegmentParams, m.ZelebrateSegmentBody),
		UpdateUserHandler)
	v1.Delete("/:id", m.Zelebrate[dto.DeleteUserRequest](m.ZelebrateSegmentParams), DeleteUserHandler)
	return v1
}
