package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Ryooooooga/zouch/pkg/errors"
)

func (cmd *Command) Add(files []string) error {
	if len(files) == 0 {
		return errors.ShowHelpAndExitError("no files specified")
	}

	for _, filename := range files {
		if err := cmd.addFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func (cmd *Command) addFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", filename)
	} else if err != nil {
		return err
	}

	templateFilename, overwritten, err := cmd.Templates.AddTemplate(filename, content, cmd.Force)
	if err != nil {
		return err
	}

	if overwritten {
		cmd.Logger.Printf("%s <- %s (overwrite)\n", templateFilename, filename)
	} else {
		cmd.Logger.Printf("%s <- %s\n", templateFilename, filename)
	}

	return nil
}
