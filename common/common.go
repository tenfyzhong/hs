package common

import (
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

func GetDir(c *cli.Context) string {
	dir := c.String(FlagAppDir)
	if dir == "" {
		dir = DefaultDir()
	}
	return dir
}

func DefaultDir() string {
	home := os.Getenv("HOME")
	return path.Join(home, ".hs")
}
