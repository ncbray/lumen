package resolved

import (
	"bytes"
	"fmt"
	"io"

	"github.com/ncbray/compilerutil/names"
	"github.com/ncbray/compilerutil/writer"
	"github.com/ncbray/lumen/glsl"
)

func glslTypeToHaxe(t Type) string {
	switch t := t.(type) {
	case *IntrinsicType:
		switch t.Name {
		case "float":
			return "Float"
		case "mat3":
			return "Matrix3"
		case "mat4":
			return "Matrix4"
		case "vec2":
			return "Vector2"
		case "vec4":
			return "Vector4"
		default:
			panic(t.Name)
		}
	default:
		panic(t)
	}
}

func GenerateHaxe(pkg string, file *File, declPrecision bool, minify bool, out io.Writer) {
	o := writer.MakeTabbedWriter("  ", out)

	o.WriteLine("package " + pkg + ";")

	imports := []string{
		"lime.graphics.opengl.GL",
		"lime.graphics.opengl.GLProgram",
		"lime.graphics.opengl.GLUniformLocation",
		"lime.math.Matrix3",
		"lime.math.Matrix4",
		"lime.math.Vector2",
		"lime.math.Vector4",
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

		shaderClassName := p.Name + "ShaderProgram"
		uniformsClassName := p.Name + "ShaderUniforms"

		o.EndOfLine()
		o.WriteLine("class " + shaderClassName + " {")
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

		vsSrc := vs.String()
		fsSrc := fs.String()

		if minify {
			vsSrc = glsl.CompactSource(vsSrc)
			fsSrc = glsl.CompactSource(fsSrc)
		}

		o.WriteLine(fmt.Sprintf("public static inline var VERTEX_SHADER_SOURCE = %q;", vsSrc))
		o.WriteLine(fmt.Sprintf("public static inline var FRAGMENT_SHADER_SOURCE = %q;", fsSrc))

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

		o.EndOfLine()
		o.WriteLine("public function createUniforms() {")
		o.Indent()
		o.WriteLine("return new " + uniformsClassName + "(this);")
		o.Dedent()
		o.WriteLine("}")

		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("class " + uniformsClassName + " {")
		o.Indent()

		o.WriteLine("private var prog:" + shaderClassName + ";")
		for _, u := range uniformData {
			ht := glslTypeToHaxe(u.Type)
			init := "new " + ht + "()"
			o.WriteLine("public var " + u.Name + ":" + ht + " = " + init + ";")
		}

		o.EndOfLine()
		o.WriteLine("public function new(prog:" + shaderClassName + ") {")
		o.Indent()
		o.WriteLine("this.prog = prog;")
		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("public function bind() {")
		o.Indent()
		o.WriteLine("GL.useProgram(prog.program);")
		for _, u := range uniformData {
			n := u.Name
			switch t := u.Type.(type) {
			case *IntrinsicType:
				switch t.Name {
				case "mat4":
					o.WriteLine("GL.uniformMatrix4fv(prog." + n + ", false, " + n + ");")
				case "vec4":

					o.WriteLine("GL.uniform4f(prog." + n + ", " + n + ".x, " + n + ".y, " + n + ".z, " + n + ".w);")
				default:
					panic(t.Name)
				}
			default:
				panic(t)
			}
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
