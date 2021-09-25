package complete

import (
	"github.com/tenfyzhong/hs/common"
	"github.com/urfave/cli/v2"
)

func Workspace(c *cli.Context) {
	completeCommand(c.Command.Flags, func(lastArg string) []string {
		testFlagNames := []string{common.FlagWorkspaceCreate, common.FlagSessionList}
		if isFlags(c.Command.Flags, testFlagNames, lastArg) {
			return nil
		}
		workspaces, err := common.GetWorkspaces(common.GetDir(c))
		if err != nil {
			return nil
		}
		return workspaces
	})
}
