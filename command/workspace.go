package command

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/tenfyzhong/hs/common"
	"github.com/urfave/cli/v2"
)

func Workspace(c *cli.Context) error {
	dir := common.GetDir(c)

	var err error
	if c.IsSet(common.FlagCreate) {
		name := c.String(common.FlagCreate)
		if name == "" {
			return cli.Exit("workspace name is required", common.CodeFlagRequired)
		}
		err = createWorkspace(dir, name)
	} else if c.IsSet(common.FlagRemove) {
		name := c.String(common.FlagRemove)
		if name == "" {
			return cli.Exit("workspace name is required", common.CodeFlagRequired)
		}
		err = removeWorkspace(dir, name)
	} else if c.IsSet(common.FlagList) {
		err = listWorkspace(dir)
	} else if c.IsSet(common.FlagShowPath) {
		name := c.String(common.FlagShowPath)
		if name == "" {
			return cli.Exit("workspace name is required", common.CodeFlagRequired)
		}
		err = showPathWorkspace(dir, name)
	} else {
		err = cli.Exit("", common.CodeUnknownFlag)
	}

	return errors.Wrapf(err, "workspace")
}

func createWorkspace(dir, name string) error {
	name = strings.ReplaceAll(name, "/", "-")
	path := filepath.Join(dir, name)
	err := os.MkdirAll(path, fs.ModePerm)
	return errors.Wrapf(err, "MkdirAll %s", path)
}

func removeWorkspace(dir, name string) error {
	path := filepath.Join(dir, name)
	err := os.RemoveAll(path)
	return errors.Wrapf(err, "RemoveAll %s", path)
}

func listWorkspace(dir string) error {
	workspaces, err := common.GetWorkspaces(dir)
	if err != nil {
		return errors.Wrapf(err, "GetWorkspaces %s", dir)
	}
	for _, name := range workspaces {
		fmt.Println(name)
	}
	return nil
}

func showPathWorkspace(dir, name string) error {
	path := filepath.Join(dir, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cli.Exit(fmt.Sprintf("%s is not exist", path), common.CodeNotExist)
	}
	fmt.Println(path)
	return nil
}
