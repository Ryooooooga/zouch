package renderer

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/Ryooooooga/zouch/pkg/repositories"
	"github.com/stoewer/go-strcase"
)

type TemplateRenderer interface {
	RenderTemplate(output io.Writer, templateFile *repositories.TemplateFile, data interface{}) error
}

type TextTemplateRenderer struct {
	FuncMap template.FuncMap
}

func NewTextTemplateRenderer() *TextTemplateRenderer {
	return &TextTemplateRenderer{
		FuncMap: template.FuncMap{
			"Shell":           shell,
			"Now":             time.Now,
			"Base":            path.Base,
			"Ext":             path.Ext,
			"Dir":             path.Dir,
			"Abs":             filepath.Abs,
			"Getwd":           os.Getwd,
			"Getenv":          os.Getenv,
			"HasPrefix":       strings.HasPrefix,
			"HasSuffix":       strings.HasSuffix,
			"LowerCamelCase":  strcase.LowerCamelCase,
			"UpperCamelCase":  strcase.UpperCamelCase,
			"SnakeCase":       strcase.SnakeCase,
			"UpperSnakeCase":  strcase.UpperSnakeCase,
			"KebabCase":       strcase.KebabCase,
			"UpperKebabCase":  strcase.UpperKebabCase,
			"Replace":         strings.Replace,
			"ReplaceAll":      strings.ReplaceAll,
			"RegexReplaceAll": regexReplace,
		},
	}
}

func (r *TextTemplateRenderer) RenderTemplate(output io.Writer, templateFile *repositories.TemplateFile, data interface{}) error {
	name := path.Base(templateFile.Path)
	content := string(templateFile.Content)

	tpl, err := template.New(name).Funcs(r.FuncMap).Parse(content)
	if err != nil {
		return err
	}

	return tpl.Execute(output, data)
}

func shell(command string) (string, error) {
	var outBuffer bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = &outBuffer
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output := outBuffer.String()
	output = strings.TrimSuffix(output, "\n")
	output = strings.TrimSuffix(output, "\r")
	return output, nil
}

func regexReplace(src string, pattern string, replace string) (string, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	return regex.ReplaceAllString(src, replace), nil
}
