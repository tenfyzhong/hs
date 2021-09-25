package main

import (
	"log"
	"os"

	"github.com/tenfyzhong/hs/command"
	"github.com/tenfyzhong/hs/common"
	"github.com/tenfyzhong/hs/complete"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "hs",
		Usage:                "A tool to enhance httpie/curl",
		Version:              "v0.1.0",
		EnableBashCompletion: true,
		BashComplete:         complete.App,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
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
		log.Fatal(err)
	}
}
