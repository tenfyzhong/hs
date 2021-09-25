package command

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/tenfyzhong/hs/common"
	"github.com/urfave/cli/v2"
)

func Workspace(c *cli.Context) error {
	dir := common.GetDir(c)

	var err error
	if c.IsSet(common.FlagWorkspaceCreate) {
		name := c.String(common.FlagWorkspaceCreate)
		if name == "" {
			return cli.Exit("workspace name is required", common.CodeFlagRequired)
		}
		err = createWorkspace(dir, name)
	} else if c.IsSet(common.FlagWorkspaceRemove) {
		name := c.String(common.FlagWorkspaceRemove)
		if name == "" {
			return cli.Exit("workspace name is required", common.CodeFlagRequired)
		}
		err = removeWorkspace(dir, name)
	} else if c.IsSet(common.FlagWorkspaceList) {
		err = listWorkspace(dir)
	} else if c.IsSet(common.FlagWorkspaceShowPath) {
		name := c.String(common.FlagWorkspaceShowPath)
		if name == "" {
			return cli.Exit("workspace name is required", common.CodeFlagRequired)
		}
		err = showPathWorkspace(dir, name)
	} else {
		err = cli.Exit("", common.CodeUnknownFlag)
	}

	return err
}

func createWorkspace(dir, name string) error {
	name = strings.ReplaceAll(name, "/", "-")
	path := filepath.Join(dir, name)
	return os.MkdirAll(path, fs.ModePerm)
}

func removeWorkspace(dir, name string) error {
	path := filepath.Join(dir, name)
	return os.RemoveAll(path)
}

func listWorkspace(dir string) error {
	workspaces, err := common.GetWorkspaces(dir)
	if err != nil {
		return err
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
