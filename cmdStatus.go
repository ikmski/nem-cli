package main

import "github.com/urfave/cli"

func cmdHeartbeat(c *cli.Context) error {

	client, err := newDefaultClient()
	if err != nil {
		return err
	}

	err = client.sendRequest("GET", "heartbeat", nil)
	if err != nil {
		return err
	}

	return nil
}

func cmdStatus(c *cli.Context) error {

	client, err := newDefaultClient()
	if err != nil {
		return err
	}

	err = client.sendRequest("GET", "status", nil)
	if err != nil {
		return err
	}

	return nil
}
