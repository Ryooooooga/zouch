package commands_test

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
)

func prepareTestDir(t *testing.T) string {
	tempDir := t.TempDir()

	_ = os.Mkdir(path.Join(tempDir, "test-dir"), 0755)
	_ = ioutil.WriteFile(path.Join(tempDir, "test.txt"), []byte("test"), 0644)
	_ = ioutil.WriteFile(path.Join(tempDir, "test2.txt"), []byte("test2"), 0644)
	_ = ioutil.WriteFile(path.Join(tempDir, "test-dir/hello.txt"), []byte("hello"), 0644)

	return tempDir
}

func TestAdd(t *testing.T) {
	scenarios := []struct {
		testname          string
		files             []string
		force             bool
		expectedTemplates []string
	}{
		{
			testname:          "test2.txt and hello.txt",
			files:             []string{"test2.txt", "./test-dir/hello.txt"},
			force:             false,
			expectedTemplates: []string{"_.txt", "hello.txt", "main.go", "test.txt", "test2.txt"},
		},
		{
			testname:          "force",
			files:             []string{"test.txt", "test-dir/hello.txt"},
			force:             true,
			expectedTemplates: []string{"_.txt", "hello.txt", "main.go", "test.txt"},
		},
	}

	for _, s := range scenarios {
		t.Run(s.testname, func(t *testing.T) {
			tempDir := prepareTestDir(t)
			_ = os.Chdir(tempDir)

			cmd := newTestCommand(t, false, s.force)

			err := cmd.Add(s.files)
			if err != nil {
				t.Fatalf("cmd.Add() returns an error %v", err)
			}

			templateFiles, err := cmd.Templates.ListTemplates()
			if err != nil {
				t.Fatalf("cmd.Templates.ListTemplates() returns an error %v", err)
			}

			if !reflect.DeepEqual(templateFiles, s.expectedTemplates) {
				t.Fatalf("templateFiles != %v, actual %v", s.expectedTemplates, templateFiles)
			}
		})
	}
}

func TestFailAdd(t *testing.T) {
	scenarios := []struct {
		testname string
		files    []string
		force    bool
	}{
		{
			testname: "empty",
			files:    []string{},
			force:    false,
		},
		{
			testname: "directory",
			files:    []string{"test-dir"},
			force:    true,
		},
		{
			testname: "not exists",
			files:    []string{"NO_SUCH_FILE"},
			force:    false,
		},
		{
			testname: "template exists",
			files:    []string{"test.txt"},
			force:    false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.testname, func(t *testing.T) {
			tempDir := prepareTestDir(t)
			_ = os.Chdir(tempDir)

			cmd := newTestCommand(t, false, s.force)

			err := cmd.Add(s.files)
			if err == nil {
				tpl, _ := cmd.Templates.FindTemplate("test.txt")
				t.Fatalf("cmd.Add() must return an error %v", string(tpl.Content))
			}
		})
	}
}
