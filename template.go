package main

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/stoewer/go-strcase"
)

type definedValues struct {
	Path         string
	TemplatePath string
}

func renderTemplate(source string, destination string) error {
	output, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer output.Close()

	tpl := template.New(path.Base(source)).Funcs(template.FuncMap{
		"Shell":          shell,
		"DateFormat":     dateFormat,
		"Now":            time.Now,
		"Base":           path.Base,
		"Dir":            path.Dir,
		"Getwd":          os.Getwd,
		"LowerCamelCase": strcase.LowerCamelCase,
		"UpperCamelCase": strcase.UpperCamelCase,
		"SnakeCase":      strcase.SnakeCase,
		"UpperSnakeCase": strcase.UpperSnakeCase,
		"KebabCase":      strcase.KebabCase,
		"UpperKebabCase": strcase.UpperKebabCase,
	})

	tpl, err = tpl.ParseFiles(source)
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

// Predefined functions
func shell(command string) (string, error) {
	bytes, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		return "", err
	}

	output := string(bytes)
	output = strings.TrimSuffix(output, "\n")
	output = strings.TrimSuffix(output, "\r")
	return output, nil
}

func dateFormat(format string, t time.Time) string {
	return t.Format(format)
}
