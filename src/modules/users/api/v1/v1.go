package v1

import (
	m "tasklist/src/middleware"
	sdto "tasklist/src/modules/_shared/dto"
	"tasklist/src/modules/users/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	router.Post("/", m.Zelebrate[dto.CreateUserRequest](m.ZelebrateSegmentBody), CreateUserHandler)
	router.Get("/", m.Zelebrate[sdto.PaginatedQuery](m.ZelebrateSegmentQuery),
		m.FilterQuery, GetUsersHandler)
	router.Get("/:id", m.Zelebrate[dto.GetUserRequest](m.ZelebrateSegmentParams), GetUserHandler)
	router.Patch("/:id", m.Zelebrate[dto.UpdateUserRequest](m.ZelebrateSegmentParams, m.ZelebrateSegmentBody),
		UpdateUserHandler)
	router.Delete("/:id", m.Zelebrate[dto.DeleteUserRequest](m.ZelebrateSegmentParams), DeleteUserHandler)
}
