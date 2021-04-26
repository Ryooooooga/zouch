package main

import (
	"fmt"
	"os"

	"github.com/Ryooooooga/zouch/pkg/cli"
)

var (
	version = "dev"
	commit  = "HEAD"
	date    = "unknown"
)

func main() {
	if err := cli.New(fmt.Sprintf("%s (rev: %s) built at %s", version, commit, date)).Run(os.Args); err != nil {
		panic(err)
	}
}
