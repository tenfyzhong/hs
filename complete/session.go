package complete

import (
	"path/filepath"

	"github.com/tenfyzhong/hs/common"
	"github.com/urfave/cli/v2"
)

func Session(c *cli.Context) {
	completeCommand(c.Command.Flags, func(lastArg string) []string {
		testFlagNames := []string{
			common.FlagSessionHttpie,
			common.FlagSessionSave,
			common.FlagSessionCurl,
			common.FlagSessionRaw,
			common.FlagSessionList,
			common.FlagSessionHTTPS}
		if isFlags(c.Command.Flags, testFlagNames, lastArg) {
			return nil
		}

		if isFlag(c.Command.Flags, common.FlagSessionWorkspace, lastArg) {
			workspaces, _ := common.GetWorkspaces(common.GetDir(c))
			return workspaces
		}

		workspace := c.String(common.FlagSessionWorkspace)
		if workspace == "" {
			return nil
		}

		dir := common.GetDir(c)
		workspacePath := filepath.Join(dir, workspace)
		sessions, _ := common.GetSessions(workspacePath)
		return sessions
	})
}
