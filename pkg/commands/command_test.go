package commands_test

import (
	"bytes"
	"log"
	"path"
	"sort"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/commands"
	"github.com/Ryooooooga/zouch/pkg/config"
	"github.com/Ryooooooga/zouch/pkg/renderer"
	"github.com/Ryooooooga/zouch/pkg/repositories"
)

func newTestCommand(t *testing.T, createDir bool, force bool) *commands.Command {
	output := bytes.NewBufferString("")
	logger := log.New(bytes.NewBufferString(""), "", 0)

	rootDir := t.TempDir()
	templateDir := path.Join(rootDir, "templates")
	cfg := &config.Config{RootDir: rootDir, TemplateDir: templateDir}

	templates := newTestTemplateRepository()
	renderer := renderer.NewTextTemplateRenderer()

	return commands.NewCommand(
		output,
		logger,
		cfg,
		templates,
		renderer,
		createDir,
		force,
	)
}

type testTemplateRepository struct {
	Files map[string]string
}

func newTestTemplateRepository() repositories.TemplateRepository {
	files := map[string]string{
		"_.txt":    "Default Text Template : {{ .TemplateFilename | Base }}\n",
		"test.txt": "Test Text Template\n{{ .Filename | Base | UpperSnakeCase }}\n",
		"main.go":  "func main() {}\n",
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
	basename := path.Base(filename)

	_, ok := r.Files[basename]
	if ok && !overwrite {
		return "", false, nil
	}

	r.Files[basename] = string(content)

	return basename, ok, nil
}

func (r *testTemplateRepository) FindTemplate(filename string) (*repositories.TemplateFile, error) {
	basename := path.Base(filename)

	content, ok := r.Files[basename]
	if ok {
		return &repositories.TemplateFile{
			Path:    basename,
			Content: []byte(content),
		}, nil
	}

	fallbackTemplate := "_" + path.Ext(basename)
	content, ok = r.Files[fallbackTemplate]
	if ok {
		return &repositories.TemplateFile{
			Path:    fallbackTemplate,
			Content: []byte(content),
		}, nil
	}

	return nil, nil
}
