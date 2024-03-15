package errs

import "fmt"

type CustomError struct {
	Type    string
	Code    string
	Message string
}

func (ce CustomError) Error() string {
	return fmt.Sprintf("Type: %s, Code: %s, Message: %s", ce.Type, ce.Code, ce.Message)
}

var ErrEmptyName = CustomError{
	Type:    "validation",
	Code:    "EMPTY_NAME",
	Message: "name wasn't provided",
}
