package middlewares

import (
	"reflect"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos/abstractions"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	serviceerrors "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services/service_errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ReqBodyValidator[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto T
		err := c.BodyParser(&dto)
		if err != nil {
			return models.NewBadRequestError("invalid request body")
		}
		if validatable, ok := any(&dto).(abstractions.Validatable); ok {
			problemDetails := make(map[string]string)
			if err := validatable.Validate(); err != nil {
				if validationErrors, ok := err.(validator.ValidationErrors); ok {
					dtoType := reflect.TypeOf(dto)
					if dtoType.Kind() == reflect.Ptr {
						dtoType = dtoType.Elem()
					}
					for _, fieldErr := range validationErrors {
						fieldName := fieldErr.StructField()
						invalidFieldName := fieldName
						if dtoType.Kind() == reflect.Struct {
							if field, found := dtoType.FieldByName(fieldName); found {
								jsonTagValue := field.Tag.Get("json")
								if jsonTagValue != "" && jsonTagValue != "-" {
									invalidFieldName = jsonTagValue
								}
							}
						}
						problemDetails[invalidFieldName] = fieldErr.Translate(*server.AppServer.Translator)
					}
				}
				validationApiError := models.NewValidationError(serviceerrors.ErrValidationFailed.Error(), problemDetails)
				return validationApiError
			}
		}
		c.Locals("validatedDto", &dto)
		return c.Next()
	}
}
