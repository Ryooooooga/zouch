package commands

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/Ryooooooga/zouch/pkg/errors"
	"github.com/Ryooooooga/zouch/pkg/repositories"
)

const (
	FilePermission      = 0644
	DirectoryPermission = 0755
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
	var fileExists bool
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fileExists = false
	} else if err != nil {
		return err
	} else if stat.IsDir() || !cmd.Force {
		return cmd.updateTimestamp(filename)
	} else {
		fileExists = true
	}

	tpl, err := cmd.Templates.FindTemplate(filename)
	if err != nil {
		return err
	}

	if cmd.CreateDir {
		if err := os.MkdirAll(path.Dir(filename), DirectoryPermission); err != nil {
			return err
		}
	}

	if tpl != nil {
		return cmd.renderTemplate(filename, tpl, fileExists)
	} else if fileExists {
		return cmd.updateTimestamp(filename)
	} else {
		return cmd.createNewFile(filename)
	}
}

func (cmd *Command) createNewFile(filename string) error {
	if err := ioutil.WriteFile(filename, []byte{}, FilePermission); err != nil {
		return err
	}

	cmd.Logger.Printf("%s (new)", filename)
	return nil
}

func (cmd *Command) updateTimestamp(filename string) error {
	now := cmd.Now()
	atime := now
	mtime := now
	if err := os.Chtimes(filename, atime, mtime); err != nil {
		return err
	}

	cmd.Logger.Printf("%s (update timestamp)", filename)
	return nil
}

func (cmd *Command) renderTemplate(filename string, tpl *repositories.TemplateFile, overwrite bool) error {
	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	data := templateVariables(filename, tpl)

	if err := cmd.Renderer.RenderTemplate(output, tpl, data); err != nil {
		return err
	}

	if overwrite {
		cmd.Logger.Printf("%s -> %s (overwrite)", tpl.Path, filename)
	} else {
		cmd.Logger.Printf("%s -> %s", tpl.Path, filename)
	}
	return nil
}

func templateVariables(filename string, tpl *repositories.TemplateFile) map[interface{}]interface{} {
	return map[interface{}]interface{}{
		"Filename":         filename,
		"TemplateFilename": tpl.Path,
	}
}
