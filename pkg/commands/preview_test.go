package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreview(t *testing.T) {
	scenarios := []struct {
		testname       string
		files          []string
		expectedOutput string
	}{
		{
			testname:       "test.txt",
			files:          []string{"./foo/test.txt"},
			expectedOutput: "Test Text Template\nTEST.TXT\n",
		},
		{
			testname:       "test.txt and main.go",
			files:          []string{"./foo/test.txt", "../main.go"},
			expectedOutput: "Test Text Template\nTEST.TXT\nfunc main() {}\n",
		},
		{
			testname:       "bar.txt",
			files:          []string{"bar.txt"},
			expectedOutput: "Default Text Template : _.txt\n",
		},
	}

	for _, s := range scenarios {
		t.Run(s.testname, func(t *testing.T) {
			cmd := newTestCommand(t, false, false)

			err := cmd.Preview(s.files)
			assert.NoError(t, err)
			assert.Equal(t, s.expectedOutput, cmd.Output.(*bytes.Buffer).String())
		})
	}
}

func TestFailPreview(t *testing.T) {
	scenarios := []struct {
		testname string
		args     []string
	}{
		{
			testname: "empty",
			args:     []string{},
		},
		{
			testname: "no template - main.cpp",
			args:     []string{"main.cpp"},
		},
		{
			testname: "no template - baz.go",
			args:     []string{"baz.go"},
		},
	}

	for _, s := range scenarios {
		t.Run(s.testname, func(t *testing.T) {
			cmd := newTestCommand(t, false, false)
			err := cmd.Preview(s.args)
			assert.Error(t, err)
		})
	}
}
