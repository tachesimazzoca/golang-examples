package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func testReadLines(t *testing.T, xs []string, sep string) {
	lines, _ := ReadLines(strings.NewReader(strings.Join(xs[:], sep) + sep))
	if len(lines) != len(xs) {
		t.Errorf("The number of lines: expeced %d, actual %d", len(xs), len(lines))
	}
	for i, x := range xs {
		if x != lines[i] {
			t.Errorf("The number of lines: expeced %s, actual %s", x, lines[i])
		}
	}
}

func TestReadLinesWithUnixNewLine(t *testing.T) {
	sep := "\n"
	testReadLines(t, []string{"foo", "bar", "", "baz", "", ""}, sep)
}

func TestReadLinesWithWindowsNewLine(t *testing.T) {
	sep := "\r\n"
	testReadLines(t, []string{"foo", "bar", "", "baz", "", ""}, sep)
}

func TestReadLinesFromFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "utils_test_")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tf.Name())

	expectedItems := []string{"hello", "ioutil", "TempFile."}
	content := strings.Join(expectedItems, "\n")

	if _, err := tf.WriteString(content); err != nil {
		t.Error(err)
	}

	if _, err := tf.Seek(0, os.SEEK_SET); err != nil {
		t.Error(err)
	}

	lines, err := ReadLines(tf)
	if err != nil {
		t.Error(err)
	}
	if len(lines) != len(expectedItems) {
		t.Errorf("The number of lines: expeced %d, actual %d", len(expectedItems), len(lines))
	}
	actualContent := strings.Join(lines, "\n")
	if content != actualContent {
		t.Errorf("The returned lines: expeced %s, actual %s", content, actualContent)
	}
}
