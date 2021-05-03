package errors_test

import (
	"fmt"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestShowHelpAndExitError(t *testing.T) {
	t.Run("ShowHelpAndExitError.Error()", func(t *testing.T) {
		err := errors.ShowHelpAndExitError("test error message")
		assert.Equal(t, "test error message", err.Error())
	})

	t.Run("IsShowHelpAndExitError", func(t *testing.T) {
		err := errors.ShowHelpAndExitError("")
		assert.True(t, errors.IsShowHelpAndExitError(err))

		err = fmt.Errorf("")
		assert.False(t, errors.IsShowHelpAndExitError(err))
	})
}
