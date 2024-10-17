package cerror

import "encoding/json"

type customError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *customError) Error() string {
	bytes, err := json.Marshal(&c)
	if err != nil {
		return "marshal error"
	}

	return string(bytes)
}

func New(code, message string) *customError {
	return &customError{
		Code:    code,
		Message: message,
	}
}

func BadRequest() *customError {
	return &customError{
		Code:    BAD_REQUEST,
		Message: "invalid input data",
	}
}

func NotFound() *customError {
	return &customError{
		Code:    NOT_FOUND,
		Message: "not found",
	}
}

func InternalServerError() *customError {
	return &customError{
		Code:    INTERNAL_SERVER_ERROR,
		Message: "Internal Server Error",
	}
}

func Forbidden() *customError {
	return &customError{
		Code:    FORBIDDEN,
		Message: "Forbidden",
	}
}
