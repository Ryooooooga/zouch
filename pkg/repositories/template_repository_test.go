package repositories_test

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/errors"
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

	t.Run("ListTemplate", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(tempDir)

		files, err := repo.ListTemplates()
		if err != nil {
			t.Fatalf("ListTemplates() returns an error %v", err)
		}

		expectedFiles := []string{"test1.txt", "test2.txt"}

		if !reflect.DeepEqual(files, expectedFiles) {
			t.Fatalf("files != %v, actual %v", expectedFiles, files)
		}
	})

	t.Run("ListTemplateNotExists", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(path.Join(tempDir, "NO_SUCH_DIR"))

		files, err := repo.ListTemplates()
		if err != nil {
			t.Fatalf("ListTemplates() returns an error %v", err)
		}

		if len(files) != 0 {
			t.Fatalf("files != {}, actual %v", files)
		}
	})

	t.Run("AddTemplate", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(tempDir)

		overwritten, err := repo.AddTemplate("add-test1.txt", []byte("add-test1"), false)
		if err != nil {
			t.Fatalf("AddTemplate() returns an error %v", err)
		}
		if overwritten {
			t.Fatalf("overwritten must be false")
		}
		defer os.Remove(path.Join(tempDir, "add-test1.txt"))

		files, err := repo.ListTemplates()
		if err != nil {
			t.Fatalf("ListTemplates() returns an error %v", err)
		}

		expectedFiles := []string{"add-test1.txt", "test1.txt", "test2.txt"}
		if !reflect.DeepEqual(files, expectedFiles) {
			t.Fatalf("files != %v, actual %v", expectedFiles, files)
		}

		_, err = repo.AddTemplate("add-test1.txt", []byte("add-test1"), false)
		if err == nil || !errors.IsTemplateExistError(err) {
			t.Fatalf("AddTemplate() returns TemplateExistError %v", err)
		}

		overwritten, err = repo.AddTemplate("add-test1.txt", []byte("add-test1"), true)
		if err != nil {
			t.Fatalf("AddTemplate() returns an error %v", err)
		}
		if !overwritten {
			t.Fatalf("overwritten must be true")
		}
	})

	t.Run("FindTemplate", func(t *testing.T) {
		scenarios := []struct {
			filename        string
			expectedPath    string
			expectedContent string
		}{
			{
				filename:        "test1.txt",
				expectedPath:    path.Join(tempDir, "test1.txt"),
				expectedContent: `Test Template {{ .Path }}`,
			},
			{
				filename:        "some/directory/test1.txt",
				expectedPath:    path.Join(tempDir, "test1.txt"),
				expectedContent: `Test Template {{ .Path }}`,
			},
			{
				filename:        "../test2.txt",
				expectedPath:    path.Join(tempDir, "test2.txt"),
				expectedContent: `Today is {{ Now.Format "2006-01-02" }}!`,
			},
		}

		for _, s := range scenarios {
			t.Run(s.filename, func(t *testing.T) {
				repo := repositories.NewTemplateRepository(tempDir)

				tpl, err := repo.FindTemplate(s.filename)
				if err != nil {
					t.Fatalf("FindTemplate() returns an error %v", err)
				}
				if tpl.Path != s.expectedPath {
					t.Fatalf("tpl.Path != %s, actual %s", s.expectedPath, tpl.Path)
				}
				if string(tpl.Content) != s.expectedContent {
					t.Fatalf("tpl.Content != %s, actual %s", s.expectedContent, string(tpl.Content))
				}
			})
		}
	})

	t.Run("FailFindTemplate", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(tempDir)

		_, err := repo.FindTemplate("test3.txt")
		if !errors.IsTemplateNotExistError(err) {
			t.Fatalf("err must be TemplateNotExistError %v", err)
		}
		if err == nil {
			t.Fatalf("FindTemplate() must return an error")
		}

		_, err = repo.FindTemplate("test-dir")
		if errors.IsTemplateNotExistError(err) {
			t.Fatalf("err must not be TemplateNotExistError %v", err)
		}
		if err == nil {
			t.Fatalf("FindTemplate() must return an error")
		}
	})

	t.Run("TemplatePathOf", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(tempDir)

		templatePath := repo.TemplatePathOf("./tests/file.md")
		expectedTemplatePath := path.Join(tempDir, "file.md")

		if templatePath != expectedTemplatePath {
			t.Fatalf("templatePath != %s, actual %s", expectedTemplatePath, templatePath)
		}
	})
}
