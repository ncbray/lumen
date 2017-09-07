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

	o.WriteLine("import { Buffer, GraphicsPipeline, VertexArray } from './vapor/graphics';")
	o.WriteLine("import { GL } from './vapor/webgl';")

	for _, v := range file.Vertex {
		o.EndOfLine()
		o.WriteLine("export class " + v.Name + "Builder {")
		o.Indent()
		o.WriteLine("private data:ArrayBuffer;")
		o.WriteLine("private u8:Uint8Array;")
		o.WriteLine("private u16:Uint16Array;")
		o.WriteLine("private f32:Float32Array;")
		o.WriteLine("private used:number;")
		o.WriteLine("private vertexCount:number;")

		o.WriteLine("private buffer:Buffer;")
		o.WriteLine("public vao:VertexArray;")

		o.EndOfLine()
		o.WriteLine("constructor(buffer:Buffer) {")
		o.Indent()

		o.WriteLine("this.data = new ArrayBuffer(256);")
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
			o.WriteLine("a.buffer = buffer;")
		}
		o.EndOfLine()
		o.WriteLine("this.buffer = buffer;")
		o.WriteLine("this.vao = vao;")

		o.Dedent()
		o.WriteLine("}")

		o.EndOfLine()
		o.WriteChunk(`
			beginData() {
				this.used = 0;
				this.vertexCount = 0;
			}

			endData(pipeline:GraphicsPipeline) {
				pipeline.bufferData(this.buffer, this.u8.subarray(0, this.used));
			}

			private alloc(required:number):number {
				const temp = this.used;
				this.used += required;
				if (this.used > this.data.byteLength) {
					this.data = new ArrayBuffer(this.data.byteLength * 2);
					const old = this.u8;
					this.syncViews();
					this.u8.set(old.subarray(0, temp));
				}
				return temp;
			}

			private syncViews() {
				this.u8 = new Uint8Array(this.data);
				this.u16 = new Uint16Array(this.data);
				this.f32 = new Float32Array(this.data);
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
		o.WriteLine("const temp = this.vertexCount;")
		o.WriteLine("this.vertexCount += 1;")
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

		// End of class.
		o.Dedent()
		o.WriteLine("}")

		// Uniforms
		o.EndOfLine()
		o.WriteLine("export class " + uniformsClassName + " {")
		o.Indent()

		// End of class.
		o.Dedent()
		o.WriteLine("}")

	}
}
