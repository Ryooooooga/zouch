package repositories

import (
	"fmt"
	"os"
	"path"
	"sort"
)

type TemplateFile struct {
	Path    string
	Content []byte
}

type TemplateRepository interface {
	ListTemplates() ([]string, error)
	AddTemplate(filename string, content []byte, overwrite bool) (templateFilename string, overwritten bool, err error)
	FindTemplate(filename string) (*TemplateFile, error)
}

type templateRepository struct {
	templateDir string
}

const (
	FilePermission      = 0644
	DirectoryPermission = 0755
)

func NewTemplateRepository(templateDir string) TemplateRepository {
	return &templateRepository{
		templateDir,
	}
}

func (r *templateRepository) ListTemplates() ([]string, error) {
	list := []string{}

	files, err := os.ReadDir(r.templateDir)
	if os.IsNotExist(err) {
		// Template directory does not exist
		return list, nil
	} else if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			list = append(list, file.Name())
		}
	}

	sort.Strings(list)
	return list, nil
}

func (r *templateRepository) AddTemplate(filename string, content []byte, overwrite bool) (string, bool, error) {
	basename := path.Base(filename)
	templatePath := path.Join(r.templateDir, basename)

	var overwritten bool
	stat, err := os.Stat(templatePath)
	if os.IsNotExist(err) {
		// Template does not exist (ok)
		overwritten = false
	} else if err != nil {
		return "", false, err
	} else if stat.IsDir() {
		return "", false, fmt.Errorf("%s is a directory", templatePath)
	} else if !overwrite {
		return "", false, fmt.Errorf("%s already exists", templatePath)
	} else {
		overwritten = true
	}

	if err := os.MkdirAll(path.Dir(templatePath), DirectoryPermission); err != nil {
		return "", false, err
	}

	if err := os.WriteFile(templatePath, content, FilePermission); err != nil {
		return "", false, err
	}

	return templatePath, overwritten, nil
}

func (r *templateRepository) FindTemplate(filename string) (*TemplateFile, error) {
	basename := path.Base(filename)
	templatePath := path.Join(r.templateDir, basename)

	content, err := os.ReadFile(templatePath)
	if err == nil {
		return &TemplateFile{
			Path:    templatePath,
			Content: content,
		}, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	ext := path.Ext(filename)
	fallbackTemplatePath := path.Join(r.templateDir, "_"+ext)

	content, err = os.ReadFile(fallbackTemplatePath)
	if err == nil {
		return &TemplateFile{
			Path:    fallbackTemplatePath,
			Content: content,
		}, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	return nil, nil
}
