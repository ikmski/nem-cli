package main

import "github.com/urfave/cli"

func cmdHeartbeat(c *cli.Context) error {

	client, err := newDefaultClient()
	if err != nil {
		return err
	}

	err = client.heartbeat()
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

	err = client.status()
	if err != nil {
		return err
	}

	return nil
}
