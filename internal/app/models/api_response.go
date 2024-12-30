package models

import "github.com/gofiber/fiber/v2"

type ApiResponse[Data interface{}] struct {
	Success        bool                    `json:"success"`
	Message        string                  `json:"message"`
	Data           *Data                   `json:"data"`
	StatusCode     int                     `json:"statusCode"`
	AdditionalInfo *map[string]interface{} `json:"additionalInfo"`
}

func NewApiResponse[Data interface{}](message string, statusCode int, data *Data, additionalInfo *map[string]interface{}) ApiResponse[Data] {
	if additionalInfo == nil {
		emptyMap := make(map[string]interface{})
		additionalInfo = &emptyMap
	}
	return ApiResponse[Data]{
		Success:        statusCode >= 200 && statusCode < 300,
		Message:        message,
		Data:           data,
		StatusCode:     statusCode,
		AdditionalInfo: additionalInfo,
	}
}

func NewOkApiResponse[Data interface{}](message string, data *Data) ApiResponse[Data] {
	return NewApiResponse(message, fiber.StatusOK, data, &map[string]interface{}{})
}
