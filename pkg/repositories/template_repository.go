package repositories

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/Ryooooooga/zouch/pkg/errors"
)

type TemplateRepository interface {
	ListTemplates() ([]string, error)
	AddTemplate(filename string, content []byte, overwrite bool) (overwritten bool, err error)
	FindTemplate(filename string) (string, error)
	TemplatePathOf(filename string) string
}

type templateRepository struct {
	rootDir string
}

const (
	FilePermission      = 0644
	DirectoryPermission = 0755
)

func NewTemplateRepository(rootDir string) TemplateRepository {
	return &templateRepository{
		rootDir,
	}
}

func (r *templateRepository) ListTemplates() ([]string, error) {
	list := []string{}

	files, err := ioutil.ReadDir(r.rootDir)
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

	return list, nil
}

func (r *templateRepository) AddTemplate(filename string, content []byte, overwrite bool) (bool, error) {
	templatePath := r.TemplatePathOf(filename)

	var templateExists bool
	if stat, err := os.Stat(templatePath); os.IsNotExist(err) {
		// `templatePath` does not exist (ok)
		templateExists = false
	} else if err != nil {
		return false, nil
	} else if stat.IsDir() {
		return false, fmt.Errorf("%s is a directory", templatePath)
	} else {
		templateExists = true
	}

	if templateExists && !overwrite {
		return false, errors.TemplateExistError("%s already exists", templatePath)
	}

	if err := os.MkdirAll(path.Dir(templatePath), DirectoryPermission); err != nil {
		return false, err
	}

	if err := ioutil.WriteFile(templatePath, content, FilePermission); err != nil {
		return templateExists, err
	}

	return templateExists, nil
}

func (r *templateRepository) FindTemplate(filename string) (string, error) {
	panic("unimplemented")
}

func (r *templateRepository) TemplatePathOf(filename string) string {
	basename := path.Base(filename)
	templatePath := path.Join(r.rootDir, basename)

	return templatePath
}
