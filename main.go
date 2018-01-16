package main

import "github.com/urfave/cli"

var (
	version  string
	revision string
)

var nisAddress string = "176.9.68.110"

func main() {

	app := cli.NewApp()
	app.Name = "nem-cli"
	app.Usage = "NEM CLI"
	app.Description = "command-line tool for NEM"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "address, addr",
			Value:       "176.9.68.110",
			Usage:       "address of the NEM super node",
			Destination: &nisAddress,
		},
	}

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
