package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateBody[T any]() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body T

		if err := ctx.BodyParser(&body); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request format.",
			})
		}

		if err := validate.Struct(body); err != nil {
			
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Validation failed.",
				"details": validationErrorsToJSON(err),
			})
		}

		ctx.Locals("body",body)


		return ctx.Next()
	}
}


func validationErrorsToJSON(err error) map[string]string {
	errorsMap := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		errorsMap[e.Field()] = e.Error()
	}

	return errorsMap
}