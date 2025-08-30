package middleware

import (
	"fmt"
	"reflect"
	"strings"
	"tasklist/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

var validate = validator.New()

type ZelebrateSegment string

const (
	ZelebrateSegmentBody    ZelebrateSegment = "body"
	ZelebrateSegmentParams  ZelebrateSegment = "params"
	ZelebrateSegmentQuery   ZelebrateSegment = "query"
	ZelebrateSegmentHeaders ZelebrateSegment = "headers"
)

const (
	ctxZelebrateRequest = "zelebrate_request" // Key used to store the validated request in the context
)

// Zelebrate is a middleware function that validates one or more of the body, params, or query of the request
// against the given struct type T.
func Zelebrate[T any](segments ...ZelebrateSegment) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		request := lo.Must(parseRequest[T](ctx, segments...))
		firstFormattedErr := ""
		errs := validate.Struct(request)
		if errs != nil {
			reflectedRequest := reflect.TypeOf(request).Elem()
			for _, err := range lo.Cast[validator.ValidationErrors](errs) {
				field, ok := reflectedRequest.FieldByName(err.StructField())
				if ok {
					messages := field.Tag.Get("messages")
					for message := range strings.SplitSeq(messages, ",") {
						messageSlice := strings.Split(message, "=")
						if len(messageSlice) == 2 && messageSlice[0] == err.Tag() {
							firstFormattedErr = messageSlice[1]
							break
						}
					}
					if firstFormattedErr == "" && messages != "" {
						firstFormattedErr = messages
					}
				}
				if firstFormattedErr == "" {
					firstFormattedErr = fmt.Sprintf("%s failed on the '%s' tag against value '%s'",
						err.Field(), err.Tag(), err.Value())
				}
			}
		}
		if firstFormattedErr != "" {
			panic(fiber.NewError(fiber.StatusUnprocessableEntity, firstFormattedErr))
		}

		ctx.Locals(ctxZelebrateRequest, request)

		return ctx.Next()
	}
}

// ZelebrateRequest extracts the validated struct of type T from the context
func ZelebrateRequest[T any](ctx *fiber.Ctx) *T {
	return ctx.Locals(ctxZelebrateRequest).(*T)
}

// Parses the specified segments of the request into a struct of type T
// thus simplifying the way that they can be accessed.
func parseRequest[T any](ctx *fiber.Ctx, segments ...ZelebrateSegment) (*T, error) {
	target := new(T)
	for _, segment := range segments {
		var err error
		switch segment {
		case ZelebrateSegmentBody:
			err = ctx.BodyParser(target)
		case ZelebrateSegmentParams:
			err = ctx.ParamsParser(target)
		case ZelebrateSegmentQuery:
			err = ctx.QueryParser(target)
		case ZelebrateSegmentHeaders:
			headers := ctx.GetReqHeaders()
			targetRef := reflect.ValueOf(target).Elem()
			targetType := targetRef.Type()
			for i := 0; i < targetRef.NumField(); i++ {
				field := targetType.Field(i)
				if headerName, ok := field.Tag.Lookup("json"); ok {
					if headerValue, exists := headers[lo.Capitalize(strings.ReplaceAll(headerName, ",omitempty", ""))]; exists {
						if targetRef.Field(i).CanSet() {
							targetRef.Field(i).Set(reflect.ValueOf(lo.LastOrEmpty(headerValue)))
						}
					}
				}
			}
		}
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}
	return target, nil
}
