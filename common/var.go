package common

import "github.com/urfave/cli/v2"

var SessionFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    FlagSessionSave,
		Aliases: []string{"s"},
		Usage:   "Save session",
	},
	&cli.StringFlag{
		Name:    FlagSessionReplay,
		Aliases: []string{"r"},
		Usage:   "Replay session",
	},
	&cli.BoolFlag{
		Name:    FlagSessionList,
		Aliases: []string{"l"},
		Usage:   "List session",
	},
	&cli.StringFlag{
		Name:  FlagSessionRemove,
		Usage: "Remove session",
	},
	&cli.StringFlag{
		Name:  FlagSessionShowPath,
		Usage: "Show the path of the session",
	},
	&cli.StringFlag{
		Name:     FlagSessionWorkspace,
		Aliases:  []string{"w"},
		Required: true,
		Usage:    "The workspace to work",
	},
	&cli.BoolFlag{
		Name:  FlagSessionHttpie,
		Usage: "Replay as httpie command",
	},
	&cli.BoolFlag{
		Name:  FlagSessionCurl,
		Usage: "Replay as curl command",
	},
	&cli.BoolFlag{
		Name:  FlagSessionRaw,
		Value: true,
		Usage: "Replay as raw http message",
	},
	&cli.BoolFlag{
		Name:  FlagSessionHTTPS,
		Usage: "Replay as https",
	},
}

var WorkspaceFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    FlagWorkspaceCreate,
		Aliases: []string{"c"},
		Usage:   "Create new workspace",
	},
	&cli.StringFlag{
		Name:    FlagWorkspaceRemove,
		Aliases: []string{"r"},
		Usage:   "Remove a workspace",
	},
	&cli.BoolFlag{
		Name:    FlagWorkspaceList,
		Aliases: []string{"l"},
		Usage:   "List workspace",
	},
	&cli.StringFlag{
		Name:  FlagWorkspaceShowPath,
		Usage: "Show the path of the workspace",
	},
}

