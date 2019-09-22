package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLines(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, s.Err()
}

func WriteLines(w io.Writer, lines []string) error {
	bw := bufio.NewWriter(w)
	for _, line := range lines {
		fmt.Fprintln(bw, line)
	}
	return bw.Flush()
}

func ReadLinesFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadLines(f)
}

func WriteLinesToFile(path string, lines []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteLines(f, lines)
}
