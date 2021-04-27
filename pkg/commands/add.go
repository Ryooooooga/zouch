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

	overwritten, err := cmd.templates.AddTemplate(filename, content, cmd.force)
	if err != nil {
		return err
	}

	if overwritten {
		fmt.Fprintf(cmd.output, "%s <- %s (overwrite)\n", cmd.templates.TemplatePathOf(filename), filename)
	} else {
		fmt.Fprintf(cmd.output, "%s <- %s\n", cmd.templates.TemplatePathOf(filename), filename)
	}

	return nil
}
