package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

var LineBreak = fmt.Sprintln()

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

func TestReadLines(t *testing.T) {
	tf, err := ioutil.TempFile("", "utils_test_")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tf.Name())

	content := `now
testing
ReadLines
`
	if err := ioutil.WriteFile(tf.Name(), []byte(content), 0644); err != nil {
		t.Error(err)
	}

	lines, err := ReadLines(tf)
	if err != nil {
		t.Error(err)
	}

	expectedLines := []string{"now", "testing", "ReadLines"}
	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("lines expected: %s, actual: %s", expectedLines, lines)
	}
}

func TestWriteLines(t *testing.T) {
	tf, err := ioutil.TempFile("", "utils_test_")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tf.Name())

	lines := []string{"now", "testing", "WriteLines"}
	if err := WriteLines(tf, lines); err != nil {
		t.Error(err)
	}

	bs, err := ioutil.ReadFile(tf.Name())
	if err != nil {
		t.Error(err)
	}

	expectedContent := `now
testing
WriteLines
`
	if string(bs) != expectedContent {
		t.Error(bs)
	}
}

func TestReadLinesFromFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "utils_test_")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tf.Name())

	content := `now
testing
ReadLinesFrom
`
	if err := ioutil.WriteFile(tf.Name(), []byte(content), 0644); err != nil {
		t.Error(err)
	}

	lines, err := ReadLinesFromFile(tf.Name())
	if err != nil {
		t.Error(err)
	}

	expectedLines := []string{"now", "testing", "ReadLinesFrom"}
	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("lines expected: %s, actual: %s", expectedLines, lines)
	}
}

func TestWriteLinesToFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "utils_test_")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tf.Name())

	lines := []string{"now", "testing", "WriteLinesTo"}
	if err := WriteLinesToFile(tf.Name(), lines); err != nil {
		t.Error(err)
	}

	bs, err := ioutil.ReadFile(tf.Name())
	if err != nil {
		t.Error(err)
	}

	expectedContent := `now
testing
WriteLinesTo
`
	if string(bs) != expectedContent {
		t.Error(bs)
	}
}
