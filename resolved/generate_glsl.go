package resolved

import (
	"io"
	"strings"

	"github.com/ncbray/compilerutil/writer"
)

func typeName(t Type) string {
	switch t := t.(type) {
	case *IntrinsicType:
		return t.Name
	default:
		panic(t)
	}
}

func generateFormat(name string, format *Format, o *writer.TabbedWriter) {
	for _, f := range format.Fields {
		if !f.Marked {
			continue
		}
		o.WriteLine(name + " " + typeName(f.Type) + " " + f.Name + ";")
	}
}

func generateArgs(args []Expr, o *writer.TabbedWriter) {
	o.WriteString("(")
	for i, arg := range args {
		if i != 0 {
			o.WriteString(", ")
		}
		generateExpr(arg, o)
	}
	o.WriteString(")")
}

func generateExpr(e Expr, o *writer.TabbedWriter) {
	switch e := e.(type) {
	case *Number:
		o.WriteString(e.Raw)
	case *GetInput:
		o.WriteString(e.Input.Name)
	case *GetLocal:
		o.WriteString(e.Local.Name)
	case *Infix:
		generateExpr(e.Left, o)
		o.WriteString(" ")
		o.WriteString(e.Op)
		o.WriteString(" ")
		generateExpr(e.Right, o)
	case *GetAttr:
		generateExpr(e.Value, o)
		o.WriteString(".")
		o.WriteString(e.Name)
	case *Constructor:
		o.WriteString(typeName(e.Type))
		generateArgs(e.Args, o)
	case *CallIntrinsic:
		o.WriteString(e.Function.Name)
		generateArgs(e.Args, o)
	default:
		panic(e)
	}
}

func generateBody(f *Function, o *writer.TabbedWriter) {
	o.WriteLine("void main() {")
	o.Indent()
	for _, lcl := range f.Locals {
		o.WriteLine(typeName(lcl.Type) + " " + lcl.Name + ";")
	}
	for _, s := range f.Body {
		switch s := s.(type) {
		case *SetLocal:
			o.WriteString(s.Local.Name)
			o.WriteString(" = ")
			generateExpr(s.Value, o)
			o.WriteString(";")
			o.EndOfLine()
		case *SetOutput:
			o.WriteString(s.Output.Name)
			o.WriteString(" = ")
			generateExpr(s.Value, o)
			o.WriteString(";")
			o.EndOfLine()
		default:
			panic(s)
		}
	}
	o.Dedent()
	o.WriteLine("}")
}

func clearMarks(format *Format) {
	for _, f := range format.Fields {
		f.Marked = false
	}
}

func findUsedExpr(e Expr) {
	switch e := e.(type) {
	case *Number, *GetLocal:
	case *GetInput:
		e.Input.Marked = true
	case *GetAttr:
		findUsedExpr(e.Value)
	case *Infix:
		findUsedExpr(e.Left)
		findUsedExpr(e.Right)
	case *Constructor:
		for _, arg := range e.Args {
			findUsedExpr(arg)
		}
	case *CallIntrinsic:
		for _, arg := range e.Args {
			findUsedExpr(arg)
		}
	default:
		panic(e)
	}
}

func findUsed(s *ShaderProgram, f *Function) {
	clearMarks(s.Uniform)
	clearMarks(s.Attribute)
	clearMarks(s.Varying)

	for _, s := range f.Body {
		switch s := s.(type) {
		case *SetOutput:
			s.Output.Marked = true
			findUsedExpr(s.Value)
		case *SetLocal:
			findUsedExpr(s.Value)
		case *Discard:
			findUsedExpr(s.Value)
		default:
			panic(s)
		}
	}
}

func GenerateGLSLVertShader(s *ShaderProgram, declPrecision bool, out io.Writer) {
	findUsed(s, s.Vs)

	o := writer.MakeTabbedWriter("  ", out)
	if declPrecision {
		o.WriteLine("precision mediump float;")
	}
	generateFormat("uniform", s.Uniform, o)
	generateFormat("attribute", s.Attribute, o)
	generateFormat("varying", s.Varying, o)
	generateBody(s.Vs, o)
}

func GenerateGLSLFragShader(s *ShaderProgram, declPrecision bool, out io.Writer) {
	findUsed(s, s.Fs)

	o := writer.MakeTabbedWriter("  ", out)
	if declPrecision {
		o.WriteLine("precision mediump float;")
	}
	generateFormat("uniform", s.Uniform, o)
	generateFormat("varying", s.Varying, o)
	generateBody(s.Fs, o)
}

func isSampler(t Type) bool {
	switch t := t.(type) {
	case *IntrinsicType:
		return strings.HasPrefix(t.Name, "sampler")
	default:
		panic(t)
	}
}
