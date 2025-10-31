package repositories_test

import (
	"os"
	"path"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/repositories"
	"github.com/stretchr/testify/assert"
)

func writeFile(t *testing.T, filename string, content string) {
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
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
	writeFile(t, path.Join(tempDir, "test2.md"), `Today is {{ Now.Format "2006-01-02" }}!`)
	writeFile(t, path.Join(tempDir, "_.md"), `Default markdown template`)
	makeDir(t, path.Join(tempDir, "test-dir"))
	writeFile(t, path.Join(tempDir, "test-dir", "test3.txt"), `hello`)

	t.Run("ListTemplate", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(tempDir)

		files, err := repo.ListTemplates()
		assert.Nil(t, err)

		expectedFiles := []string{"_.md", "test1.txt", "test2.md"}
		assert.Equal(t, expectedFiles, files)
	})

	t.Run("ListTemplateNotExists", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(path.Join(tempDir, "NO_SUCH_DIR"))

		files, err := repo.ListTemplates()
		assert.Nil(t, err)
		assert.Empty(t, files)
	})

	t.Run("AddTemplate", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(tempDir)

		templateFilename, overwritten, err := repo.AddTemplate("add-test1.txt", []byte("add-test1"), false)
		assert.Nil(t, err)

		expectedTemplateFilename := path.Join(tempDir, "add-test1.txt")
		assert.Equal(t, expectedTemplateFilename, templateFilename)
		assert.False(t, overwritten)
		defer func() {
			_ = os.Remove(path.Join(tempDir, "add-test1.txt"))
		}()

		files, err := repo.ListTemplates()
		assert.Nil(t, err)

		expectedFiles := []string{"_.md", "add-test1.txt", "test1.txt", "test2.md"}
		assert.Equal(t, expectedFiles, files)

		_, _, err = repo.AddTemplate("add-test1.txt", []byte("add-test1"), false)
		assert.Error(t, err)

		templateFilename, overwritten, err = repo.AddTemplate("add-test1.txt", []byte("add-test1"), true)
		assert.Nil(t, err)
		assert.Equal(t, expectedTemplateFilename, templateFilename)
		assert.True(t, overwritten)
	})

	t.Run("FindTemplate", func(t *testing.T) {
		scenarios := []struct {
			filename string
			expected *repositories.TemplateFile
		}{
			{
				filename: "test1.txt",
				expected: &repositories.TemplateFile{
					Path:    path.Join(tempDir, "test1.txt"),
					Content: []byte(`Test Template {{ .Path }}`),
				},
			},
			{
				filename: "some/directory/test1.txt",
				expected: &repositories.TemplateFile{
					Path:    path.Join(tempDir, "test1.txt"),
					Content: []byte(`Test Template {{ .Path }}`),
				},
			},
			{
				filename: "../test2.md",
				expected: &repositories.TemplateFile{
					Path:    path.Join(tempDir, "test2.md"),
					Content: []byte(`Today is {{ Now.Format "2006-01-02" }}!`),
				},
			},
			{
				filename: "../fallback.md",
				expected: &repositories.TemplateFile{
					Path:    path.Join(tempDir, "_.md"),
					Content: []byte(`Default markdown template`),
				},
			},
			{
				filename: "test3.txt",
				expected: nil,
			},
		}

		for _, s := range scenarios {
			t.Run(s.filename, func(t *testing.T) {
				repo := repositories.NewTemplateRepository(tempDir)

				tpl, err := repo.FindTemplate(s.filename)
				assert.Nil(t, err)
				assert.Equal(t, s.expected, tpl)
			})
		}
	})

	t.Run("FailFindTemplate", func(t *testing.T) {
		repo := repositories.NewTemplateRepository(tempDir)

		_, err := repo.FindTemplate("test-dir")
		assert.Error(t, err)
	})
}
