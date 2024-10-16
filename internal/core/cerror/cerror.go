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
