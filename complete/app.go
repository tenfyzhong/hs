package complete

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func App(c *cli.Context) {
	args := appArgs()
	nargs := len(args)

	prefix := ""
	if nargs > 0 {
		prefix = args[nargs-1]
	}
	commands := c.App.VisibleCommands()
	flags := c.App.VisibleFlags()

	items := make([]Item, 0, len(flags)+len(commands))
	items = append(items, CommandItems(commands)...)
	items = append(items, FlagItems(flags)...)

	matchItems := ItemsHasPrefix(items, prefix)
	if len(matchItems) == 1 && matchItems[0].Name == prefix {
		return
	}

	MatchItemsSort(matchItems)
	for _, item := range matchItems {
		fmt.Println(item.String())
	}
}
