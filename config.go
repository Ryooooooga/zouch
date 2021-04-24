package main

import (
	"os"
	"path"
)

const (
	rootEnvKey = "ZOUCH_ROOT"
)

func getRootDir() string {
	if root, ok := os.LookupEnv(rootEnvKey); ok {
		return root
	}
	if xdgConfigHome, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		return path.Join(xdgConfigHome, "zouch")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return path.Join(home, ".config", "zouch")
	}
	panic(err)
}
