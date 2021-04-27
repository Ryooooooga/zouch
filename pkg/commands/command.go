package commands

import (
	"io"
	"log"

	"github.com/Ryooooooga/zouch/pkg/repositories"
)

type Command struct {
	output    io.Writer
	logger    *log.Logger
	templates repositories.TemplateRepository
	createDir bool
	force     bool
}

func NewCommand(output io.Writer, logger *log.Logger, templates repositories.TemplateRepository, createDir bool, force bool) *Command {
	return &Command{
		output,
		logger,
		templates,
		createDir,
		force,
	}
}
