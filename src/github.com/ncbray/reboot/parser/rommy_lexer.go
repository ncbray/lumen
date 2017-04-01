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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 13, 99, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 7, 9, 53, 10, 9, 12, 9, 14, 9, 56, 11, 9, 3, 10, 6, 10,
	59, 10, 10, 13, 10, 14, 10, 60, 3, 10, 3, 10, 6, 10, 65, 10, 10, 13, 10,
	14, 10, 66, 5, 10, 69, 10, 10, 3, 10, 3, 10, 5, 10, 73, 10, 10, 3, 10,
	6, 10, 76, 10, 10, 13, 10, 14, 10, 77, 5, 10, 80, 10, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 7, 11, 86, 10, 11, 12, 11, 14, 11, 89, 11, 11, 3, 11, 3,
	11, 3, 12, 6, 12, 94, 10, 12, 13, 12, 14, 12, 95, 3, 12, 3, 12, 2, 2, 13,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 3, 2, 9, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2,
	50, 59, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 4, 2, 12, 12, 15,
	15, 5, 2, 11, 12, 15, 15, 34, 34, 107, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 28, 3, 2, 2, 2, 7, 32,
	3, 2, 2, 2, 9, 34, 3, 2, 2, 2, 11, 36, 3, 2, 2, 2, 13, 38, 3, 2, 2, 2,
	15, 43, 3, 2, 2, 2, 17, 50, 3, 2, 2, 2, 19, 58, 3, 2, 2, 2, 21, 81, 3,
	2, 2, 2, 23, 93, 3, 2, 2, 2, 25, 26, 7, 93, 2, 2, 26, 27, 7, 95, 2, 2,
	27, 4, 3, 2, 2, 2, 28, 29, 7, 120, 2, 2, 29, 30, 7, 99, 2, 2, 30, 31, 7,
	116, 2, 2, 31, 6, 3, 2, 2, 2, 32, 33, 7, 61, 2, 2, 33, 8, 3, 2, 2, 2, 34,
	35, 7, 125, 2, 2, 35, 10, 3, 2, 2, 2, 36, 37, 7, 127, 2, 2, 37, 12, 3,
	2, 2, 2, 38, 39, 7, 103, 2, 2, 39, 40, 7, 112, 2, 2, 40, 41, 7, 119, 2,
	2, 41, 42, 7, 111, 2, 2, 42, 14, 3, 2, 2, 2, 43, 44, 7, 117, 2, 2, 44,
	45, 7, 118, 2, 2, 45, 46, 7, 116, 2, 2, 46, 47, 7, 119, 2, 2, 47, 48, 7,
	101, 2, 2, 48, 49, 7, 118, 2, 2, 49, 16, 3, 2, 2, 2, 50, 54, 9, 2, 2, 2,
	51, 53, 9, 3, 2, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3,
	2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 18, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57,
	59, 9, 4, 2, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2,
	2, 60, 61, 3, 2, 2, 2, 61, 68, 3, 2, 2, 2, 62, 64, 7, 48, 2, 2, 63, 65,
	9, 4, 2, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 62, 3, 2, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 79, 3, 2, 2, 2, 70, 72, 9, 5, 2, 2, 71, 73, 9, 6, 2, 2, 72,
	71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 76, 9, 4, 2,
	2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78,
	3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 70, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2,
	80, 20, 3, 2, 2, 2, 81, 82, 7, 49, 2, 2, 82, 83, 7, 49, 2, 2, 83, 87, 3,
	2, 2, 2, 84, 86, 10, 7, 2, 2, 85, 84, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87,
	85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3, 2, 2,
	2, 90, 91, 8, 11, 2, 2, 91, 22, 3, 2, 2, 2, 92, 94, 9, 8, 2, 2, 93, 92,
	3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2,
	96, 97, 3, 2, 2, 2, 97, 98, 8, 12, 2, 2, 98, 24, 3, 2, 2, 2, 12, 2, 54,
	60, 66, 68, 72, 77, 79, 87, 95, 3, 2, 3, 2,
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
	"", "", "", "", "", "", "", "", "ID", "NUM", "SINGLE_LINE_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ID", "NUM", "SINGLE_LINE_COMMENT",
	"WS",
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
	RommyLexerT__0                = 1
	RommyLexerT__1                = 2
	RommyLexerT__2                = 3
	RommyLexerT__3                = 4
	RommyLexerT__4                = 5
	RommyLexerT__5                = 6
	RommyLexerT__6                = 7
	RommyLexerID                  = 8
	RommyLexerNUM                 = 9
	RommyLexerSINGLE_LINE_COMMENT = 10
	RommyLexerWS                  = 11
)
