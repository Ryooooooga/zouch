package errors_test

import (
	"fmt"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/errors"
)

func TestShowHelpAndExitError(t *testing.T) {
	t.Run("ShowHelpAndExitError.Error()", func(t *testing.T) {
		err := errors.ShowHelpAndExitError("test error message")
		if err.Error() != "test error message" {
			t.Fatalf("err.Error() != %v, actual %v", "test error message", err.Error())
		}
	})

	t.Run("IsShowHelpAndExitError", func(t *testing.T) {
		err := errors.ShowHelpAndExitError("")
		if !errors.IsShowHelpAndExitError(err) {
			t.Fatalf("IsShowHelpAndExitError(err) must return true")
		}

		err = fmt.Errorf("")
		if errors.IsShowHelpAndExitError(err) {
			t.Fatalf("IsShowHelpAndExitError(err) must return false")
		}
	})
}
