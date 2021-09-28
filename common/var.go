package common

import "github.com/urfave/cli/v2"

var SessionFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     FlagWorkspace,
		Aliases:  []string{"w"},
		Required: true,
		Usage:    "the workspace to work",
	},
	&cli.StringFlag{
		Name:    FlagCreate,
		Aliases: []string{"c"},
		Usage:   "create session",
	},
	&cli.BoolFlag{
		Name:    FlagList,
		Aliases: []string{"l"},
		Usage:   "list session",
	},
	&cli.StringFlag{
		Name:    FlagRemove,
		Aliases: []string{"R"},
		Usage:   "remove session",
	},
	&cli.StringFlag{
		Name:    FlagShowPath,
		Aliases: []string{"P"},
		Usage:   "show the path of the session",
	},
	&cli.StringFlag{
		Name:    FlagReplay,
		Aliases: []string{"r"},
		Usage:   "replay session",
	},
	&cli.BoolFlag{
		Name:    FlagHttpie,
		Aliases: []string{"H"},
		Usage:   "use httpie to replay",
	},
	&cli.BoolFlag{
		Name:    FlagCurl,
		Aliases: []string{"C"},
		Usage:   "use curl to replay",
	},
	&cli.BoolFlag{
		Name:  FlagRaw,
		Value: true,
		Usage: "print the http message",
	},
	&cli.BoolFlag{
		Name:    FlagHTTPS,
		Aliases: []string{"S"},
		Usage:   "replay as https",
	},
}

var WorkspaceFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    FlagCreate,
		Aliases: []string{"c"},
		Usage:   "create a workspace",
	},
	&cli.StringFlag{
		Name:    FlagRemove,
		Aliases: []string{"R"},
		Usage:   "remove a workspace",
	},
	&cli.BoolFlag{
		Name:    FlagList,
		Aliases: []string{"l"},
		Usage:   "list workspace",
	},
	&cli.StringFlag{
		Name:    FlagShowPath,
		Aliases: []string{"P"},
		Usage:   "show the path of the workspace",
	},
}
