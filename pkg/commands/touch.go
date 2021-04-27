package commands

import (
	"github.com/Ryooooooga/zouch/pkg/errors"
)

func (cmd *Command) Touch(files []string) error {
	if len(files) == 0 {
		return errors.ShowHelpAndExitError("no files specified")
	}

	for _, file := range files {
		cmd.touchFile(file)
	}

	return nil
}

func (cmd *Command) touchFile(file string) error {

	return nil
}
