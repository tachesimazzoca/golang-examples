package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"os"
	"strconv"
)

func main() {
	var err error
	var port int
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		assertNoError(err)
	} else {
		port = 1025
	}
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	assertNoError(err)
	for {
		fmt.Println("Listening at ", ln)
		conn, err := ln.Accept()
		assertNoError(err)
		fmt.Println("Accepted: ", conn)
		go handleConnection(conn)
	}
}

func assertNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()
	fmt.Println("Reading lines ...")
	r := textproto.NewReader(bufio.NewReader(conn))
	lines, err := r.ReadDotLines()
	if err != nil {
		return err
	}
	fmt.Println(lines)
	return nil
}
