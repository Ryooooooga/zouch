package config

import (
	"os"
	"path"
)

const (
	rootEnvKey = "ZOUCH_ROOT"
)

func GetRootDir() string {
	// "$ZOUCH_ROOT"
	if root, ok := os.LookupEnv(rootEnvKey); ok {
		return root
	}
	// "$XDG_CONFIG_HOME/zouch"
	if xdgConfigHome, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		return path.Join(xdgConfigHome, "zouch")
	}
	// "$HOME/.config/zouch"
	home, err := os.UserHomeDir()
	if err != nil {
		return path.Join(home, ".config", "zouch")
	}
	panic(err)
}
