package handlers

type BaseResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Errors     any    `json:"errors"`
}

func SuccessResponse(data any, message *string) *BaseResponse {
	var successMessage string

	if message == nil {
		successMessage = "Success"
	} else {
		successMessage = *message
	}

	return &BaseResponse{
		StatusCode: 200,
		Data:       data,
		Message:    successMessage,
	}
}

func ErrorResponse(message string, statusCode int) *BaseResponse {
	return &BaseResponse{
		StatusCode: statusCode,
		Data:       nil,
		Message:    message,
	}
}

func ErrorValidateResponse(errors any) *BaseResponse {
	return &BaseResponse{
		StatusCode: 400,
		Data:       nil,
		Message:    "validation errors",
		Errors:     errors,
	}
}
