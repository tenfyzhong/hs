package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/tenfyzhong/hs/command"
	"github.com/tenfyzhong/hs/common"
	"github.com/tenfyzhong/hs/complete"
	"github.com/urfave/cli/v2"
)

var enableLog bool

func init() {
	enableLogStr := os.Getenv("ENABLE_HS_LOG")
	if enableLogStr == "1" {
		if w, err := os.OpenFile("hs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
			enableLog = true
			log.SetOutput(w)
			log.SetFlags(log.Lshortfile | log.LstdFlags)
		}
	}
}

func main() {
	app := &cli.App{
		Name:                 "hs",
		Usage:                "A tool to enhance httpie/curl",
		Version:              "v0.1.0",
		EnableBashCompletion: true,
		BashComplete:         complete.App,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    common.FlagAppDir,
				Value:   common.DefaultDir(),
				Usage:   "the directory to store the data",
				Aliases: []string{"d"},
				EnvVars: []string{"HS_DIR"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:         "session",
				Usage:        "Manage session and replay session",
				Aliases:      []string{"s"},
				Action:       command.Session,
				Flags:        common.SessionFlags,
				BashComplete: complete.Session,
			},
			{
				Name:         "workspace",
				Usage:        "Manage workspace",
				Aliases:      []string{"w"},
				Action:       command.Workspace,
				Flags:        common.WorkspaceFlags,
				BashComplete: complete.Workspace,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		if enableLog {
			log.Printf("cmd:%v err:%+v", os.Args, err)
		}
		fmt.Fprintf(app.ErrWriter, "%v", errors.Cause(err).Error())
		os.Exit(common.CodeUnknown)
	}
}
