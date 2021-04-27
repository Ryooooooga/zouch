package errors_test

import (
	"fmt"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/errors"
)

func TestTemplateExistError(t *testing.T) {
	t.Run("TemplateExistError.Error()", func(t *testing.T) {
		err := errors.TemplateExistError("test error message %d", 42)
		if err.Error() != "test error message 42" {
			t.Fatalf("err.Error() != %v, actual %v", "test error message 42", err.Error())
		}
	})

	t.Run("IsTemplateExistError", func(t *testing.T) {
		err := errors.TemplateExistError("")
		if !errors.IsTemplateExistError(err) {
			t.Fatalf("IsTemplateExistError(err) must return true")
		}

		err = fmt.Errorf("")
		if errors.IsTemplateExistError(err) {
			t.Fatalf("IsTemplateExistError(err) must return false")
		}
	})
}
