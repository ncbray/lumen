package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ncbray/reboot/ast"
	"github.com/ncbray/reboot/parser"
)

func main() {
	filename := os.Args[1]
	input := antlr.NewFileStream(filename)
	lexer := parser.NewRommyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRommyParser(stream)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	fileTree := p.File()
	conv := &ast.ASTConverter{Filename: filename}
	file := conv.ConvertFile(fileTree.(*parser.FileContext))
	ast.GenerateGo(file, os.Stdout)
}
