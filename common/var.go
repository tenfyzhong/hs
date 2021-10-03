package common

import "github.com/urfave/cli/v2"

var SessionFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     FlagWorkspace,
		Aliases:  []string{"w"},
		Required: true,
		Usage:    "special the workspace to work",
	},
	&cli.StringFlag{
		Name:    FlagCreate,
		Aliases: []string{"c"},
		Usage:   "create a session into the workspace, suggest to get the http message by call the 'httpie --offline' command",
	},
	&cli.BoolFlag{
		Name:    FlagList,
		Aliases: []string{"l"},
		Usage:   "list the sessions in the workspace",
	},
	&cli.StringFlag{
		Name:    FlagRemove,
		Aliases: []string{"R"},
		Usage:   "remove a session in the workspace",
	},
	&cli.StringFlag{
		Name:    FlagShowPath,
		Aliases: []string{"P"},
		Usage:   "show the path of the session in the workspace",
	},
	&cli.StringFlag{
		Name:    FlagReplay,
		Aliases: []string{"r"},
		Usage:   "replay a session in the workspace, the rest args will be add to the 'httpie'/'curl' command",
	},
	&cli.BoolFlag{
		Name:    FlagHttpie,
		Aliases: []string{"H"},
		Usage:   "use httpie to replay the session, only works with '--replay'",
	},
	&cli.BoolFlag{
		Name:    FlagCurl,
		Aliases: []string{"C"},
		Usage:   "use curl to replay the session, only works with '--replay'",
	},
	&cli.BoolFlag{
		Name:  FlagRaw,
		Value: true,
		Usage: "print the raw http message, only works with '--replay'",
	},
	&cli.BoolFlag{
		Name:    FlagHTTPS,
		Aliases: []string{"S"},
		Usage:   "replay as https, only works with '--httpie' or '--curl'",
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

var TransferFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    FlagHttpie,
		Aliases: []string{"H"},
		Usage:   "transfer to httpie command",
	},
	&cli.BoolFlag{
		Name:    FlagCurl,
		Aliases: []string{"C"},
		Value:   true,
		Usage:   "transfer to curl command",
	},
	&cli.BoolFlag{
		Name:    FlagHTTPS,
		Aliases: []string{"S"},
		Usage:   "use https",
	},
}
