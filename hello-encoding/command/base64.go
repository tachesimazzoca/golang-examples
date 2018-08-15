package command

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Base64EncodeCommand struct {
	Meta
}

func (c *Base64EncodeCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println(c.Help())
		return 1
	}
	data := []byte(args[0])
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	return 0
}

func (c *Base64EncodeCommand) Help() string {
	helpText := `
Usage: %s base64 encode <string>
`
	return strings.TrimSpace(fmt.Sprintf(helpText, c.AppName))
}

func (c *Base64EncodeCommand) Synopsis() string {
	return "encode strings into base64 characters"
}

type Base64DecodeCommand struct {
	Meta
}

func (c *Base64DecodeCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println(c.Help())
		return 1
	}
	data, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Printf("%q\n", data)
	return 0
}

func (c *Base64DecodeCommand) Help() string {
	helpText := `
Usage: %s base64 decode <string>
`
	return strings.TrimSpace(fmt.Sprintf(helpText, c.AppName))
}

func (c *Base64DecodeCommand) Synopsis() string {
	return "decode base64 characters"
}
