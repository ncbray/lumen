package util

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Location struct {
	File   string
	Line   int
	Column int
}

func GetLocation(file string, tok antlr.Token) Location {
	return Location{
		File:   file,
		Line:   tok.GetLine(),
		Column: tok.GetColumn(),
	}
}
