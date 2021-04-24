package main

import "log"

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
	return nil
}

func (app *App) touchFiles(files []string) error {
	return nil
}
