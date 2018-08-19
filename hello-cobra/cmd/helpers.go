package cmd

import "log"

func debug(msg string) {
	if verbose {
		log.Println(msg)
	}
}
