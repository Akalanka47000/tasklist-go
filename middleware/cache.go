package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

// Caches successful responses for 30 minutes. Please refer the fiber cache documentation
// for more details: https://docs.gofiber.io/api/middleware/cache
// Important: This middleware respects the Cache-Control header.
// If the Cache-Control header is set to no-cache, the cache will not be used.
var CacheSuccess = cache.New(cache.Config{
	KeyGenerator: func(ctx *fiber.Ctx) string {
		return ctx.OriginalURL()
	},
	Expiration: time.Minute * 30,
})
