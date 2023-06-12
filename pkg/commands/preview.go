package commands

import (
	"fmt"

	"github.com/Ryooooooga/zouch/pkg/errors"
)

func (cmd *Command) Preview(files []string) error {
	if len(files) == 0 {
		return errors.ShowHelpAndExitError("no files specified")
	}

	for _, filename := range files {
		if err := cmd.previewFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func (cmd *Command) previewFile(filename string) error {
	tpl, err := cmd.Templates.FindTemplate(filename)
	if err != nil {
		return err
	}
	if tpl == nil {
		return fmt.Errorf("template for %s does not exist", filename)
	}

	data := templateVariables(filename, tpl)

	if err := cmd.Renderer.RenderTemplate(cmd.Output, tpl, data); err != nil {
		return err
	}

	return nil
}
