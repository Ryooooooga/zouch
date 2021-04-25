package config

import (
	"os"
	"path"
)

const (
	RootEnvKey          = "ZOUCH_ROOT"
	XdgConfigHomeEnvKey = "XDG_COFNIG_HOME"
)

type Config struct {
	userHomeDir func() (string, error)
	lookupEnv   func(key string) (string, bool)
}

func NewConfig() *Config {
	return &Config{
		userHomeDir: os.UserHomeDir,
		lookupEnv:   os.LookupEnv,
	}
}

func (c *Config) RootDir() string {
	// "$ZOUCH_ROOT"
	if root, ok := c.lookupEnv(RootEnvKey); ok {
		return root
	}
	// "$XDG_CONFIG_HOME/zouch"
	if xdgConfigHome, ok := c.lookupEnv(XdgConfigHomeEnvKey); ok {
		return path.Join(xdgConfigHome, "zouch")
	}
	// "$HOME/.config/zouch"
	home, err := c.userHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".config", "zouch")
}
