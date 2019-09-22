package hello

import (
	"fmt"
	"os"

	"github.com/mitchellh/colorstring"
)

func Success(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, colorstring.Color("[green]OK ")+format, args...)
}

func Failure(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, colorstring.Color("[red]NG ")+format, args...)
}
