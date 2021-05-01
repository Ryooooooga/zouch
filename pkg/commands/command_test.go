package commands_test

import (
	"sort"

	"github.com/Ryooooooga/zouch/pkg/repositories"
)

type testTemplateRepository struct {
	Files map[string]string
}

func newTestTemplateRepository() repositories.TemplateRepository {
	files := map[string]string{
		"_.txt":    "Default Text Template {{ .Filename | Base }}",
		"test.txt": "Test Text Template {{ .TemplateFilename | Base | UpperSnakeCase }}",
		"main.go":  "func main() {}",
	}

	return &testTemplateRepository{
		Files: files,
	}
}

func (r *testTemplateRepository) ListTemplates() ([]string, error) {
	list := []string{}

	for filename := range r.Files {
		list = append(list, filename)
	}

	sort.Strings(list)

	return list, nil
}

func (r *testTemplateRepository) AddTemplate(filename string, content []byte, overwrite bool) (string, bool, error) {
	_, ok := r.Files[filename]
	if ok && !overwrite {
		return "", false, nil
	}

	r.Files[filename] = string(content)

	return filename, ok, nil
}

func (r *testTemplateRepository) FindTemplate(filename string) (*repositories.TemplateFile, error) {
	content, ok := r.Files[filename]
	if !ok {
		return nil, nil
	}

	return &repositories.TemplateFile{
		Path:    filename,
		Content: []byte(content),
	}, nil
}
