package repositories

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/Ryooooooga/zouch/pkg/errors"
	"github.com/Ryooooooga/zouch/pkg/file"
)

type TemplateFile struct {
	Path    string
	Content []byte
}

type TemplateRepository interface {
	ListTemplates() ([]string, error)
	AddTemplate(filename string, content []byte, overwrite bool) (overwritten bool, err error)
	FindTemplate(filename string) (TemplateFile, error)
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

	templateExists, err := file.IsFile(templatePath)
	if err != nil {
		return false, err
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

func (r *templateRepository) FindTemplate(filename string) (TemplateFile, error) {
	basename := path.Base(filename)
	templatePath := path.Join(r.rootDir, basename)

	content, err := ioutil.ReadFile(templatePath)
	if os.IsNotExist(err) {
		return TemplateFile{}, errors.TemplateNotExistError("%s does not exist", templatePath)
	} else if err != nil {
		return TemplateFile{}, err
	}

	return TemplateFile{
		Path:    templatePath,
		Content: content,
	}, nil
}

func (r *templateRepository) TemplatePathOf(filename string) string {
	basename := path.Base(filename)
	templatePath := path.Join(r.rootDir, basename)

	return templatePath
}
