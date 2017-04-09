package resolved

import (
	"bytes"
	"fmt"
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
		o.WriteString(e.Name)
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

func GenerateGLSLVertShader(s *ShaderProgram, out io.Writer) {
	o := writer.MakeTabbedWriter("  ", out)
	generateFormat("uniform", s.Uniform, o)
	generateFormat("attribute", s.Attribute, o)
	generateFormat("varying", s.Varying, o)
	generateBody(s.Vs, o)
}

func GenerateGLSLFragShader(s *ShaderProgram, declPrecision bool, out io.Writer) {
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

func GenerateHaxe(pkg string, file *File, declPrecision bool, out io.Writer) {
	o := writer.MakeTabbedWriter("  ", out)

	o.WriteLine(pkg + ";")

	imports := []string{
		"lime.graphics.opengl.GL",
		"lime.graphics.opengl.GLProgram",
		"lime.graphics.opengl.GLUniformLocation",
	}

	o.EndOfLine()
	for _, imp := range imports {
		o.WriteLine("import " + imp + ";")
	}

	for _, p := range file.Programs {
		var vs bytes.Buffer
		GenerateGLSLVertShader(p, &vs)
		var fs bytes.Buffer
		GenerateGLSLFragShader(p, true, &fs)

		o.EndOfLine()
		o.WriteLine("class " + p.Name + "ShaderProgram {")
		o.Indent()

		uniformData := []*Field{}
		uniformSamplers := []*Field{}
		for _, f := range p.Uniform.Fields {
			if isSampler(f.Type) {
				uniformSamplers = append(uniformSamplers, f)
			} else {
				uniformData = append(uniformData, f)
			}
		}

		o.WriteLine(fmt.Sprintf("public static inline var VERTEX_SHADER_SOURCE = %q;", vs.String()))
		o.WriteLine(fmt.Sprintf("public static inline var FRAGMENT_SHADER_SOURCE = %q;", fs.String()))

		o.WriteString("public static var ATTRIBUTE_NAMES = [")
		for i, a := range p.Attribute.Fields {
			if i != 0 {
				o.WriteString(", ")
			}
			o.WriteString("\"" + a.Name + "\"")
		}
		o.WriteString("];")
		o.EndOfLine()

		o.WriteString("public static var SAMPLER_NAMES = [")
		for i, u := range uniformSamplers {
			if i != 0 {
				o.WriteString(", ")
			}
			o.WriteString("\"" + u.Name + "\"")
		}
		o.WriteString("];")
		o.EndOfLine()

		// Fields
		o.EndOfLine()
		o.WriteLine("public var program:GLProgram;")
		for _, u := range uniformData {
			o.WriteLine("public var " + u.Name + ":GLUniformLocation;")
		}

		// Constructor
		o.EndOfLine()
		o.WriteLine("public function new(p:GLProgram) {")
		o.Indent()
		o.WriteLine("program = p;")
		for _, u := range uniformData {
			o.WriteLine(u.Name + " = GL.getUniformLocation(program, \"" + u.Name + "\");")
		}
		o.Dedent()
		o.WriteLine("}")

		o.Dedent()
		o.WriteLine("}")

	}

	o.EndOfLine()
	o.WriteLine("class ShaderPrograms {")
	o.Indent()
	o.Dedent()
	o.WriteLine("}")
}
