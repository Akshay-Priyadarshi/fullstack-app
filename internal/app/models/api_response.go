package models

type ApiResponse[Data interface{}] struct {
	Success        bool                    `json:"success"`
	Message        string                  `json:"message"`
	Data           *Data                   `json:"data"` // Pointer to make it nullable
	StatusCode     int                     `json:"statusCode"`
	AdditionalInfo *map[string]interface{} `json:"additionalInfo"` // Pointer to make it nullable
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
