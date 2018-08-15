package command

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

type HexDumpCommand struct {
	Meta
}

func (c *HexDumpCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		return 1
	}
	f, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer f.Close()

	d := hex.Dumper(os.Stdout)
	defer d.Close()
	for {
		buf := make([]byte, 4096)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		if _, e := d.Write(buf[:n]); e != nil {
			fmt.Fprintln(os.Stderr, e)
			return 1
		}
	}
	return 0
}

func (c *HexDumpCommand) Help() string {
	helpText := `
Usage: %s hex dump <path/to/file>
`
	return strings.TrimSpace(fmt.Sprintf(helpText, c.AppName))
}

func (c *HexDumpCommand) Synopsis() string {
	return "dump a file in hex and ASCII"
}
