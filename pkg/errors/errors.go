package errors

type showHelpAndExitError struct {
	message string
}

func ShowHelpAndExitError(message string) error {
	return &showHelpAndExitError{
		message,
	}
}
func (err *showHelpAndExitError) Error() string {
	return err.message
}

func IsShowHelpAndExitError(err error) bool {
	_, ok := err.(*showHelpAndExitError)
	return ok
}
