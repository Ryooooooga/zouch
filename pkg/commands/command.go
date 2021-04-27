package commands

import (
	"io"
	"log"
)

type Command struct {
	output    io.Writer
	logger    *log.Logger
	rootDir   string
	createDir bool
	force     bool
}

func NewCommand(output io.Writer, logger *log.Logger, rootDir string, createDir bool, force bool) *Command {
	return &Command{
		output,
		logger,
		rootDir,
		createDir,
		force,
	}
}
