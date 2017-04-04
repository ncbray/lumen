package log

import (
	"fmt"
	"io"

	"github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
)

type CompilerLogger interface {
	Error(message string)
	ErrorAtLocation(filename string, line int, column int, message string)
	NumErrors() int
}

type consoleLogger struct {
	stdout    io.Writer
	stderr    io.Writer
	numErrors int
}

var _ CompilerLogger = (*consoleLogger)(nil)

func (l *consoleLogger) Error(message string) {
	l.stderr.Write([]byte(ansi.Red))
	fmt.Fprintf(l.stderr, "ERROR %s\n", message)
	l.stderr.Write([]byte(ansi.Reset))
	l.numErrors += 1

}

func (l *consoleLogger) ErrorAtLocation(filename string, line int, column int, message string) {
	l.Error(fmt.Sprintf("%s:%d:%d: %s", filename, line, column, message))
}

func (l *consoleLogger) NumErrors() int {
	return l.numErrors
}

func CreateConsoleLogger() CompilerLogger {
	return &consoleLogger{
		stdout: colorable.NewColorableStdout(),
		stderr: colorable.NewColorableStderr(),
	}
}
