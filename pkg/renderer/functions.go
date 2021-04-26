package renderer

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/stoewer/go-strcase"
)

func DefaultFuncMap() template.FuncMap {
	return template.FuncMap{
		"Shell":          shell,
		"Now":            time.Now,
		"Base":           path.Base,
		"Dir":            path.Dir,
		"Getwd":          os.Getwd,
		"Getenv":         os.Getenv,
		"LowerCamelCase": strcase.LowerCamelCase,
		"UpperCamelCase": strcase.UpperCamelCase,
		"SnakeCase":      strcase.SnakeCase,
		"UpperSnakeCase": strcase.UpperSnakeCase,
		"KebabCase":      strcase.KebabCase,
		"UpperKebabCase": strcase.UpperKebabCase,
	}
}

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
