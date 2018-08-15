package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/tachesimazzoca/golang-examples/hello-encoding/command"
)

const AppName = "hello-encoding"

func main() {
	commandMeta := &command.Meta{
		AppName: AppName,
	}

	cmd := cli.NewCLI(commandMeta.AppName, "0.0.0")
	cmd.Args = os.Args[1:]
	cmd.Commands = map[string]cli.CommandFactory{
		"base64 encode": func() (cli.Command, error) {
			return &command.Base64EncodeCommand{
				Meta: *commandMeta,
			}, nil
		},
		"base64 decode": func() (cli.Command, error) {
			return &command.Base64DecodeCommand{
				Meta: *commandMeta,
			}, nil
		},
		"hex dump": func() (cli.Command, error) {
			return &command.HexDumpCommand{
				Meta: *commandMeta,
			}, nil
		},
	}

	status, err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(status)
}
