package cli

import (
	"io"
	"log"
	"os"

	"github.com/Ryooooooga/zouch/pkg/app"
	"github.com/Ryooooooga/zouch/pkg/config"
	"github.com/urfave/cli/v2"
)

func New(version string) *cli.App {
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

	const usage = `zouch [files...]
   zouch --add [files...]`

	return &cli.App{
		Name:                   "zouch",
		Usage:                  "Create a new file from a template",
		UsageText:              usage,
		Version:                version,
		Flags:                  flags,
		HideHelpCommand:        true,
		Action:                 runCommand,
		Authors:                authors,
		UseShortOptionHandling: true,
	}
}

func runCommand(ctx *cli.Context) error {
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

	app := app.NewApp(logger, rootDir, createDirFlag, forceFlag)

	if addFlag {
		return app.AddTemplateFiles(files)
	} else {
		return app.TouchFiles(files)
	}
}

func newLogger(verbose bool) *log.Logger {
	writer := io.Discard
	if verbose {
		writer = os.Stderr
	}
	return log.New(writer, "", 0)
}
