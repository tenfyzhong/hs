package common

import (
	"io/ioutil"
	"strings"
)

func GetSessions(workspacePath string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(workspacePath)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		name := fileInfo.Name()
		if !strings.HasSuffix(name, SessionSuffix) {
			continue
		}
		name = name[0 : len(name)-5]
		res = append(res, name)
	}
	return res, nil
}
