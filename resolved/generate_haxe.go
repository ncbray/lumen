package resolved

import (
	"bytes"
	"fmt"
	"io"

	"github.com/ncbray/compilerutil/names"
	"github.com/ncbray/compilerutil/writer"
)

func GenerateHaxe(pkg string, file *File, declPrecision bool, out io.Writer) {
	o := writer.MakeTabbedWriter("  ", out)

	o.WriteLine("package " + pkg + ";")

	imports := []string{
		"lime.graphics.opengl.GL",
		"lime.graphics.opengl.GLProgram",
		"lime.graphics.opengl.GLUniformLocation",
		"mecha.graphics.state.RenderPipeline",
	}

	o.EndOfLine()
	for _, imp := range imports {
		o.WriteLine("import " + imp + ";")
	}

	for _, p := range file.Programs {
		var vs bytes.Buffer
		GenerateGLSLVertShader(p, declPrecision, &vs)
		var fs bytes.Buffer
		GenerateGLSLFragShader(p, declPrecision, &fs)

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
	for _, p := range file.Programs {
		fieldName := names.JoinCamelCase(names.SplitCamelCase(p.Name), false)
		o.WriteLine("public var " + fieldName + ":" + p.Name + "ShaderProgram;")
	}
	o.EndOfLine()
	o.WriteLine("public function new(pipeline:RenderPipeline) {")
	o.Indent()

	for _, p := range file.Programs {
		fieldName := names.JoinCamelCase(names.SplitCamelCase(p.Name), false)
		clsName := p.Name + "ShaderProgram"

		o.WriteLine("{")
		o.Indent()
		o.WriteLine("var prog = pipeline.createShaderProgram(" + clsName + ".VERTEX_SHADER_SOURCE, " +
			clsName + ".FRAGMENT_SHADER_SOURCE, " +
			clsName + ".ATTRIBUTE_NAMES, " +
			clsName + ".SAMPLER_NAMES);")
		o.WriteLine(fieldName + " = new " + clsName + "(prog);")
		o.Dedent()
		o.WriteLine("}")

	}
	o.Dedent()
	o.WriteLine("}")

	o.Dedent()
	o.WriteLine("}")
}
