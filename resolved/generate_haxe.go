package resolved

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/ncbray/compilerutil/names"
	"github.com/ncbray/compilerutil/writer"
	"github.com/ncbray/lumen/glsl"
)

func glslTypeToHaxe(t Type) string {
	switch t := t.(type) {
	case *ScalarType:
		switch t.Name {
		case "float":
			return "Float"
		default:
			panic(t.Name)
		}
	case *VectorType:
		switch t.Name {
		case "vec2":
			return "Vector2"
		case "vec4":
			return "Vector4"
		default:
			panic(t.Name)
		}
	case *MatrixType:
		switch t.Name {
		case "mat3":
			return "Matrix3"
		case "mat4":
			return "Matrix4"
		default:
			panic(t.Name)
		}
	default:
		panic(t)
	}
}

func encodedSize(encoding string) int {
	switch encoding {
	case "float":
		return 4
	case "uint16":
		return 2
	case "uint8":
		return 1
	default:
		panic(encoding)
	}
}

func sizeInBytes(t *ScalarType) int {
	switch t.Name {
	case "float":
		return 4
	default:
		panic(t)
	}
}

func align(offset int, size int) int {
	return (offset + size - 1) & ^(size - 1)
}

func GenerateHaxe(pkg string, file *File, declPrecision bool, minify bool, out io.Writer) {
	for _, v := range file.Vertex {

		// Calculate offsets.
		firstByteSize := 1
		offset := 0
		for i, comp := range v.Components {
			encodedScalarByteSize := encodedSize(comp.Encoding)
			switch t := comp.Type.(type) {
			case *VectorType:
				scalarByteSize := encodedScalarByteSize
				if i == 1 {
					firstByteSize = scalarByteSize
				}
				offset = align(offset, scalarByteSize)
				comp.ByteOffset = offset
				comp.ByteSize = scalarByteSize * t.Width

				offset += comp.ByteSize
			default:
				panic(t)
			}
		}
		v.ByteSize = align(offset, firstByteSize)
	}

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
		"mecha.graphics.state.Batch",
		"mecha.graphics.build.DataBuffer",
		"mecha.graphics.build.IndexBuilder",
		"mecha.graphics.state.RenderPipeline",
		"mecha.graphics.state.VertexArray",
	}

	o.EndOfLine()
	for _, imp := range imports {
		o.WriteLine("import " + imp + ";")
	}

	for _, v := range file.Vertex {
		o.EndOfLine()
		o.WriteLine("class " + v.Name + "Builder {")
		o.Indent()
		o.WriteLine("private var data:DataBuffer;")
		o.WriteLine("private var vao:VertexArray;")
		o.WriteLine("private var index:Int;")

		o.EndOfLine()
		o.WriteLine("public function new(pipeline:RenderPipeline) {")
		o.Indent()

		o.WriteLine("data = new DataBuffer(256);")

		o.WriteLine(fmt.Sprintf("vao = new VertexArray(%d);", len(v.Components)))
		o.WriteLine("vao.buffer = pipeline.createBuffer(GL.ARRAY_BUFFER, GL.DYNAMIC_DRAW);")

		for i, comp := range v.Components {
			o.EndOfLine()
			switch t := comp.Type.(type) {
			case *VectorType:
				o.WriteLine(fmt.Sprintf("vao.attrib[%d].size = %d;", i, t.Width))
			default:
				panic(t)
			}
			switch comp.Encoding {
			case "float":
				o.WriteLine(fmt.Sprintf("vao.attrib[%d].type = F32;", i))
				o.WriteLine(fmt.Sprintf("vao.attrib[%d].normalized = false;", i))
			case "uint16":
				o.WriteLine(fmt.Sprintf("vao.attrib[%d].type = U16;", i))
				o.WriteLine(fmt.Sprintf("vao.attrib[%d].normalized = true;", i))
			case "uint8":
				o.WriteLine(fmt.Sprintf("vao.attrib[%d].type = U8;", i))
				o.WriteLine(fmt.Sprintf("vao.attrib[%d].normalized = true;", i))
			default:
				panic(comp.Encoding)
			}
			// Interleaved
			o.WriteLine(fmt.Sprintf("vao.attrib[%d].stride = %d;", i, v.ByteSize))
			o.WriteLine(fmt.Sprintf("vao.attrib[%d].offset = %d;", i, comp.ByteOffset))
		}

		o.Dedent()
		o.WriteLine("}")

		args := []string{}
		for _, comp := range v.Components {
			switch t := comp.Type.(type) {
			case *VectorType:
				for i := 0; i < t.Width; i++ {
					args = append(args, fmt.Sprintf("%s_%d:Float", comp.Name, i))
				}
			default:
				panic(t)
			}
		}

		o.EndOfLine()
		o.WriteLine("public function beginBuffer() {")
		o.Indent()
		o.WriteLine("data.reset();")
		o.WriteLine("index = 0;")
		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("public function beginBatch(batch:Batch) {")
		o.Indent()
		o.WriteLine("batch.vao = vao;")
		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("public function vert(" + strings.Join(args, ", ") + "):Int {")
		o.Indent()

		o.WriteLine(fmt.Sprintf("var offset = data.alloc(%d);", v.ByteSize))
		for _, comp := range v.Components {
			o.EndOfLine()
			switch t := comp.Type.(type) {
			case *VectorType:
				for i := 0; i < t.Width; i++ {
					switch comp.Encoding {
					case "float":
						o.WriteLine(fmt.Sprintf("data.f32[(offset+%d)>>2] = %s_%d;", comp.ByteOffset+i*4, comp.Name, i))
					case "uint16":
						o.WriteLine(fmt.Sprintf("data.u16[(offset+%d)>>1] = Math.round(%s_%d*0xffff);", comp.ByteOffset+i*2, comp.Name, i))
					case "uint8":
						o.WriteLine(fmt.Sprintf("data.u8[(offset+%d)] = Math.round(%s_%d*0xff);", comp.ByteOffset+i, comp.Name, i))
					default:
						panic(comp.Encoding)
					}
				}
			default:
				panic(t)
			}
		}

		o.EndOfLine()
		o.WriteLine("var temp = index;")
		o.WriteLine("index += 1;")
		o.WriteLine("return temp;")

		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("public function endBatch(batch:Batch) {")
		o.Indent()
		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("public function endBuffer(pipeline:RenderPipeline) {")
		o.Indent()
		o.WriteLine("pipeline.uploadBuffer(vao.buffer, data.u8.subarray(0, data.used), data.used);")
		o.Dedent()
		o.WriteLine("}")

		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("interface " + v.Name + "Generator {")
		o.Indent()
		o.WriteLine("public function buildBuffer(index:IndexBuilder, data:" + v.Name + "Builder):Void;")
		o.Dedent()
		o.WriteLine("}")
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
			case *VectorType:
				switch t.Name {
				case "vec4":
					o.WriteLine("GL.uniform4f(prog." + n + ", " + n + ".x, " + n + ".y, " + n + ".z, " + n + ".w);")
				default:
					panic(t.Name)
				}
			case *MatrixType:
				switch t.Name {
				case "mat4":
					o.WriteLine("GL.uniformMatrix4fv(prog." + n + ", false, " + n + ");")
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
