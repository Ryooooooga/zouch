package commands

import (
	"fmt"
)

func (cmd *Command) List(files []string) error {
	templateFiles, err := cmd.Templates.ListTemplates()
	if err != nil {
		return err
	}

	for _, tpl := range templateFiles {
		fmt.Fprintf(cmd.Output, "%s\n", tpl)
	}

	return nil
}
