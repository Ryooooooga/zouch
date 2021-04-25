package config

import (
	"testing"
)

func TestRootDir(t *testing.T) {
	type envs struct {
		home          string
		xdgConfigHome string
		zouchRoot     string
	}
	type scenario struct {
		name            string
		expectedRootDir string
		envs            envs
	}

	scenarios := []scenario{
		{
			name:            "default root dir",
			expectedRootDir: "/home/USER/.config/zouch",
			envs: envs{
				home:          "/home/USER",
				xdgConfigHome: "",
				zouchRoot:     "",
			},
		},
		{
			name:            "respect XDG_CONFIG_HOME",
			expectedRootDir: "/home/USER/xdgConfig/zouch",
			envs: envs{
				home:          "/home/USER",
				xdgConfigHome: "/home/USER/xdgConfig",
				zouchRoot:     "",
			},
		},
		{
			name:            "respect ZOUCH_ROOT",
			expectedRootDir: "/tmp/zouch",
			envs: envs{
				home:          "/home/USER",
				xdgConfigHome: "/home/USER/xdgConfig",
				zouchRoot:     "/tmp/zouch",
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			c := newTestConfig(s.envs.home, s.envs.xdgConfigHome, s.envs.zouchRoot)

			if rootDir := c.RootDir(); rootDir != s.expectedRootDir {
				t.Fatalf("expected rootDir == %v, actual %v", s.expectedRootDir, rootDir)
			}
		})
	}
}

func newTestConfig(home string, xdgConfigHome string, zouchRoot string) *Config {
	return &Config{
		// Stub of os.UserHomeDir
		userHomeDir: func() (string, error) {
			return home, nil
		},
		// Stub of os.LookupEnv
		lookupEnv: func(key string) (string, bool) {
			switch key {
			case XdgConfigHomeEnvKey:
				return xdgConfigHome, len(xdgConfigHome) > 0
			case RootEnvKey:
				return zouchRoot, len(zouchRoot) > 0
			default:
				return "", false
			}
		},
	}
}
