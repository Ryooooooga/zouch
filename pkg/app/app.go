package app

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/Ryooooooga/zouch/pkg/file"
)

const (
	DirectoryPermission = 0755
	FilePermission      = 0644
)

type App struct {
	logger        *log.Logger
	rootDir       string
	createDirFlag bool
	forceFlag     bool
}

func NewApp(logger *log.Logger, rootDir string, createDirFlag bool, forceFlag bool) *App {
	return &App{
		logger,
		rootDir,
		createDirFlag,
		forceFlag,
	}
}

func (app *App) AddTemplateFiles(files []string) error {
	for _, file := range files {
		if err := app.addTemplateFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) addTemplateFile(filename string) error {
	if stat, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", filename)
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("%s is a directory", filename)
	}

	overwriteTemplate := false
	destination := app.resolveTemplatePath(filename)
	if stat, err := os.Stat(destination); os.IsNotExist(err) {
		// File does not exist (ok)
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("%s is a directory", destination)
	} else if !app.forceFlag {
		return fmt.Errorf("template %s already exists", destination)
	} else {
		overwriteTemplate = true
	}

	if err := os.MkdirAll(path.Dir(destination), DirectoryPermission); err != nil {
		return err
	}

	if err := file.Copy(filename, destination); err != nil {
		return err
	}

	if overwriteTemplate {
		app.logger.Printf("overwrite template: %s -> %s", filename, destination)
	} else {
		app.logger.Printf("add new template: %s -> %s", filename, destination)
	}

	return nil
}

func (app *App) TouchFiles(files []string) error {
	for _, file := range files {
		if err := app.touchFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) touchFile(filename string) error {
	if stat, err := os.Stat(filename); os.IsNotExist(err) {
		// File does not exist (ok)
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("%s is a directory", filename)
	} else if !app.forceFlag {
		t := time.Now()
		atime := t
		mtime := t
		if err := os.Chtimes(filename, atime, mtime); err != nil {
			return err
		}
		app.logger.Printf("update timestamp of %s\n", filename)
		return nil
	}

	source := app.resolveTemplatePath(filename)
	templateExists := true
	if stat, err := os.Stat(source); os.IsNotExist(err) {
		templateExists = false
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("template %s is a directory", source)
	}

	if app.createDirFlag {
		if err := os.MkdirAll(path.Dir(filename), DirectoryPermission); err != nil {
			return err
		}
	}

	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	if !templateExists {
		app.logger.Printf("create empty file %s\n", filename)
		return nil
	}

	if err := renderTemplate(output, source, filename); err != nil {
		return err
	}

	app.logger.Printf("create from template: %s -> %s\n", source, filename)
	return nil
}

func (app *App) resolveTemplatePath(file string) string {
	basename := filepath.Base(file)

	return path.Join(app.rootDir, basename)
}
