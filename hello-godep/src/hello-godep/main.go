package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

type Hello struct {
	CommandName string
}

func (sc *Hello) Help() string {
	return "This sub-command just prints \"Hello\"."
}

func (sc *Hello) Run(args []string) int {
	fmt.Printf("Hello\n")
	return 0
}

func (sc *Hello) Synopsis() string {
	return sc.CommandName + " hello\n"
}

func main() {
	c := cli.NewCLI("hello-godep", "0.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"hello": func() (cli.Command, error) {
			return &Hello{os.Args[0]}, nil
		},
	}

	exitStatus, _ := c.Run()

	os.Exit(exitStatus)
}
