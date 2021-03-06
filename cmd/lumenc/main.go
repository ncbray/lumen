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
	var minify bool
	var input string
	var typeScript bool
	var haxeTarget string
	var haxePkg string
	var output string

	inputFile := &cmdline.FilePath{
		MustExist: true,
	}

	targetEnum := &cmdline.Enum{Possible: []string{"android", "html5", "ios", "linux", "mac", "windows"}}

	outputFile := &cmdline.FilePath{
		MustExist: false,
	}

	app := cmdline.MakeApp("lumenc")
	app.Flags([]*cmdline.Flag{
		{
			Long: "minify",
			Call: cmdline.SetTrue(&minify),
		},
		{
			Long: "ts",
			Call: cmdline.SetTrue(&typeScript),
		},
		{
			Long:  "haxe-target",
			Value: targetEnum.Set(&haxeTarget),
		},
		{
			Long:  "haxe-package",
			Value: cmdline.String.Set(&haxePkg),
		},
	})
	app.RequiredArgs([]*cmdline.Argument{
		{
			Name:  "input",
			Value: inputFile.Set(&input),
		},
		{
			Name:  "output",
			Value: outputFile.Set(&output),
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

	out := fsys.OutputFile(output, 0644)
	ow, err := out.GetWriter()
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer ow.Close()

	if typeScript {
		resolved.GenerateTypeScript(rfile, minify, ow)
	}

	// TODO better error checking.
	if haxeTarget != "" {
		resolved.GenerateHaxe(haxePkg, rfile, haxeTarget == "android" || haxeTarget == "html5" || haxeTarget == "ios", minify, ow)
	}

	fsys.Commit()
	return true
}

func main() {
	if !run() {
		os.Exit(1)
	}
}
