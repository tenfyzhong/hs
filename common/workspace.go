package common

import (
	"io/ioutil"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func GetWorkspaces(dir string) ([]string, error) {
	fileInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return []string{}, nil
	}
	if !fileInfo.IsDir() {
		return nil, cli.Exit(dir+" is not a directory", CodeNoDirectoryHs)
	}
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, cli.Exit(err.Error(), CodeNoDirectoryHs)
	}
	res := make([]string, 0, len(fileInfos))
	for _, info := range fileInfos {
		if !info.IsDir() {
			continue
		}
		res = append(res, info.Name())
	}
	sort.Strings(res)
	return res, nil
}
