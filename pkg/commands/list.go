package commands

import (
	"fmt"
)

func (cmd *Command) List(files []string) error {
	templateFiles, err := cmd.templates.ListTemplates()
	if err != nil {
		return err
	}

	for _, tpl := range templateFiles {
		fmt.Fprintf(cmd.output, "%s\n", tpl)
	}

	return nil
}
