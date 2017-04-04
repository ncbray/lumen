package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/metacompiler/ast"
	"github.com/ncbray/lumen/metacompiler/parser"
	"github.com/ncbray/lumen/metacompiler/resolved"
)

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

type Context struct {
	FileSystem fs.FileSystem
	Logger     log.CompilerLogger
}

func parseFile(filename string, ctx *Context) *ast.File {
	src := ctx.FileSystem.InputFile(filename)
	b, err := src.GetBytes()
	if err != nil {
		ctx.Logger.Error(err.Error())
		return nil
	}

	input := CreateFileStream(filename, string(b))
	lexer := parser.NewRommyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRommyParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	listener := &ErrorListener{Filename: filename, Logger: ctx.Logger}
	p.AddErrorListener(listener)
	p.BuildParseTrees = true
	fileTree := p.File()
	if ctx.Logger.NumErrors() > 0 {
		return nil
	}
	conv := &ast.ASTConverter{Filename: filename}
	return conv.ConvertFile(fileTree.(*parser.FileContext))
}

func compile(filename string, isAst bool, outdir string, ctx *Context) {
	file := parseFile(filename, ctx)
	if ctx.Logger.NumErrors() > 0 {
		return
	}

	rfile := resolved.FromAST(file, ctx.Logger)
	if ctx.Logger.NumErrors() > 0 {
		return
	}

	_, pkg := filepath.Split(outdir)

	out := ctx.FileSystem.OutputFile(filepath.Join(outdir, "generated.go"), 0644)
	ow, err := out.GetWriter()
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	defer ow.Close()
	resolved.GenerateGo(pkg, rfile, isAst, ow)
}

func run() bool {
	logger := log.CreateConsoleLogger()

	tdir, err := fs.MakeTempDir("mecha_")
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer tdir.Cleanup()

	fsys := fs.MakeBufferedFileSystem(tdir)
	ctx := &Context{
		FileSystem: fsys,
		Logger:     logger,
	}

	pkgPath := "src/github.com/ncbray/lumen"
	metacompilerPath := filepath.Join(pkgPath, "metacompiler")

	compile(filepath.Join(metacompilerPath, "_dsl/ast.dsl"), true, filepath.Join(metacompilerPath, "ast"), ctx)
	if ctx.Logger.NumErrors() > 0 {
		fmt.Fprintf(os.Stderr, "%d errors, stopping\n", logger.NumErrors())
		return false
	}

	compile(filepath.Join(metacompilerPath, "_dsl/resolved.dsl"), false, filepath.Join(metacompilerPath, "resolved"), ctx)
	if ctx.Logger.NumErrors() > 0 {
		fmt.Fprintf(os.Stderr, "%d errors, stopping\n", logger.NumErrors())
		return false
	}

	fsys.Commit()
	return true
}

func main() {
	if !run() {
		os.Exit(1)
	}
}
