package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Ryooooooga/zouch/pkg/config"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "HEAD"
	date    = "unknown"
)

func main() {
	if err := newCliApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newCliApp() *cli.App {
	authors := []*cli.Author{
		{Name: "Ryooooooga", Email: "eial5q265e5+github@gmail.com"},
	}

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:     "add",
			Aliases:  []string{"A"},
			Usage:    "add [files...] as templates",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "",
			Aliases:  []string{"r"},
			Usage:    "create directories as required",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "force",
			Aliases:  []string{"f"},
			Usage:    "force update existing files",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "verbose",
			Aliases:  []string{"V"},
			Usage:    "display verbose output",
			Required: false,
		},
	}

	return &cli.App{
		Name:                   "zouch",
		Usage:                  "Create a new file from a template",
		UsageText:              "zouch [files...]\n   zouch --add [files...]",
		Version:                fmt.Sprintf("%s, rev: %s, built at %s", version, commit, date),
		Flags:                  flags,
		HideHelpCommand:        true,
		Action:                 runCommand,
		Authors:                authors,
		UseShortOptionHandling: true,
	}
}

func runCommand(c *cli.Context) error {
	addFlag := c.Bool("add")
	createDirFlag := c.Bool("r")
	forceFlag := c.Bool("force")
	verboseFlag := c.Bool("verbose")
	files := c.Args().Slice()

	if len(files) == 0 {
		cli.ShowAppHelpAndExit(c, 1)
	}

	logger := newLogger(verboseFlag)
	rootDir := config.NewConfig().RootDir()

	app := newApp(logger, rootDir, createDirFlag, forceFlag)

	if addFlag {
		return app.addTemplateFiles(files)
	} else {
		return app.touchFiles(files)
	}
}

func newLogger(verbose bool) *log.Logger {
	writer := io.Discard
	if verbose {
		writer = os.Stderr
	}
	return log.New(writer, "", 0)
}
