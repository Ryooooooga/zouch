package repositories_test

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/repositories"
)

func writeFile(t *testing.T, filename string, content string) {
	if err := ioutil.WriteFile(filename, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create test file %s: %v", filename, err)
	}
}

func makeDir(t *testing.T, filename string) {
	if err := os.Mkdir(filename, 0755); err != nil {
		t.Fatalf("failed to create test directory %s: %v", filename, err)
	}
}

func TestTemplateRepository(t *testing.T) {
	tempDir := t.TempDir()
	writeFile(t, path.Join(tempDir, "test1.txt"), `Test Template {{ .Path }}`)
	writeFile(t, path.Join(tempDir, "test2.txt"), `Today is {{ Now.Format "2006-01-02" }}!`)
	makeDir(t, path.Join(tempDir, "test-dir"))
	writeFile(t, path.Join(tempDir, "test-dir", "test3.txt"), `hello`)

	repo := repositories.NewTemplateRepository(tempDir)

	t.Run("ListTemplate", func(t *testing.T) {
		files, err := repo.ListTemplates()
		if err != nil {
			t.Fatalf("repo.ListTemplates() returns an error %v", err)
		}

		expectedFiles := []string{"test1.txt", "test2.txt"}

		if !reflect.DeepEqual(files, expectedFiles) {
			t.Fatalf("files != %v, actual %v", expectedFiles, files)
		}
	})

	t.Run("TemplatePathOf", func(t *testing.T) {
		templatePath := repo.TemplatePathOf("./tests/file.md")
		expectedTemplatePath := path.Join(tempDir, "file.md")

		if templatePath != expectedTemplatePath {
			t.Fatalf("templatePath != %s, actual %s", expectedTemplatePath, templatePath)
		}
	})
}
