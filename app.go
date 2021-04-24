package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

type App struct {
	logger        *log.Logger
	rootDir       string
	createDirFlag bool
	forceFlag     bool
}

func newApp(logger *log.Logger, rootDir string, createDirFlag bool, forceFlag bool) *App {
	return &App{
		logger,
		rootDir,
		createDirFlag,
		forceFlag,
	}
}

func (app *App) addTemplateFiles(files []string) error {
	for _, file := range files {
		if err := app.addTemplateFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) addTemplateFile(file string) error {
	if stat, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", file)
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("%s is a directory", file)
	}

	overwriteTemplate := false
	destination := app.resolveTemplatePath(file)
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

	if err := os.MkdirAll(path.Dir(destination), 0755); err != nil {
		return err
	}

	if err := copyFile(file, destination); err != nil {
		return err
	}

	if overwriteTemplate {
		app.logger.Printf("overwrite template: %s -> %s", file, destination)
	} else {
		app.logger.Printf("add new template: %s -> %s", file, destination)
	}

	return nil
}

func (app *App) touchFiles(files []string) error {
	for _, file := range files {
		if err := app.touchFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) touchFile(file string) error {
	if stat, err := os.Stat(file); os.IsNotExist(err) {
		// File does not exist (ok)
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("%s is a directory", file)
	} else if !app.forceFlag {
		if err := updateTimestamp(file, time.Now()); err != nil {
			return err
		}
		app.logger.Printf("update timestamp of %s\n", file)
		return nil
	}

	source := app.resolveTemplatePath(file)
	templateExists := true
	if stat, err := os.Stat(source); os.IsNotExist(err) {
		templateExists = false
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("template %s is a directory", source)
	}

	if app.createDirFlag {
		if err := os.MkdirAll(path.Dir(file), 0755); err != nil {
			return err
		}
	}

	if templateExists {
		if err := renderTemplate(source, file); err != nil {
			return err
		}

		app.logger.Printf("create from template: %s -> %s\n", source, file)
		return nil
	} else {
		if err := createFile(file); err != nil {
			return err
		}

		app.logger.Printf("create empty file %s\n", file)
		return nil
	}
}

func (app *App) resolveTemplatePath(file string) string {
	basename := filepath.Base(file)

	return path.Join(app.rootDir, basename)
}
