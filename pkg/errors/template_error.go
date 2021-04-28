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

type templateNotExistError struct {
	message string
}

func TemplateNotExistError(format string, args ...interface{}) error {
	return &templateNotExistError{
		message: fmt.Sprintf(format, args...),
	}
}
func (err *templateNotExistError) Error() string {
	return err.message
}

func IsTemplateNotExistError(err error) bool {
	_, ok := err.(*templateNotExistError)
	return ok
}
