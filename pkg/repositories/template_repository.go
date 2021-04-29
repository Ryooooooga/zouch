package repositories

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/Ryooooooga/zouch/pkg/errors"
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

	files, err := ioutil.ReadDir(r.templateDir)
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

func (r *templateRepository) AddTemplate(filename string, content []byte, overwrite bool) (bool, error) {
	templatePath := r.TemplatePathOf(filename)

	var overwritten bool
	stat, err := os.Stat(templatePath)
	if os.IsNotExist(err) {
		// Template does not exist (ok)
		overwritten = false
	} else if err != nil {
		return false, err
	} else if stat.IsDir() {
		return false, fmt.Errorf("%s is a directory", templatePath)
	} else if !overwrite {
		return false, errors.TemplateExistError("%s already exists", templatePath)
	} else {
		overwritten = true
	}

	if err := os.MkdirAll(path.Dir(templatePath), DirectoryPermission); err != nil {
		return false, err
	}

	if err := ioutil.WriteFile(templatePath, content, FilePermission); err != nil {
		return false, err
	}

	return overwritten, nil
}

func (r *templateRepository) FindTemplate(filename string) (TemplateFile, error) {
	basename := path.Base(filename)
	templatePath := path.Join(r.templateDir, basename)

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
	templatePath := path.Join(r.templateDir, basename)

	return templatePath
}
