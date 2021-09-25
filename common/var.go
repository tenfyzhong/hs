package common

import "github.com/urfave/cli/v2"

var SessionFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     FlagSessionWorkspace,
		Aliases:  []string{"w"},
		Required: true,
		Usage:    "the workspace to work",
	},
	&cli.StringFlag{
		Name:    FlagSessionSave,
		Aliases: []string{"s"},
		Usage:   "save session",
	},
	&cli.BoolFlag{
		Name:    FlagSessionList,
		Aliases: []string{"l"},
		Usage:   "list session",
	},
	&cli.StringFlag{
		Name:  FlagSessionRemove,
		Usage: "remove session",
	},
	&cli.StringFlag{
		Name:  FlagSessionShowPath,
		Usage: "show the path of the session",
	},
	&cli.StringFlag{
		Name:    FlagSessionReplay,
		Aliases: []string{"r"},
		Usage:   "replay session",
	},
	&cli.BoolFlag{
		Name:  FlagSessionHttpie,
		Usage: "use httpie to replay",
	},
	&cli.BoolFlag{
		Name:  FlagSessionCurl,
		Usage: "use curl to replay",
	},
	&cli.BoolFlag{
		Name:  FlagSessionRaw,
		Value: true,
		Usage: "print the http message",
	},
	&cli.BoolFlag{
		Name:  FlagSessionHTTPS,
		Usage: "replay as https",
	},
}

var WorkspaceFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    FlagWorkspaceCreate,
		Aliases: []string{"c"},
		Usage:   "create new a workspace",
	},
	&cli.StringFlag{
		Name:    FlagWorkspaceRemove,
		Aliases: []string{"r"},
		Usage:   "remove a workspace",
	},
	&cli.BoolFlag{
		Name:    FlagWorkspaceList,
		Aliases: []string{"l"},
		Usage:   "list workspace",
	},
	&cli.StringFlag{
		Name:  FlagWorkspaceShowPath,
		Usage: "show the path of the workspace",
	},
}
