package commands

import (
	"io"
	"log"
	"time"

	"github.com/Ryooooooga/zouch/pkg/config"
	"github.com/Ryooooooga/zouch/pkg/renderer"
	"github.com/Ryooooooga/zouch/pkg/repositories"
)

type Command struct {
	Output    io.Writer
	Logger    *log.Logger
	Config    *config.Config
	Templates repositories.TemplateRepository
	Renderer  renderer.TemplateRenderer
	CreateDir bool
	Force     bool

	Now func() time.Time
}

func NewCommand(
	output io.Writer,
	logger *log.Logger,
	config *config.Config,
	templates repositories.TemplateRepository,
	renderer renderer.TemplateRenderer,
	createDir bool,
	force bool,
) *Command {
	return &Command{
		Output:    output,
		Logger:    logger,
		Config:    config,
		Templates: templates,
		Renderer:  renderer,
		CreateDir: createDir,
		Force:     force,

		Now: time.Now,
	}
}
