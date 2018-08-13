package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
)

func main() {
	dir, err := homedir.Dir()

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Your home directory is %s.\n", dir)
}
