package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var host string
	var port int
	serverOpt := flag.NewFlagSet("serverOpt", flag.ContinueOnError)
	serverOpt.StringVar(&host, "host", "localhost", "A hostname of serverOpt")
	serverOpt.IntVar(&port, "port", 8080, "A port number of serverOpt")
	//serverOpt.Usage = func() {}
	if err := serverOpt.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}
	fmt.Println("host:", host)
	fmt.Println("port:", port)
	fmt.Println("host via flagSet:", serverOpt.Lookup("host").Value)
	fmt.Println("port via flagSet:", serverOpt.Lookup("port").Value)
	fmt.Println("args:", serverOpt.Args())
}
