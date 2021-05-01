package commands_test

import (
	"bytes"
	"testing"
)

func TestList(t *testing.T) {
	cmd := newTestCommand(t, false, false)

	err := cmd.List([]string{"THIS", "IS", "IGNORED"})
	if err != nil {
		t.Fatalf("cmd.List() returns an error %v", err)
	}

	expectedOutput := "_.txt\nmain.go\ntest.txt\n"

	result := cmd.Output.(*bytes.Buffer).String()
	if result != expectedOutput {
		t.Fatalf("result != %s, actual %s", expectedOutput, result)
	}
}
