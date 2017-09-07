package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/metacompiler/ast"
	"github.com/ncbray/lumen/metacompiler/parser"
	"github.com/ncbray/lumen/metacompiler/resolved"
	"github.com/ncbray/lumen/util"
)

type Context struct {
	FileSystem fs.FileSystem
	Logger     log.CompilerLogger
}

type Batch struct {
	SrcDir        string
	DstDir        string
	ParserPackage string
	Files         []string
}

func parseFile(filename string, ctx *Context) *ast.File {
	src := ctx.FileSystem.InputFile(filename)
	b, err := src.GetBytes()
	if err != nil {
		ctx.Logger.Error(err.Error())
		return nil
	}

	input := util.CreateFileStream(filename, string(b))
	lexer := parser.NewRommyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRommyParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	listener := &util.ErrorListener{Filename: filename, Logger: ctx.Logger}
	p.AddErrorListener(listener)
	p.BuildParseTrees = true
	fileTree := p.File()
	if ctx.Logger.NumErrors() > 0 {
		return nil
	}
	conv := &ast.ASTConverter{Filename: filename}
	return conv.ConvertFile(fileTree.(*parser.FileContext))
}

func compile(filename string, parserPackage string, outdir string, ctx *Context) {
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
	resolved.GenerateGo(pkg, rfile, parserPackage, ow)
}

func run() bool {
	logger := log.CreateConsoleLogger()

	tdir, err := fs.MakeTempDir("lumen_")
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

	batches := []*Batch{
		{
			SrcDir:        "src/github.com/ncbray/lumen/metacompiler/_dsl",
			DstDir:        "src/github.com/ncbray/lumen/metacompiler",
			ParserPackage: "github.com/ncbray/lumen/metacompiler/parser",
			Files: []string{
				"ast.dsl",
				"resolved.dsl",
			},
		},
		{
			SrcDir:        "src/github.com/ncbray/lumen/_dsl",
			DstDir:        "src/github.com/ncbray/lumen",
			ParserPackage: "github.com/ncbray/lumen/parser",
			Files: []string{
				"ast.dsl",
				"resolved.dsl",
				"vecgen.dsl",
			},
		},
	}

	for _, b := range batches {
		for _, f := range b.Files {
			dstPkg := f[:strings.Index(f, ".")]
			compile(filepath.Join(b.SrcDir, f), b.ParserPackage, filepath.Join(b.DstDir, dstPkg), ctx)
			if ctx.Logger.NumErrors() > 0 {
				fmt.Fprintf(os.Stderr, "%d errors, stopping\n", logger.NumErrors())
				return false
			}
		}
	}

	fsys.Commit()
	return true
}

func main() {
	if !run() {
		os.Exit(1)
	}
}
