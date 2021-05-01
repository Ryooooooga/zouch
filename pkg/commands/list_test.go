package commands_test

import (
	"bytes"
	"log"
	"os"
	"path"
	"testing"

	"github.com/Ryooooooga/zouch/pkg/commands"
	"github.com/Ryooooooga/zouch/pkg/config"
	"github.com/Ryooooooga/zouch/pkg/renderer"
)

func TestList(t *testing.T) {
	output := bytes.NewBufferString("")
	logger := log.New(os.Stderr, "", 0)

	rootDir := t.TempDir()
	templateDir := path.Join(rootDir, "templates")
	cfg := &config.Config{RootDir: rootDir, TemplateDir: templateDir}

	templates := newTestTemplateRepository()
	renderer := renderer.NewTextTemplateRenderer()

	cmd := commands.NewCommand(
		output,
		logger,
		cfg,
		templates,
		renderer,
		false,
		false,
	)

	err := cmd.List([]string{"THIS", "IS", "IGNORED"})
	if err != nil {
		t.Fatalf("cmd.List() returns an error %v", err)
	}

	expectedOutput := "_.txt\nmain.go\ntest.txt\n"

	result := output.String()
	if result != expectedOutput {
		t.Fatalf("result != %s, actual %s", expectedOutput, result)
	}
}
