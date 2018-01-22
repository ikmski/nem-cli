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
		{
			Name:  "account",
			Usage: "NIS account related commands",
			Subcommands: []cli.Command{
				{
					Name:  "data",
					Usage: "show account data",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "public-key",
							Usage: "search from public key",
						},
					},
					Action: cmdAccountData,
				},
			},
		},
		{
			Name:  "chain",
			Usage: "block chain related commands",
			Subcommands: []cli.Command{
				{
					Name:   "height",
					Usage:  "get the current height of the block chain",
					Action: cmdChainHeight,
				},
				{
					Name:  "block",
					Usage: "get a block from the chain that has the given height",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "json",
							Usage: "argument in JSON format",
						},
					},
					Action: cmdChainBlock,
				},
			},
		},
	}

	app.Action = cmdStatus
	app.RunAndExitOnError()
}
