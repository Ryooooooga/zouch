package cli

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Ryooooooga/zouch/pkg/commands"
	"github.com/Ryooooooga/zouch/pkg/config"
	"github.com/Ryooooooga/zouch/pkg/errors"
	"github.com/Ryooooooga/zouch/pkg/renderer"
	"github.com/Ryooooooga/zouch/pkg/repositories"
	"github.com/urfave/cli/v2"
)

func New(version string) *cli.App {
	var authors = []*cli.Author{
		{Name: "Ryooooooga", Email: "eial5q265e5+github@gmail.com"},
	}

	var flags = []cli.Flag{
		&cli.BoolFlag{
			Name:     "list",
			Aliases:  []string{"l"},
			Usage:    "list template files",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "preview",
			Aliases:  []string{"p"},
			Usage:    "show template preview",
			Required: false,
		},
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
zouch --list
zouch --preview [files...]
zouch --add     [files...]`

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
	listFlag := ctx.Bool("list")
	previewFlag := ctx.Bool("preview")
	addFlag := ctx.Bool("add")

	createDirFlag := ctx.Bool("r")
	forceFlag := ctx.Bool("force")
	verboseFlag := ctx.Bool("verbose")
	files := ctx.Args().Slice()

	output := os.Stdout
	logger := newLogger(verboseFlag)

	cfg := config.NewConfig()
	templates := repositories.NewTemplateRepository(cfg.TemplateDir)
	renderer := renderer.NewTextTemplateRenderer()

	cmd := commands.NewCommand(
		output,
		logger,
		cfg,
		templates,
		renderer,
		createDirFlag,
		forceFlag,
	)

	var err error
	if listFlag {
		err = cmd.List(files)
	} else if previewFlag {
		err = cmd.Preview(files)
	} else if addFlag {
		err = cmd.Add(files)
	} else {
		err = cmd.Touch(files)
	}

	if errors.IsShowHelpAndExitError(err) {
		err = cli.ShowAppHelp(ctx)
		fmt.Fprintln(ctx.App.Writer)
	}

	return err
}

func newLogger(verbose bool) *log.Logger {
	writer := io.Discard
	if verbose {
		writer = os.Stderr
	}
	return log.New(writer, "", 0)
}
