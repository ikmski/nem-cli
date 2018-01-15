package main

import (
	"github.com/urfave/cli"
)

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

	app.RunAndExitOnError()
}
