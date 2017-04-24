package utils

import (
	"testing"
)

func TestUcFirst(t *testing.T) {
	xs := []struct {
		in, want string
	}{
		{"foo", "Foo"},
		{"123", "123"},
		{"aBcD", "Abcd"},
		{"", ""},
	}
	for _, x := range xs {
		y := UcFirst(x.in)
		if y != x.want {
			t.Errorf("UcFirst(%q) == %q, want %q", x.in, y, x.want)
		}
	}
}

func TestCamelize(t *testing.T) {
	xs := []struct {
		in, want string
	}{
		{"string_utils", "StringUtils"},
		{"multibyte_日本語_string", "Multibyte日本語String"},
	}
	for _, x := range xs {
		y := Camelize(x.in)
		if y != x.want {
			t.Errorf("Camelize(%q) == %q, want %q", x.in, y, x.want)
		}
	}
}
