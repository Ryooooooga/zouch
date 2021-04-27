package errors

import "fmt"

type templateExistError struct {
	message string
}

func TemplateExistError(format string, args ...interface{}) error {
	return &templateExistError{
		message: fmt.Sprintf(format, args...),
	}
}
func (err *templateExistError) Error() string {
	return err.message
}

func IsTemplateExistError(err error) bool {
	_, ok := err.(*templateExistError)
	return ok
}
