package responses

import (
	"strings"

	"github.com/dasha-kinsely/ostruct/models/responses"
)

// If something wrong happens before this response is generated, write errors as a struct obj
func ErrorResponse(issues string, message string, data interface{}) responses.CommonResponse {
	splittedError := strings.Split(issues, "\n")
	res := responses.CommonResponse{
		Status: false,
		Message: message,
		Errors: splittedError,
		Data: data,
	}
	return res
}

// This is how a normal, data-carrying response should behave
func NormalResponse(message string, data interface{}) responses.CommonResponse {
	res := responses.CommonResponse{
		Status: true,
		Message: message,
		Errors: nil,
		Data: data,
	}
	return res
}