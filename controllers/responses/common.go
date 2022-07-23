package responses

import (
	"strings"
)

type CommonResponse struct{
	Status bool `json:"status"`
	Message string `json:"message"`
	Errors interface{} `json:"errors"`
	Data interface{} `json:"data"`
}

// If something wrong happens before this response is generated, write errors as a struct obj
func ErrorResponse(issues string, message string, data interface{}) CommonResponse {
	splittedError := strings.Split(issues, "\n")
	res := CommonResponse{
		Status: false,
		Message: message,
		Errors: splittedError,
		Data: data,
	}
	return res
}

// This is how a normal, data-carrying response should behave
func NormalResponse(message string, data interface{}) CommonResponse {
	res := CommonResponse{
		Status: true,
		Message: message,
		Errors: nil,
		Data: data,
	}
	return res
}