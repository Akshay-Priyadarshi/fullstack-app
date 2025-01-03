package middlewares

import (
	"log/slog"
	"reflect"
	"strings"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos/abstractions"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/responses"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BodyValidator[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto T
		err := c.BodyParser(&dto)
		if err != nil {
			return services.New400ServiceError(err, "failed to parse request body")
		}
		if validatable, ok := any(&dto).(abstractions.Validatable); ok {
			problemDetails := make(map[string]interface{})
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
						problemDetails[invalidFieldName] = strings.ToLower(fieldErr.Translate(*server.AppServer.Translator))
						slog.Info("Logging Problem Details", "problemDetails", problemDetails)
					}
				}
				var problem responses.Problem = responses.Problem{
					Type:    responses.ProblemTypeValidation,
					Details: problemDetails,
				}
				additionalInfo := make(map[string]interface{})
				additionalInfo["problem"] = problem
				validationApiError := services.New422ServiceError(err, "validation failed", &additionalInfo)
				return validationApiError
			}
		}
		c.Locals("validatedDto", &dto)
		return c.Next()
	}
}
