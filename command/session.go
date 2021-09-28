package command

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/tenfyzhong/hs/common"
	"github.com/urfave/cli/v2"
)

type ReplayType int

const (
	replayRaw ReplayType = iota
	replayHttpie
	replayCurl
)

func Session(c *cli.Context) error {
	dir := common.GetDir(c)
	workspace := c.String(common.FlagWorkspace)
	if workspace == "" {
		return cli.Exit("--workspace is required", common.CodeFlagRequired)
	}

	workspacePath := filepath.Join(dir, workspace)
	if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
		return cli.Exit(workspacePath+" is not exist", common.CodeNotExist)
	}

	var err error
	if c.IsSet(common.FlagCreate) {
		name := c.String(common.FlagCreate)
		if name == "" {
			return cli.Exit("create session name is required", common.CodeFlagRequired)
		}
		err = createSession(workspacePath, name)
	} else if c.IsSet(common.FlagReplay) {
		name := c.String(common.FlagReplay)
		if name == "" {
			return cli.Exit("replay session name is required", common.CodeFlagRequired)
		}
		replayType := getReplayType(c)
		isHTTPS := c.Bool(common.FlagHTTPS)
		args := c.Args().Slice()
		err = replaySession(workspacePath, name, replayType, isHTTPS, args)
	} else if c.IsSet(common.FlagList) {
		err = listSession(workspacePath)
	} else if c.IsSet(common.FlagRemove) {
		name := c.String(common.FlagRemove)
		if name == "" {
			return cli.Exit("remove session name is required", common.CodeFlagRequired)
		}
		err = removeSession(workspacePath, name)
	} else if c.IsSet(common.FlagShowPath) {
		name := c.String(common.FlagShowPath)
		if name == "" {
			return cli.Exit("show-path session name is required", common.CodeFlagRequired)
		}
		err = showPathSession(workspacePath, name)
	} else {
		err = cli.Exit("", common.CodeUnknownFlag)
	}
	return errors.WithStack(err)
}

func getReplayType(c *cli.Context) ReplayType {
	replayType := replayRaw
	if c.Bool(common.FlagHttpie) {
		replayType = replayHttpie
	} else if c.Bool(common.FlagCurl) {
		replayType = replayCurl
	}
	return replayType
}

func createSession(workspacePath, name string) error {
	name = strings.ReplaceAll(name, "/", "-")
	name += common.SessionSuffix
	path := filepath.Join(workspacePath, name)
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return errors.Wrapf(err, "io.ReadAll from os.Stdin")
	}
	err = ioutil.WriteFile(path, data, 0644)
	return errors.Wrapf(err, "ioutil.WriteFile %s", path)
}

func replaySession(workspacePath, name string, replayType ReplayType, isHTTPS bool, args []string) error {
	name += common.SessionSuffix
	path := filepath.Join(workspacePath, name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrapf(err, "ioutil.ReadFile %s", path)
	}

	req, err := buildHTTPRequest(data)
	if err != nil {
		return errors.Wrapf(err, "buildHttpRequest")
	}

	msg := ""
	switch replayType {
	case replayRaw:
		msg = string(data)
	case replayHttpie:
		msg, err = buildHttpieCmd(req, isHTTPS, args)
		if err != nil {
			return errors.Wrapf(err, "buildHttpieCmd data:%s", string(data))
		}
	case replayCurl:
		msg, err = buildCurlCmd(req, isHTTPS, args)
		if err != nil {
			return errors.Wrapf(err, "buildCurlCmd data:%s", string(data))
		}
	}
	fmt.Printf("%s", msg)
	return nil
}

func listSession(workspacePath string) error {
	res, err := common.GetSessions(workspacePath)
	if err != nil {
		return errors.Wrapf(err, "GetSessions workspacePath: %s", workspacePath)
	}
	for _, name := range res {
		fmt.Println(name)
	}
	return nil
}

func removeSession(workspacePath, name string) error {
	name += common.SessionSuffix
	path := filepath.Join(workspacePath, name)
	err := os.Remove(path)
	return errors.Wrapf(err, "os.Remove %s", path)
}

func showPathSession(workspacePath, name string) error {
	name += common.SessionSuffix
	path := filepath.Join(workspacePath, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cli.Exit(fmt.Sprintf("%s is not exist", path), common.CodeNotExist)
	}
	fmt.Println(path)
	return nil
}
