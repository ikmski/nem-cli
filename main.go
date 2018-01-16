package main

import "github.com/urfave/cli"

var (
	version  string
	revision string
)

func main() {

	app := cli.NewApp()
	app.Name = "nem-cli"
	app.Usage = "NEM CLI"
	app.Description = "command-line tool for NEM"
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:  "status",
			Usage: "NIS status related commands",
			Subcommands: []cli.Command{
				{
					Name:   "heartbeat",
					Usage:  "show NIS heartbeat",
					Action: cmdHeartbeat,
				},
				{
					Name:   "status",
					Usage:  "show NIS status",
					Action: cmdStatus,
				},
			},
		},
	}

	app.Action = cmdStatus
	app.RunAndExitOnError()
}
