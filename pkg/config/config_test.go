package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootDir(t *testing.T) {
	type envs struct {
		home          string
		xdgConfigHome string
		zouchRoot     string
	}
	type scenario struct {
		name                string
		expectedRootDir     string
		expectedTemplateDir string
		envs                envs
	}

	scenarios := []scenario{
		{
			name:                "default root dir",
			expectedRootDir:     "/home/USER/.config/zouch",
			expectedTemplateDir: "/home/USER/.config/zouch/templates",
			envs: envs{
				home:          "/home/USER",
				xdgConfigHome: "",
				zouchRoot:     "",
			},
		},
		{
			name:                "respect XDG_CONFIG_HOME",
			expectedRootDir:     "/home/USER/xdgConfig/zouch",
			expectedTemplateDir: "/home/USER/xdgConfig/zouch/templates",
			envs: envs{
				home:          "/home/USER",
				xdgConfigHome: "/home/USER/xdgConfig",
				zouchRoot:     "",
			},
		},
		{
			name:                "respect ZOUCH_ROOT",
			expectedRootDir:     "/tmp/zouch",
			expectedTemplateDir: "/tmp/zouch/templates",
			envs: envs{
				home:          "/home/USER",
				xdgConfigHome: "/home/USER/xdgConfig",
				zouchRoot:     "/tmp/zouch",
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			setOrUnsetEnv(t, "HOME", s.envs.home)
			setOrUnsetEnv(t, "XDG_CONFIG_HOME", s.envs.xdgConfigHome)
			setOrUnsetEnv(t, "ZOUCH_ROOT", s.envs.zouchRoot)

			c := NewConfig()
			assert.NotNil(t, c)
			assert.Equal(t, s.expectedRootDir, c.RootDir)
			assert.Equal(t, s.expectedTemplateDir, c.TemplateDir)
		})
	}
}

func setOrUnsetEnv(t *testing.T, key string, value string) {
	var err error
	if len(value) > 0 {
		err = os.Setenv(key, value)
	} else {
		err = os.Unsetenv(key)
	}

	if err != nil {
		t.Fatalf("setOrUnsetEnv(%s, %s) failed %v", key, value, err)
	}
}
