package main

import (
	"bufio"
	"io"
)

func ReadLines(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, s.Err()
}
