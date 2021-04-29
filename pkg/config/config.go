package config

import (
	"os"
	"path"
)

type Config struct {
	RootDir     string
	TemplateDir string
}

func NewConfig() *Config {
	rootDir := getRootDir()
	templateDir := path.Join(rootDir, "templates")

	return &Config{
		RootDir:     rootDir,
		TemplateDir: templateDir,
	}
}

func getRootDir() string {
	// "$ZOUCH_ROOT"
	if root, ok := os.LookupEnv("ZOUCH_ROOT"); ok {
		return root
	}
	// "$XDG_CONFIG_HOME/zouch"
	if xdgConfigHome, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		return path.Join(xdgConfigHome, "zouch")
	}
	// "$HOME/.config/zouch"
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".config", "zouch")
}
