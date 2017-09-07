package main

import (
	"os"

	"github.com/ncbray/cmdline"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/vecgen"
)

const appName = "vecgen"

type cmdLineOptions struct {
	output string
}

func run(options *cmdLineOptions, fsys fs.BufferedFileSystem, logger log.CompilerLogger) bool {
	out := fsys.OutputFile(options.output, 0644)
	ow, err := out.GetWriter()
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer ow.Close()
	vecgen.GenerateTypeScript(ow)
	return true
}

func wrapper() bool {
	outputFile := &cmdline.FilePath{
		MustExist: false,
	}

	options := &cmdLineOptions{}

	app := cmdline.MakeApp(appName)
	app.RequiredArgs([]*cmdline.Argument{
		{
			Name:  "output",
			Value: outputFile.Set(&options.output),
		},
	})
	app.Run(os.Args[1:])

	logger := log.CreateConsoleLogger()

	tdir, err := fs.MakeTempDir(appName + "_")
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer tdir.Cleanup()
	fsys := fs.MakeBufferedFileSystem(tdir)

	if run(options, fsys, logger) {
		fsys.Commit()
		return true
	}
	return false
}

func main() {
	if !wrapper() {
		os.Exit(1)
	}
}
