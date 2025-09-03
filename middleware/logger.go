package middleware

import (
	"tasklist/global"
	. "tasklist/modules/users/api/v1/models"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// List of paths which will be ignored by the logger.
var ZapWhitelists = []string{
	"/system/health",
	"/system/liveness",
	"/system/readiness",
}

// getCustomLogFields extracts custom fields from the request context for logging purposes.
func getCustomLogFields(c *fiber.Ctx) []any {
	headers := c.GetReqHeaders()
	fields := []any{
		"ip", c.IP(),
		"status", c.Response().StatusCode(),
		"method", c.Method(),
		"path", c.Path(),
		"user-agent", lo.FirstOrEmpty(headers[fiber.HeaderUserAgent]),
		"query", c.Queries(),
	}
	user, ok := c.Locals(global.CtxUser).(*User)
	if ok && user != nil {
		fields = append(fields, "user-id", user.ID.Hex())
	}
	return fields
}

// getZapLogFields converts custom log fields into zap.Field format for structured logging.
func getZapLogFields(c *fiber.Ctx) []zap.Field {
	fields := []zap.Field{
		zap.String(global.CtxCorrelationID, lo.Cast[string](c.Locals(global.CtxCorrelationID))),
	}
	customFields := getCustomLogFields(c)
	for i := 0; i < len(customFields); i += 2 {
		key := lo.Cast[string](customFields[i])
		fields = append(fields, zap.Any(key, customFields[i+1]))
	}
	return fields
}

// getLogFields prepares log fields in a key-value format for non-zap loggers.
func getLogFields(c *fiber.Ctx) []any {
	fields := getCustomLogFields(c)
	payload := string(c.Body())
	if len(payload) > 0 {
		fields = append(fields, "payload", payload)
	}
	return fields
}

// Zapped is a middleware that overrides the default logger with zapcore and sets up an http request logger.
// The logs are populated with the correlation ID associated with the request.
func Zapped(c *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	log.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		ZapOptions: []zap.Option{
			zap.Fields(getZapLogFields(c)...),
		},
	}))

	path := c.Path()

	if lo.Contains(ZapWhitelists, path) {
		return c.Next()
	}

	log.Infow("Request initiated", getLogFields(c)...)

	return fiberzap.New(fiberzap.Config{
		Logger: logger,
		FieldsFunc: func(c *fiber.Ctx) []zap.Field {
			return getZapLogFields(c)
		},
		Fields:   []string{"latency"},
		Messages: []string{"Server error", "Client error", "Request completed"},
	})(c)
}

// fiberzapPostRecoveryLog is a temporary solution invoked at the default `errorHandler`
// to log the status of failed http requests since the panic + recover flow we use
// doesn't trigger the fiberzap logger on request completion.
func fiberzapPostRecoveryLog(c *fiber.Ctx) {
	log.Errorw("Request failed", getLogFields(c)...)
}
