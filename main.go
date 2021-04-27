package main

import (
	"fmt"
	"os"

	"github.com/Ryooooooga/zouch/pkg/cli"
	"github.com/fatih/color"
)

var (
	version = "dev"
	commit  = "HEAD"
	date    = "unknown"
)

func main() {
	if err := cli.New(fmt.Sprintf("%s (rev: %s) built at %s", version, commit, date)).Run(os.Args); err != nil {
		color.New(color.FgRed).Fprintf(os.Stderr, "zouch: %s\n", err)
	}
}
