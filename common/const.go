package common

const (
	FlagAppDir = "dir"

	FlagSessionWorkspace = "workspace"
	FlagSessionSave      = "save"
	FlagSessionReplay    = "replay"
	FlagSessionHttpie    = "httpie"
	FlagSessionCurl      = "curl"
	FlagSessionRaw       = "raw"
	FlagSessionHTTPS     = "https"
	FlagSessionList      = "ls"
	FlagSessionRemove    = "rm"
	FlagSessionShowPath  = "show-path"

	FlagWorkspaceCreate   = "create"
	FlagWorkspaceRemove   = "rm"
	FlagWorkspaceList     = "ls"
	FlagWorkspaceShowPath = "show-path"

	SessionSuffix = ".http"
)

const (
	CodeNoDirectoryHs = 255 - iota
	CodeFlagRequired
	CodeNotExist
	CodeUnknownFlag
	CodeUnknown
)
