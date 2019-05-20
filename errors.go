package ytrwrap

import (
	"fmt"
)

//
// APICode represents HTTP code w/ special meaning
//
type APICode int

const (
	//
	// OK is http.StatusOK equivalent
	//
	OK = APICode(200)

	//
	// KEY_WRONG indicated invalid api key
	//
	KEY_WRONG = APICode(403)

	//
	// KEY_BLOCKED for state when key is valid but not allowed
	//
	KEY_BLOCKED = APICode(402)

	//
	// LIMIT_DAILY_EXCEEDED indicates when amount of chars
	// for 24h exceeds limit
	//
	LIMIT_DAILY_EXCEEDED = APICode(404)

	//
	// LIMIT_TEXTSIZE_EXCEEDED indicates
	// current request is too big
	//
	LIMIT_TEXTSIZE_EXCEEDED = APICode(413)

	//
	// CANNOT_TRANSLATE - arguments are valid but text is not translateable
	//
	CANNOT_TRANSLATE = APICode(422)

	//
	// NOT_SUPPORTED_DIRECTION for invalid source/destination combination
	//
	NOT_SUPPORTED_DIRECTION = APICode(501)

	//
	// WRAPPER_INTERNAL_ERROR for all the cases when library
	// fails to send network request, decode result etc
	//
	WRAPPER_INTERNAL_ERROR = APICode(500)
)

//
// GenericResponse is common part of all the service responses
//
type GenericResponse struct {
	//
	// Description is a comment about code
	//
	Description string `json:"message,omitempty"`

	//
	// ErrorCode - numeric code indicates if operation successful
	//
	ErrorCode APICode `json:"code,omitempty"`
}

type apiError struct {
	GenericResponse
}

func newError(resp GenericResponse) *apiError {
	return &apiError{
		GenericResponse: resp,
	}
}

func apiErrorf(format string, a ...interface{}) *apiError {
	return newError(
		GenericResponse{
			ErrorCode:   WRAPPER_INTERNAL_ERROR,
			Description: fmt.Sprintf(format, a...),
		})
}

func (p *apiError) Error() string {
	return fmt.Sprintf("%v | %v", p.ErrorCode, p.Description)
}
