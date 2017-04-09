package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ncbray/cmdline"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/lumen/ast"
	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/parser"
	"github.com/ncbray/lumen/resolved"
	"github.com/ncbray/lumen/util"
)

func parseFile(filename string, fsys fs.FileSystem, logger log.CompilerLogger) *ast.File {
	src := fsys.InputFile(filename)
	b, err := src.GetBytes()
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	input := util.CreateFileStream(filename, string(b))
	lexer := parser.NewLumenLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLumenParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	listener := &util.ErrorListener{Filename: filename, Logger: logger}
	p.AddErrorListener(listener)
	p.BuildParseTrees = true
	fileTree := p.File()
	if logger.NumErrors() > 0 {
		return nil
	}
	conv := &ast.ASTConverter{Filename: filename}
	return conv.ConvertFile(fileTree.(*parser.FileContext))
}

func run() bool {
	var pkg string
	var input string

	inputFile := &cmdline.FilePath{
		MustExist: true,
	}

	app := cmdline.MakeApp("lumenc")
	app.RequiredArgs([]*cmdline.Argument{
		{
			Name:  "package",
			Value: cmdline.String.Set(&pkg),
		},
		{
			Name:  "input",
			Value: inputFile.Set(&input),
		},
	})
	app.Run(os.Args[1:])

	logger := log.CreateConsoleLogger()

	tdir, err := fs.MakeTempDir("lumen_")
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer tdir.Cleanup()
	fsys := fs.MakeBufferedFileSystem(tdir)

	file := parseFile(input, fsys, logger)
	if logger.NumErrors() > 0 {
		return false
	}
	rfile := resolved.FromAST(file, logger)
	if logger.NumErrors() > 0 {
		return false
	}

	resolved.GenerateHaxe(pkg, rfile, true, os.Stdout)

	fsys.Commit()
	return true
}

func main() {
	if !run() {
		os.Exit(1)
	}
}
