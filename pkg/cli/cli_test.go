package cli_test

import (
	"bytes"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/cli"
)

func TestHelpAndVersion(t *testing.T) {
	type scenario struct {
		name string
		args []string
	}

	scenarios := []scenario{
		{
			name: "-h",
			args: []string{"zouch", "-h"},
		},
		{
			name: "--help",
			args: []string{"zouch", "--help"},
		},
		{
			name: "-v",
			args: []string{"zouch", "-v"},
		},
		{
			name: "--version",
			args: []string{"zouch", "--version"},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			writer := bytes.NewBufferString("")
			app := cli.New("test")
			app.Writer = writer

			if err := app.Run(s.args); err != nil {
				t.Fatal("must not return an error")
			}

			if output := writer.String(); len(output) == 0 {
				t.Fatal("must show something")
			}
		})
	}
}
