package renderer

import (
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/Ryooooooga/zouch/pkg/repositories"
	"github.com/stoewer/go-strcase"
)

type TemplateRenderer interface {
	RenderTemplate(output io.Writer, templateFile repositories.TemplateFile, data interface{}) error
}

type TextTemplateRenderer struct {
	FuncMap template.FuncMap
}

func NewTextTemplateRenderer() *TextTemplateRenderer {
	return &TextTemplateRenderer{
		FuncMap: template.FuncMap{
			"Shell":          shell,
			"Now":            time.Now,
			"Base":           path.Base,
			"Dir":            path.Dir,
			"Abs":            filepath.Abs,
			"Getwd":          os.Getwd,
			"Getenv":         os.Getenv,
			"LowerCamelCase": strcase.LowerCamelCase,
			"UpperCamelCase": strcase.UpperCamelCase,
			"SnakeCase":      strcase.SnakeCase,
			"UpperSnakeCase": strcase.UpperSnakeCase,
			"KebabCase":      strcase.KebabCase,
			"UpperKebabCase": strcase.UpperKebabCase,
		},
	}
}

func (r *TextTemplateRenderer) RenderTemplate(output io.Writer, templateFile repositories.TemplateFile, data interface{}) error {
	name := path.Base(templateFile.Path)
	content := string(templateFile.Content)

	tpl, err := template.New(name).Funcs(r.FuncMap).Parse(content)
	if err != nil {
		return err
	}

	return tpl.Execute(output, data)
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
