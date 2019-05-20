package ytrwrap

import "fmt"

type APICode int

const (
	OK                      = APICode(200)
	KEY_WRONG               = APICode(403)
	KEY_BLOCKED             = APICode(402)
	LIMIT_DAILY_EXCEEDED    = APICode(404)
	LIMIT_TEXTSIZE_EXCEEDED = APICode(413)
	CANNOT_TRANSLATE        = APICode(422)
	NOT_SUPPORTED_DIRECTION = APICode(501)
	WRAPPER_INTERNAL_ERROR  = APICode(500)
)

type GenericResponse struct {
	Description string  `json:"message,omitempty"`
	ErrorCode   APICode `json:"code,omitempty"`
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

func Code(raw interface{}) APICode {
	err := raw.(*apiError)
	return err.ErrorCode
}

func Description(raw interface{}) string {
	err := raw.(*apiError)
	return err.Description
}
