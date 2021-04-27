package commands

import (
	"github.com/Ryooooooga/zouch/pkg/errors"
)

func (cmd *Command) Touch(files []string) error {
	if len(files) == 0 {
		return errors.ShowHelpAndExitError("no files specified")
	}

	for _, filename := range files {
		if err := cmd.touchFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func (cmd *Command) touchFile(filename string) error {

	return nil
}
