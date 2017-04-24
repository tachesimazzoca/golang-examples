package utils

import (
	"strings"
	"unicode"
)

func UcFirst(s string) string {
	xs := []rune(strings.ToLower(s))
	ys := make([]rune, len(xs))

	for i, x := range xs {
		if i == 0 {
			ys[i] = unicode.ToUpper(x)
		} else {
			ys[i] = x
		}
	}

	return string(ys)
}

func Camelize(s string) string {
	xs := strings.Split(s, "_")
	n := len(xs)
	ys := make([]string, n)

	for i, x := range xs {
		ys[i] = UcFirst(x)
	}

	return strings.Join(ys, "")
}
