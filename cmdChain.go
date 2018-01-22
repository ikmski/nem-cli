package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func cmdChainHeight(c *cli.Context) error {

	client, err := newDefaultClient()
	if err != nil {
		return err
	}

	return client.sendRequest("GET", "chain/height", nil)
}

func cmdChainBlock(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("argument is required")
	}

	client, err := newDefaultClient()
	if err != nil {
		return err
	}

	json := ""

	if c.Bool("json") {
		for _, str := range c.Args() {
			json += str
		}
	} else {
		json = fmt.Sprintf("{\"height\": %s}", c.Args().Get(0))
	}

	return client.sendRequest("POST", "block/at/public", strings.NewReader(json))
}
