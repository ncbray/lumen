package resolved

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/ncbray/compilerutil/writer"
	"github.com/ncbray/lumen/glsl"
)

func glslTypeToTypescript(t Type) string {
	switch t := t.(type) {
	case *ScalarType:
		switch t.Name {
		case "float":
			return "number"
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

func fieldNameList(fields []*Field) string {
	names := make([]string, len(fields))
	for i, f := range fields {
		names[i] = strconv.Quote(f.Name)
	}
	return strings.Join(names, ", ")
}

func GenerateTypeScript(file *File, minify bool, out io.Writer) {
	calcOffsets(file)

	o := writer.MakeTabbedWriter("  ", out)

	o.WriteLine("import { Matrix4, Vector4 } from './vapor/math';")
	o.WriteLine("import { Buffer, Geometry, GraphicsPipeline, TextureUnit, VertexArray } from './vapor/graphics';")
	o.WriteLine("import { GL } from './vapor/webgl';")

	for _, v := range file.Vertex {
		o.EndOfLine()
		o.WriteLine("export class " + v.Name + "Builder {")
		o.Indent()
		o.WriteLine("private raw_data:ArrayBuffer;")
		o.WriteLine("private u8:Uint8Array;")
		o.WriteLine("private u16:Uint16Array;")
		o.WriteLine("private f32:Float32Array;")
		o.WriteLine("private used:number;")
		o.WriteLine("private vertex_count:number;")

		o.WriteLine("private gpu_buffer:Buffer;")
		o.WriteLine("private vao:VertexArray;")

		o.EndOfLine()
		o.WriteLine("constructor(gpu_buffer:Buffer) {")
		o.Indent()

		o.WriteLine("this.raw_data = new ArrayBuffer(256);")
		o.WriteLine("this.syncViews();")
		o.WriteLine("this.used = 0;")

		o.WriteLine(fmt.Sprintf("const vao = new VertexArray(%d);", len(v.Components)))
		o.WriteLine("var a;")
		for i, comp := range v.Components {
			o.EndOfLine()
			o.WriteLine(fmt.Sprintf("a = vao.attrib[%d];", i))
			switch t := comp.Type.(type) {
			case *VectorType:
				o.WriteLine(fmt.Sprintf("a.size = %d;", t.Width))
			default:
				panic(t)
			}
			switch comp.Encoding {
			case "float":
				o.WriteLine("a.type = GL.FLOAT;")
				o.WriteLine("a.normalized = false;")
			case "uint16":
				o.WriteLine("a.type = GL.UNSIGNED_SHORT;")
				o.WriteLine("a.normalized = true;")
			case "uint8":
				o.WriteLine("a.type = GL.UNSIGNED_BYTE;")
				o.WriteLine("a.normalized = true;")
			default:
				panic(comp.Encoding)
			}
			// Interleaved
			o.WriteLine(fmt.Sprintf("a.stride = %d;", v.ByteSize))
			o.WriteLine(fmt.Sprintf("a.offset = %d;", comp.ByteOffset))
			o.WriteLine("a.buffer = gpu_buffer;")
		}
		o.EndOfLine()
		o.WriteLine("this.gpu_buffer = gpu_buffer;")
		o.WriteLine("this.vao = vao;")

		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteChunk(`
			geometry(geometry:Geometry) {
				geometry.vao = this.vao;
			}

			uploadData(pipeline:GraphicsPipeline) {
				pipeline.bufferData(this.gpu_buffer, this.u8.subarray(0, this.used));
				this.used = 0;
				this.vertex_count = 0;
			}

			private alloc(required:number):number {
				const temp = this.used;
				this.used += required;
				if (this.used > this.raw_data.byteLength) {
					this.raw_data = new ArrayBuffer(this.raw_data.byteLength * 2);
					const old = this.u8;
					this.syncViews();
					this.u8.set(old.subarray(0, temp));
				}
				return temp;
			}

			private syncViews() {
				this.u8 = new Uint8Array(this.raw_data);
				this.u16 = new Uint16Array(this.raw_data);
				this.f32 = new Float32Array(this.raw_data);
			}
		`)

		o.EndOfLine()

		params := []string{}
		for _, comp := range v.Components {
			switch t := comp.Type.(type) {
			case *VectorType:
				for j := 0; j < t.Width; j++ {
					params = append(params, fmt.Sprintf("%s_%d:number", comp.Name, j))
				}
			default:
				panic(t)
			}
		}

		o.WriteLine(fmt.Sprintf("vert(%s):number {", strings.Join(params, ", ")))
		o.Indent()
		o.WriteLine(fmt.Sprintf("const offset = this.alloc(%d);", v.ByteSize))

		for _, comp := range v.Components {
			o.EndOfLine()
			switch t := comp.Type.(type) {
			case *VectorType:
				for j := 0; j < t.Width; j++ {
					param := fmt.Sprintf("%s_%d", comp.Name, j)
					switch comp.Encoding {
					case "float":
						o.WriteLine(fmt.Sprintf("this.f32[(offset+%d)>>2] = %s;", comp.ByteOffset+j*4, param))
					case "uint16":
						o.WriteLine(fmt.Sprintf("this.u16[(offset+%d)>>1] = %s * 65535 + 0.5;", comp.ByteOffset+j*2, param))
					case "uint8":
						o.WriteLine(fmt.Sprintf("this.u8[(offset+%d)>>0] = %s * 255 + 0.5;", comp.ByteOffset+j*1, param))
					default:
						panic(comp.Encoding)
					}
				}
			default:
				panic(t)
			}
		}

		o.EndOfLine()
		o.WriteLine("const temp = this.vertex_count;")
		o.WriteLine("this.vertex_count += 1;")
		o.WriteLine("return temp;")
		o.Dedent()
		o.WriteLine("}")

		// End of class.
		o.Dedent()
		o.WriteLine("}")
	}

	for _, p := range file.Programs {
		const declPrecision = true
		var vs bytes.Buffer
		GenerateGLSLVertShader(p, declPrecision, &vs)
		var fs bytes.Buffer
		GenerateGLSLFragShader(p, declPrecision, &fs)

		shaderClassName := p.Name + "ShaderProgram"
		uniformsClassName := p.Name + "ShaderUniforms"

		// Shader
		uniformData, uniformSamplers := splitDataAndSamplers(p.Uniform.Fields)

		o.EndOfLine()
		o.WriteLine("export class " + shaderClassName + " {")
		o.Indent()

		o.WriteLine("readonly prog:WebGLProgram;")
		for _, u := range uniformData {
			o.WriteLine("readonly " + u.Name + ":WebGLUniformLocation;")
		}

		vsSrc := vs.String()
		fsSrc := fs.String()
		if minify {
			vsSrc = glsl.CompactSource(vsSrc)
			fsSrc = glsl.CompactSource(fsSrc)
		}

		o.EndOfLine()
		o.WriteLine("constructor(pipeline:GraphicsPipeline) {")
		o.Indent()

		o.WriteChunk(fmt.Sprintf(`
			this.prog = pipeline.createShaderProgram(
				%q,
				%q,
				[%s],
				[%s],
			);
`, vsSrc, fsSrc, fieldNameList(p.Attribute.Fields), fieldNameList(uniformSamplers)))

		o.WriteLine("const gl = pipeline.gl;")
		for _, u := range uniformData {
			o.WriteLine(fmt.Sprintf("this.%s = gl.getUniformLocation(this.prog, %q)!;", u.Name, u.Name))
		}

		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("createUniforms():" + uniformsClassName + " {")
		o.Indent()
		o.WriteLine("return new " + uniformsClassName + "(this);")
		o.Dedent()
		o.WriteLine("}")

		// End of class.
		o.Dedent()
		o.WriteLine("}")

		// Uniforms
		o.EndOfLine()
		o.WriteLine("export class " + uniformsClassName + " {")
		o.Indent()

		o.WriteLine("private prog:" + shaderClassName + ";")
		o.WriteLine("private units:Array<TextureUnit>;")

		for _, u := range uniformData {
			ht := glslTypeToTypescript(u.Type)
			init := "new " + ht + "()"
			o.WriteLine(u.Name + ":" + ht + " = " + init + ";")
		}
		for _, u := range uniformSamplers {
			o.WriteLine(u.Name + ":TextureUnit = new TextureUnit();")
		}

		o.EndOfLine()
		o.WriteLine("constructor(prog:" + shaderClassName + ") {")
		o.Indent()
		o.WriteLine("this.prog = prog;")
		o.WriteString("this.units = [")
		for i, u := range uniformSamplers {
			if i != 0 {
				o.WriteString(", ")
			}
			o.WriteString("this." + u.Name)
		}
		o.WriteString("];")
		o.EndOfLine()

		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteLine("bind(pipeline:GraphicsPipeline) {")
		o.Indent()
		o.WriteLine("const gl = pipeline.gl;")
		o.WriteLine("const prog = this.prog;")
		o.WriteLine("gl.useProgram(prog.prog);")
		for _, u := range uniformData {
			n := u.Name
			switch t := u.Type.(type) {
			case *VectorType:
				switch t.Name {
				case "vec4":
					o.WriteLine("const " + n + " = this." + n + ";")
					o.WriteLine("gl.uniform4f(prog." + n + ", " + n + ".x, " + n + ".y, " + n + ".z, " + n + ".w);")
				default:
					panic(t.Name)
				}
			case *MatrixType:
				switch t.Name {
				case "mat4":
					o.WriteLine("this." + n + ".copyToMemory(pipeline.matrix4Temp, 0);")
					o.WriteLine("gl.uniformMatrix4fv(prog." + n + ", false, pipeline.matrix4Temp);")
				default:
					panic(t.Name)
				}
			default:
				panic(t)
			}
		}
		o.WriteLine("pipeline.bindTextures(this.units);")
		o.Dedent()
		o.WriteLine("}")

		// End of class.
		o.Dedent()
		o.WriteLine("}")

	}
}
