package middleware

import (
	"fmt"
	"reflect"
	"strings"
	"tasklist/src/pkg/validator"

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

// Zelebrate is a middleware function that validates one or more of the body, params, or query of the request
// against the given struct type T.
func Zelebrate[T any](segments ...ZelebrateSegment) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		target := new(T)

		targetValue := reflect.ValueOf(target).Elem()

		for _, segment := range segments {
			segmentTarget := new(T)
			var err error
			switch segment {
			case ZelebrateSegmentBody:
				err = ctx.BodyParser(segmentTarget)
			case ZelebrateSegmentParams:
				err = ctx.ParamsParser(segmentTarget)
			case ZelebrateSegmentQuery:
				err = ctx.QueryParser(segmentTarget)
			case ZelebrateSegmentHeaders:
				*segmentTarget = lo.CastJSON[T](ctx.GetReqHeaders())
			}
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			if len(segments) == 1 {
				target = segmentTarget
			} else {
				segmentValue := reflect.ValueOf(segmentTarget).Elem()
				for i := 0; i < segmentValue.NumField(); i++ {
					field := segmentValue.Field(i)
					if field.IsZero() {
						continue
					}
					targetField := targetValue.Field(i)
					if targetField.CanSet() {
						targetField.Set(field)
					}
				}
			}
		}

		firstFormattedErr := ""
		errs := validate.Struct(target)
		if errs != nil {
			reflectedTarget := reflect.TypeOf(target).Elem()
			for _, err := range lo.Cast[validator.ValidationErrors](errs) {
				field, ok := reflectedTarget.FieldByName(err.StructField())
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
		return ctx.Next()
	}
}
