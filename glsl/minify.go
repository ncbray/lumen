package glsl

import (
	"regexp"
	"strings"
)

var spaces = regexp.MustCompile(`\s+`)

// Characters that are only part of single-character tokens.
// Note: `.` can appear in floating point literals, but close enough.
var distinct = regexp.MustCompile(`\s*([{}()\[\].,;?:~])\s*`)

const cautious = `+*/%<>=!^|&`

var careful1 = regexp.MustCompile(`([` + cautious + `])\s+([^` + cautious + `])`)
var careful2 = regexp.MustCompile(`([^` + cautious + `])\s+([` + cautious + `])`)

func CompactSource(src string) string {
	src = strings.TrimSpace(src)
	src = spaces.ReplaceAllLiteralString(src, " ")
	src = distinct.ReplaceAllString(src, "$1")
	src = careful1.ReplaceAllString(src, "$1$2")
	src = careful2.ReplaceAllString(src, "$1$2")
	return src
}
