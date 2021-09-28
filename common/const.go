package common

const (
	FlagAppDir    = "dir"
	FlagHTTPS     = "https"
	FlagWorkspace = "workspace"
	FlagReplay    = "replay"
	FlagHttpie    = "httpie"
	FlagCurl      = "curl"
	FlagRaw       = "raw"
	FlagList      = "ls"
	FlagRemove    = "rm"
	FlagShowPath  = "show-path"
	FlagCreate    = "create"

	SessionSuffix = ".http"
)

const (
	CodeNoDirectoryHs = 255 - iota
	CodeFlagRequired
	CodeNotExist
	CodeUnknownFlag
	CodeUnknown
)
