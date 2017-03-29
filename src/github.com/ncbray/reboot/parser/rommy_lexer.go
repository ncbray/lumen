// Generated from Rommy.g by ANTLR 4.6.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 12, 83, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	7, 9, 51, 10, 9, 12, 9, 14, 9, 54, 11, 9, 3, 10, 6, 10, 57, 10, 10, 13,
	10, 14, 10, 58, 3, 10, 3, 10, 6, 10, 63, 10, 10, 13, 10, 14, 10, 64, 5,
	10, 67, 10, 10, 3, 10, 3, 10, 5, 10, 71, 10, 10, 3, 10, 6, 10, 74, 10,
	10, 13, 10, 14, 10, 75, 5, 10, 78, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 3, 2, 8, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124,
	3, 2, 50, 59, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 11, 12,
	15, 15, 34, 34, 89, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3,
	2, 2, 2, 5, 26, 3, 2, 2, 2, 7, 30, 3, 2, 2, 2, 9, 32, 3, 2, 2, 2, 11, 34,
	3, 2, 2, 2, 13, 36, 3, 2, 2, 2, 15, 41, 3, 2, 2, 2, 17, 48, 3, 2, 2, 2,
	19, 56, 3, 2, 2, 2, 21, 79, 3, 2, 2, 2, 23, 24, 7, 93, 2, 2, 24, 25, 7,
	95, 2, 2, 25, 4, 3, 2, 2, 2, 26, 27, 7, 120, 2, 2, 27, 28, 7, 99, 2, 2,
	28, 29, 7, 116, 2, 2, 29, 6, 3, 2, 2, 2, 30, 31, 7, 61, 2, 2, 31, 8, 3,
	2, 2, 2, 32, 33, 7, 125, 2, 2, 33, 10, 3, 2, 2, 2, 34, 35, 7, 127, 2, 2,
	35, 12, 3, 2, 2, 2, 36, 37, 7, 103, 2, 2, 37, 38, 7, 112, 2, 2, 38, 39,
	7, 119, 2, 2, 39, 40, 7, 111, 2, 2, 40, 14, 3, 2, 2, 2, 41, 42, 7, 117,
	2, 2, 42, 43, 7, 118, 2, 2, 43, 44, 7, 116, 2, 2, 44, 45, 7, 119, 2, 2,
	45, 46, 7, 101, 2, 2, 46, 47, 7, 118, 2, 2, 47, 16, 3, 2, 2, 2, 48, 52,
	9, 2, 2, 2, 49, 51, 9, 3, 2, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2,
	52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 18, 3, 2, 2, 2, 54, 52, 3,
	2, 2, 2, 55, 57, 9, 4, 2, 2, 56, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58,
	56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 66, 3, 2, 2, 2, 60, 62, 7, 48,
	2, 2, 61, 63, 9, 4, 2, 2, 62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 60, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 77, 3, 2, 2, 2, 68, 70, 9, 5, 2, 2, 69, 71, 9,
	6, 2, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72,
	74, 9, 4, 2, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2,
	2, 75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 68, 3, 2, 2, 2, 77, 78,
	3, 2, 2, 2, 78, 20, 3, 2, 2, 2, 79, 80, 9, 7, 2, 2, 80, 81, 3, 2, 2, 2,
	81, 82, 8, 11, 2, 2, 82, 22, 3, 2, 2, 2, 10, 2, 52, 58, 64, 66, 70, 75,
	77, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'[]'", "'var'", "';'", "'{'", "'}'", "'enum'", "'struct'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "ID", "NUM", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ID", "NUM", "WS",
}

type RommyLexer struct {
	*antlr.BaseLexer
	modeNames []string
	// TODO: EOF string
}

func NewRommyLexer(input antlr.CharStream) *RommyLexer {
	var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	l := new(RommyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Rommy.g"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RommyLexer tokens.
const (
	RommyLexerT__0 = 1
	RommyLexerT__1 = 2
	RommyLexerT__2 = 3
	RommyLexerT__3 = 4
	RommyLexerT__4 = 5
	RommyLexerT__5 = 6
	RommyLexerT__6 = 7
	RommyLexerID   = 8
	RommyLexerNUM  = 9
	RommyLexerWS   = 10
)
