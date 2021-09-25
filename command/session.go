package command

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	workspace := c.String(common.FlagSessionWorkspace)
	if workspace == "" {
		return cli.Exit("--workspace is required", common.CodeFlagRequired)
	}

	workspacePath := filepath.Join(dir, workspace)
	if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
		return cli.Exit(workspacePath+" is not exist", common.CodeNotExist)
	}

	var err error
	if c.IsSet(common.FlagSessionSave) {
		name := c.String(common.FlagSessionSave)
		if name == "" {
			return cli.Exit("save session name is required", common.CodeFlagRequired)
		}
		err = saveSession(workspacePath, name)
	} else if c.IsSet(common.FlagSessionReplay) {
		name := c.String(common.FlagSessionReplay)
		if name == "" {
			return cli.Exit("replay session name is required", common.CodeFlagRequired)
		}
		replayType := getReplayType(c)
		isHTTPS := c.Bool(common.FlagSessionHTTPS)
		args := c.Args().Slice()
		err = replaySession(workspacePath, name, replayType, isHTTPS, args)
	} else if c.IsSet(common.FlagSessionList) {
		err = listSession(workspacePath)
	} else if c.IsSet(common.FlagSessionRemove) {
		name := c.String(common.FlagSessionRemove)
		if name == "" {
			return cli.Exit("remove session name is required", common.CodeFlagRequired)
		}
		err = removeSession(workspacePath, name)
	} else if c.IsSet(common.FlagSessionShowPath) {
		name := c.String(common.FlagSessionShowPath)
		if name == "" {
			return cli.Exit("show-path session name is required", common.CodeFlagRequired)
		}
		err = showPathSession(workspacePath, name)
	} else {
		err = cli.Exit("", common.CodeUnknownFlag)
	}
	return err
}

func getReplayType(c *cli.Context) ReplayType {
	replayType := replayRaw
	if c.Bool(common.FlagSessionHttpie) {
		replayType = replayHttpie
	} else if c.Bool(common.FlagSessionCurl) {
		replayType = replayCurl
	}
	return replayType
}

func saveSession(workspacePath, name string) error {
	name = strings.ReplaceAll(name, "/", "-")
	name += common.SessionSuffix
	path := filepath.Join(workspacePath, name)
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func replaySession(workspacePath, name string, replayType ReplayType, isHTTPS bool, args []string) error {
	name += common.SessionSuffix
	path := filepath.Join(workspacePath, name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	r := bufio.NewReader(bytes.NewReader(data))
	req, err := http.ReadRequest(r)
	if err != nil {
		return err
	}

	msg := ""
	switch replayType {
	case replayRaw:
		msg = string(data)
	case replayHttpie:
		msg, err = buildHttpieCmd(req, isHTTPS, args)
	case replayCurl:
		msg, err = buildCurlCmd(req, isHTTPS, args)
	}
	fmt.Printf("%s", msg)
	return err
}

func listSession(workspacePath string) error {
	res, err := common.GetSessions(workspacePath)
	if err != nil {
		return err
	}
	for _, name := range res {
		fmt.Println(name)
	}
	return nil
}

func removeSession(workspacePath, name string) error {
	name += common.SessionSuffix
	path := filepath.Join(workspacePath, name)
	return os.Remove(path)
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
