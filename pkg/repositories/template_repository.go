package repositories

import (
	"io/ioutil"
	"os"
)

type TemplateRepository interface {
	ListTemplates() ([]string, error)
	AddTemplate(name string, content string) error
	FindTemplate(filepath string) (string, error)
}

type templateRepository struct {
	rootDir string
}

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

func (r *templateRepository) AddTemplate(name string, content string) error {
	panic("unimplemented")
}

func (r *templateRepository) FindTemplate(filepath string) (string, error) {
	panic("unimplemented")
}
