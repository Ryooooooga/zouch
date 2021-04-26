package cli

import (
	"io"
	"log"
	"os"

	"github.com/Ryooooooga/zouch/pkg/app"
	"github.com/Ryooooooga/zouch/pkg/config"
	"github.com/urfave/cli/v2"
)

type Cli struct {
	cliApp *cli.App
}

func New(version string) *Cli {
	var authors = []*cli.Author{
		{Name: "Ryooooooga", Email: "eial5q265e5+github@gmail.com"},
	}

	var flags = []cli.Flag{
		&cli.BoolFlag{
			Name:     "add",
			Aliases:  []string{"A"},
			Usage:    "add [files...] as new templates",
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

	cliApp := &cli.App{
		Name:                   "zouch",
		Usage:                  "Create a new file from a template",
		UsageText:              "zouch [files...]\n   zouch --add [files...]",
		Version:                version,
		Flags:                  flags,
		HideHelpCommand:        true,
		Authors:                authors,
		UseShortOptionHandling: true,
	}

	c := &Cli{
		cliApp,
	}

	c.cliApp.Action = c.runCommand

	return c
}

func (c *Cli) Run(args []string) error {
	return c.cliApp.Run(args)
}

func (c *Cli) runCommand(ctx *cli.Context) error {
	addFlag := ctx.Bool("add")
	createDirFlag := ctx.Bool("r")
	forceFlag := ctx.Bool("force")
	verboseFlag := ctx.Bool("verbose")
	files := ctx.Args().Slice()

	if len(files) == 0 {
		cli.ShowAppHelpAndExit(ctx, 1)
	}

	logger := newLogger(verboseFlag)
	rootDir := config.NewConfig().RootDir()

	a := app.NewApp(logger, rootDir, createDirFlag, forceFlag)

	if addFlag {
		return a.AddTemplateFiles(files)
	} else {
		return a.TouchFiles(files)
	}
}

func newLogger(verbose bool) *log.Logger {
	writer := io.Discard
	if verbose {
		writer = os.Stderr
	}
	return log.New(writer, "", 0)
}
