package complete

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

type Item struct {
	Names []string
	Info  string
}

type MatchItem struct {
	Name string
	Info string
}

func (i MatchItem) String() string {
	if i.Info == "" {
		return i.Name
	}
	return i.Name + ":" + i.Info
}

func FlagItem(flag cli.DocGenerationFlag) Item {
	names := flag.Names()
	strLenSort(names)

	usage := flag.GetUsage()
	for i, name := range names {
		if len(name) == 1 {
			names[i] = "-" + name
		} else {
			names[i] = "--" + name
		}
	}
	return Item{
		Names: names,
		Info:  usage,
	}
}

func FlagItems(flags []cli.Flag) []Item {
	res := make([]Item, 0, len(flags))
	for _, flag := range flags {
		t, _ := flag.(cli.DocGenerationFlag)
		res = append(res, FlagItem(t))
	}
	return res
}

func CommandItem(command *cli.Command) Item {
	names := command.Names()
	strLenSort(names)

	usage := command.Usage
	return Item{
		Names: names,
		Info:  usage,
	}
}

func CommandItems(commands []*cli.Command) []Item {
	res := make([]Item, 0, len(commands))
	for _, command := range commands {
		res = append(res, CommandItem(command))
	}
	return res
}

func ItemsHasPrefix(items []Item, prefix string) []MatchItem {
	res := make([]MatchItem, 0)
	for _, item := range items {
		for _, name := range item.Names {
			if strings.HasPrefix(name, prefix) {
				res = append(res, MatchItem{
					Name: name,
					Info: item.Info,
				})
				break
			}
		}
	}
	return res
}

func MatchItemsSort(items []MatchItem) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Name < items[j].Name
	})
}

func strLenSort(strs []string) {
	sort.Slice(strs, func(i, j int) bool {
		if len(strs[i]) == len(strs[j]) {
			return strs[i] < strs[j]
		}
		return len(strs[i]) < len(strs[j])
	})
}

func commandArgs() []string {
	nargs := len(os.Args)
	args := os.Args[2 : nargs-1]
	return args
}

func appArgs() []string {
	nargs := len(os.Args)
	args := os.Args[1 : nargs-1]
	return args
}

func isFlag(flags []cli.Flag, flagName, name string) bool {
	var specialFlag cli.Flag
	for _, flag := range flags {
		if flag.Names()[0] == flagName {
			specialFlag = flag
			break
		}
	}
	if specialFlag != nil {
		trimmed := strings.TrimLeft(name, "-")
		for _, name := range specialFlag.Names() {
			if name == trimmed {
				return true
			}
		}
	}
	return false
}

func isFlags(flags []cli.Flag, testFlagNames []string, name string) bool {
	for _, flagName := range testFlagNames {
		if isFlag(flags, flagName, name) {
			return true
		}
	}
	return false
}

func completeCommand(flags []cli.Flag, resultFunc func(lastArg string) []string) {
	args := commandArgs()

	lastArg := ""
	if len(args) > 0 {
		lastArg = args[len(args)-1]
	}

	items := FlagItems(flags)
	matchItems := ItemsHasPrefix(items, lastArg)

	if len(matchItems) == 1 && matchItems[0].Name == lastArg {
		result := resultFunc(lastArg)
		for _, str := range result {
			fmt.Println(str)
		}
		return
	}

	MatchItemsSort(matchItems)
	for _, item := range matchItems {
		fmt.Println(item.String())
	}
}
