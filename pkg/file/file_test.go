package file_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/file"
)

func TestIsFile(t *testing.T) {
	// Prepare
	tmpDir := t.TempDir()
	testFile := path.Join(tmpDir, "test-file")
	testDir := path.Join(tmpDir, "test-dir")
	testSymFile := path.Join(tmpDir, "test-symfile")
	testSymDir := path.Join(tmpDir, "test-symdir")
	testNoSuch := path.Join(tmpDir, "NO_SUCH_FILE_OR_DIRECTORY")
	if err := ioutil.WriteFile(testFile, []byte{}, 0644); err != nil {
		t.Fatalf("failed to create a test file %s", testFile)
	}
	if err := os.Mkdir(testDir, 0755); err != nil {
		t.Fatalf("failed to create a test directory %s", testDir)
	}
	if err := os.Symlink(testFile, testSymFile); err != nil {
		t.Fatalf("failed to create a test symlink %s", testSymFile)
	}
	if err := os.Symlink(testDir, testSymDir); err != nil {
		t.Fatalf("failed to create a test symlink %s", testSymDir)
	}

	t.Run("IsFile for a file must return true", func(t *testing.T) {
		isFile, err := file.IsFile(testFile)
		if err != nil {
			t.Fatalf("IsFile(%v) returns an error %v", testFile, err)
		} else if !isFile {
			t.Fatalf("IsFile(%v) must return true", testFile)
		}
	})

	t.Run("IsFile for a directory must return false", func(t *testing.T) {
		_, err := file.IsFile(testDir)
		if err == nil {
			t.Fatalf("IsFile(%v) must return an error", testDir)
		}
	})

	t.Run("IsFile for a symlink to file must return true", func(t *testing.T) {
		isFile, err := file.IsFile(testSymFile)
		if err != nil {
			t.Fatalf("IsFile(%v) returns an error %v", testSymFile, err)
		} else if !isFile {
			t.Fatalf("IsFile(%v) must return true", testSymFile)
		}
	})

	t.Run("IsFile for a symlink to directory must return false", func(t *testing.T) {
		_, err := file.IsFile(testSymDir)
		if err == nil {
			t.Fatalf("IsFile(%v) must return an error", testSymDir)
		}
	})

	t.Run("IsFile fails when a file or a directory does not exist", func(t *testing.T) {
		isFile, err := file.IsFile(testNoSuch)
		if err != nil {
			t.Fatalf("IsFile(%v) returns an error %v", testNoSuch, err)
		} else if isFile {
			t.Fatalf("IsFile(%v) must return false", testNoSuch)
		}
	})
}
