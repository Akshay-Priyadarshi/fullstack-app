package responses

import (
	"github.com/gofiber/fiber/v2"
)

func NewOkResponse[Data interface{}](message string, data *Data) ApiResponse[Data] {
	response := NewResponse(message, fiber.StatusOK, data, &map[string]interface{}{})
	return response
}
