package complete

import (
	"path/filepath"

	"github.com/tenfyzhong/hs/common"
	"github.com/urfave/cli/v2"
)

func Session(c *cli.Context) {
	completeCommand(c.Command.Flags, func(lastArg string) []string {
		testFlagNames := []string{
			common.FlagHttpie,
			common.FlagCreate,
			common.FlagCurl,
			common.FlagRaw,
			common.FlagList,
			common.FlagHTTPS}
		if isFlags(c.Command.Flags, testFlagNames, lastArg) {
			return nil
		}

		if isFlag(c.Command.Flags, common.FlagWorkspace, lastArg) {
			workspaces, _ := common.GetWorkspaces(common.GetDir(c))
			return workspaces
		}

		workspace := c.String(common.FlagWorkspace)
		if workspace == "" {
			return nil
		}

		dir := common.GetDir(c)
		workspacePath := filepath.Join(dir, workspace)
		sessions, _ := common.GetSessions(workspacePath)
		return sessions
	})
}
