package command

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func getURL(req *http.Request) string {
	url := req.Host + req.URL.Path
	if req.URL.RawQuery != "" {
		url += "?" + req.URL.RawQuery
	}
	return url
}

func buildHttpieCmd(req *http.Request, isHTTPS bool, args []string) (string, error) {
	cmd := ""
	if isHTTPS {
		cmd = "https"
	} else {
		cmd = "http"
	}
	cmd += " --ignore-stdin"

	for _, arg := range args {
		cmd += " " + arg
	}

	url := getURL(req)
	cmd += fmt.Sprintf(" '%s'", url)

	for header, values := range req.Header {
		for _, value := range values {
			cmd += fmt.Sprintf(" %s:'%s'", header, value)
		}
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", errors.Wrapf(err, "io.ReadAll")
	}
	if len(body) > 0 {
		cmd += fmt.Sprintf(" --raw='%s'", string(body))
	}

	return cmd, nil
}

func buildCurlCmd(req *http.Request, isHTTPS bool, args []string) (string, error) {
	cmd := "curl"

	for _, arg := range args {
		cmd += " " + arg
	}

	cmd += fmt.Sprintf(" -X %s", req.Method)

	scheme := "http"
	if isHTTPS {
		scheme = "https"
	}
	url := getURL(req)
	url = fmt.Sprintf("%s://%s", scheme, url)

	cmd += fmt.Sprintf(" '%s'", url)

	for header, values := range req.Header {
		for _, value := range values {
			cmd += fmt.Sprintf(" -H '%s: %s'", header, value)
		}
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", errors.Wrapf(err, "io.ReadAll")
	}
	if len(body) > 0 {
		cmd += fmt.Sprintf(" -d '%s'", string(body))
	}

	return cmd, nil
}

func buildHTTPRequest(data []byte) (*http.Request, error) {
	r := bufio.NewReader(bytes.NewReader(data))
	req, err := http.ReadRequest(r)
	if err != nil {
		return nil, errors.Wrapf(err, "http.ReadRequest data:%s", string(data))
	}
	req.Header.Del("Accept-Encoding")
	req.Header.Del("User-Agent")
	return req, nil
}
