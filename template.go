package main

import (
	"io"
	"path"
	"path/filepath"
	"text/template"

	"github.com/Ryooooooga/zouch/pkg/renderer"
)

type definedValues struct {
	Path         string
	TemplatePath string
}

func renderTemplate(output io.Writer, source string, destination string) error {
	tpl, err := template.New(path.Base(source)).Funcs(renderer.DefaultFuncMap()).ParseFiles(source)
	if err != nil {
		return err
	}

	definedValues := definedValues{
		Path:         tryGetAbsPath(destination),
		TemplatePath: tryGetAbsPath(source),
	}

	if err := tpl.Execute(output, definedValues); err != nil {
		return err
	}

	return nil
}

func tryGetAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return absPath
}
