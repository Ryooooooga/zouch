package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cmd := newTestCommand(t, false, false)

	err := cmd.List([]string{"THIS", "IS", "IGNORED"})
	assert.Nil(t, err)

	expectedOutput := "_.txt\nmain.go\ntest.txt\n"
	assert.Equal(t, expectedOutput, cmd.Output.(*bytes.Buffer).String())
}
