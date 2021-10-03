package command

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/tenfyzhong/hs/common"
	"github.com/urfave/cli/v2"
)

func Transfer(c *cli.Context) error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return errors.Wrapf(err, "io.ReadAll from os.Stdin")
	}

	req, err := buildHTTPRequest(data)
	if err != nil {
		return errors.Wrapf(err, "buildHTTPRequest")
	}

	isHTTPS := c.Bool(common.FlagHTTPS)
	msg := ""
	args := c.Args().Slice()
	if c.Bool(common.FlagCurl) {
		msg, err = buildCurlCmd(req, isHTTPS, args)
		if err != nil {
			return errors.Wrapf(err, "buildCurlCmd data:%s", string(data))
		}
	} else if c.Bool(common.FlagHttpie) {
		msg, err = buildHttpieCmd(req, isHTTPS, args)
		if err != nil {
			return errors.Wrapf(err, "buildHttpieCmd data:%s", string(data))
		}
	}
	fmt.Printf("%s", msg)
	return nil
}
