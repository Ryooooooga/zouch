package cli_test

import (
	"bytes"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/cli"
	"github.com/stretchr/testify/assert"
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

			err := app.Run(s.args)
			assert.Nil(t, err)

			assert.NotEmpty(t, writer.String())
		})
	}
}
