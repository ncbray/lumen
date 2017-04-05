package util

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ncbray/lumen/log"
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

type FileStream struct {
	*antlr.InputStream
	filename string
}

func (f *FileStream) GetSourceName() string {
	return f.filename
}

func CreateFileStream(filename string, data string) *FileStream {
	return &FileStream{
		InputStream: antlr.NewInputStream(data),
		filename:    filename,
	}
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Filename string
	Logger   log.CompilerLogger
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	l.Logger.ErrorAtLocation(l.Filename, line, column, msg)
}
