package resolved

import (
	"io"

	"github.com/ncbray/compilerutil/writer"
)

type generateGoStructs struct {
	out   *writer.TabbedWriter
	isAst bool
}

func (g *generateGoStructs) genStructTypeRef(t Type) {
	switch t := t.(type) {
	case *Intrinsic:
		g.out.WriteString(t.Name)
	case *Struct:
		g.out.WriteString("*")
		g.out.WriteString(t.Name)
	case *Enum:
		g.out.WriteString(t.Name)
	case *List:
		g.out.WriteString("[]")
		g.genStructTypeRef(t.Element)
	default:
		panic(t)
	}

}

func (g *generateGoStructs) genStructField(f *Field) {
	g.out.WriteString(f.Name)
	g.out.WriteString(" ")
	g.genStructTypeRef(f.Type)
	g.out.EndOfLine()
}

func (g *generateGoStructs) genStructFile(file *File) {
	for _, d := range file.Types {
		switch d := d.(type) {
		case *Enum:
			g.out.EndOfLine()
			g.out.WriteLine("type " + d.Name + " interface {")
			g.out.Indent()
			g.out.WriteLine("is" + d.Name + "()")
			g.out.Dedent()
			g.out.WriteLine("}")

			for _, v := range d.Variants {
				g.out.EndOfLine()
				g.out.WriteLine("type " + v.Name + " struct {")
				g.out.Indent()
				if g.isAst {
					g.out.WriteLine("Loc util.Location")
					// HACK to make translation simpler.
					g.out.WriteLine("Temp interface{}")
				}
				for _, f := range v.Fields {
					g.genStructField(f)
				}
				g.out.Dedent()
				g.out.WriteLine("}")

				g.out.EndOfLine()
				g.out.WriteLine("func (n *" + v.Name + ") is" + d.Name + "() {")
				g.out.Indent()
				g.out.Dedent()
				g.out.WriteLine("}")
			}
		case *Struct:
			g.out.EndOfLine()
			g.out.WriteLine("type " + d.Name + " struct {")
			g.out.Indent()
			if g.isAst {
				g.out.WriteLine("Loc util.Location")
				// HACK to make translation simpler.
				g.out.WriteLine("Temp interface{}")
			}
			for _, f := range d.Fields {
				g.genStructField(f)
			}
			g.out.Dedent()
			g.out.WriteLine("}")
		default:
			panic(d)
		}
	}
}

func (g *generateGoStructs) genConvField(f *Field) {
	name := f.Name
	g.out.WriteString(name)
	g.out.WriteString(": ")
	switch t := f.Type.(type) {
	case *Intrinsic:
		if t.Name != "string" {
			panic(t.Name)
		}
		g.out.WriteString("ctx.Get" + name + "().GetText()")
	case *Struct:
		g.out.WriteString("c.Convert" + t.Name + "(ctx.Get" + name + "())")
	case *Enum:
		g.out.WriteString("c.Convert" + t.Name + "(ctx.Get" + name + "())")
	case *List:
		var elName string
		switch et := t.Element.(type) {
		case *Intrinsic:
			elName = et.Name
		case *Enum:
			elName = et.Name
		case *Struct:
			elName = et.Name
		default:
			panic(et)
		}
		g.out.WriteString("c.Convert" + elName + "List(ctx.Get" + name + "())")
	default:
		panic(t)
	}
	g.out.WriteString(",")
	g.out.EndOfLine()
}

func (g *generateGoStructs) genConvFile(file *File) {
	convCls := "ASTConverter"

	g.out.EndOfLine()
	g.out.WriteLine("type " + convCls + " struct {")
	g.out.Indent()
	g.out.WriteLine("Filename string")
	g.out.Dedent()
	g.out.WriteLine("}")

	for _, d := range file.Types {
		switch d := d.(type) {
		case *Enum:
			g.out.EndOfLine()
			g.out.WriteLine("func (c *" + convCls + ") Convert" + d.Name + "(ctx parser.I" + d.Name + "Context) " + d.Name + " {")
			g.out.Indent()
			g.out.WriteLine("switch ctx := ctx.(type) {")

			for _, v := range d.Variants {
				g.out.WriteLine("case *parser." + v.Name + "Context:")
				g.out.Indent()
				g.out.WriteLine("return &" + v.Name + " {")
				g.out.Indent()
				if g.isAst {
					g.out.WriteLine("Loc: util.GetLocation(c.Filename, ctx.GetStart()),")
				}
				for _, f := range v.Fields {
					g.genConvField(f)
				}
				g.out.Dedent()
				g.out.WriteLine("}")
				g.out.Dedent()
			}

			g.out.WriteLine("default:")
			g.out.Indent()
			g.out.WriteLine("panic(ctx)")
			g.out.Dedent()
			g.out.WriteLine("}")

			g.out.Dedent()
			g.out.WriteLine("}")

			generateConvList(convCls, d.Name, false, g.out)

		case *Struct:
			g.out.EndOfLine()
			g.out.WriteLine("func (c *" + convCls + ") Convert" + d.Name + "(ctx parser.I" + d.Name + "Context) *" + d.Name + " {")
			g.out.Indent()

			g.out.WriteLine("switch ctx := ctx.(type) {")

			g.out.WriteLine("case *parser." + d.Name + "Context:")
			g.out.Indent()
			g.out.WriteLine("return &" + d.Name + " {")
			g.out.Indent()
			if g.isAst {
				g.out.WriteLine("Loc: util.GetLocation(c.Filename, ctx.GetStart()),")
			}
			for _, f := range d.Fields {
				g.genConvField(f)
			}
			g.out.Dedent()
			g.out.WriteLine("}")
			g.out.Dedent()

			g.out.WriteLine("default:")
			g.out.Indent()
			g.out.WriteLine("panic(ctx)")
			g.out.Dedent()
			g.out.WriteLine("}")
			g.out.Dedent()
			g.out.WriteLine("}")
			generateConvList(convCls, d.Name, true, g.out)
		default:
			panic(d)
		}
	}
}

func generateConvList(cls string, name string, ptr bool, out *writer.TabbedWriter) {
	t := name
	if ptr {
		t = "*" + t
	}
	out.EndOfLine()
	out.WriteLine("func (c *" + cls + ") Convert" + name + "List(src []parser.I" + name + "Context) []" + t + " {")
	out.Indent()
	out.WriteLine("dst := make([]" + t + ", len(src))")

	out.WriteLine("for i, child := range src {")
	out.Indent()
	out.WriteLine("dst[i] = c.Convert" + name + "(child)")
	out.Dedent()
	out.WriteLine("}")

	out.WriteLine("return dst")
	out.Dedent()
	out.WriteLine("}")
}

func generateGoHeader(pkg string, imports []string, out *writer.TabbedWriter) {
	out.WriteLine("package " + pkg)

	if len(imports) > 0 {
		out.EndOfLine()
		out.WriteLine("import (")
		out.Indent()
		for _, imp := range imports {
			out.WriteLine("\"" + imp + "\"")
		}
		out.Dedent()
		out.WriteLine(")")
	}
}

func GenerateGo(pkg string, file *File, isAst bool, parserPackage string, out io.Writer) {
	g := &generateGoStructs{
		out:   writer.MakeTabbedWriter("\t", out),
		isAst: isAst,
	}

	imports := []string{}
	if isAst {
		imports = append(imports, parserPackage)
		imports = append(imports, "github.com/ncbray/lumen/util")
	}
	generateGoHeader(pkg, imports, g.out)

	g.genStructFile(file)
	if isAst {
		g.genConvFile(file)
	}
}
