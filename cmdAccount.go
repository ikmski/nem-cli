package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdAccountData(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("argument is required")
	}
	arg := c.Args().Get(0)

	client, err := newDefaultClient()
	if err != nil {
		return err
	}

	url := ""
	if c.Bool("public-key") {
		url = fmt.Sprintf("account/get/from-public-key?publicKey=%s", arg)
	} else {
		url = fmt.Sprintf("account/get?address=%s", arg)
	}

	return client.sendRequest("GET", url, nil)
}
