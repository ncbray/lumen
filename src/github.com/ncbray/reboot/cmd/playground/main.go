package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ncbray/cmdline"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/reboot/ast"
	"github.com/ncbray/reboot/log"
	"github.com/ncbray/reboot/parser"
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

func compile(filename string, isAst bool, ctx *Context) {
	file := parseFile(filename, ctx)
	if ctx.Logger.NumErrors() > 0 {
		return
	}
	dir, _ := filepath.Split(filename)
	_, pkg := filepath.Split(dir[:len(dir)-1])
	//ext := filepath.Ext(leaf)
	//leaf = leaf[:len(leaf)-len(ext)] + ".go"

	out := ctx.FileSystem.OutputFile(dir+"generated.go", 0644)
	ow, err := out.GetWriter()
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	defer ow.Close()
	ast.GenerateGo(pkg, file, isAst, ow)
}

func run(input string, isAst bool) bool {
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

	compile(input, isAst, ctx)

	if ctx.Logger.NumErrors() > 0 {
		fmt.Fprintf(os.Stderr, "%d errors, stopping\n", logger.NumErrors())
		return false
	}

	fsys.Commit()
	return true
}

func main() {
	inputFile := &cmdline.FilePath{
		MustExist: true,
	}

	var isAst bool
	var input string

	app := cmdline.MakeApp("playground")
	app.Flags([]*cmdline.Flag{
		{
			Long: "ast",
			Call: cmdline.SetTrue(&isAst),
		},
	})
	app.RequiredArgs([]*cmdline.Argument{
		{
			Name:  "input",
			Value: inputFile.Set(&input),
		},
	})
	app.Run(os.Args[1:])

	if !run(input, isAst) {
		os.Exit(1)
	}
}
