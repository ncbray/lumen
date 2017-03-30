package ast

import (
	"io"

	"github.com/ncbray/compilerutil/writer"
)

type generateGoStructs struct {
	out   *writer.TabbedWriter
	lut   map[string]Declaration
	isAst bool
}

func (g *generateGoStructs) genStructTypeRef(t TypeRef) {
	switch t := t.(type) {
	case *TypeName:
		isStruct := false
		d, ok := g.lut[t.Name]
		if ok {
			_, ok = d.(*StructDecl)
			if ok {
				isStruct = true
			}
		}
		if isStruct {
			g.out.WriteString("*")
		}
		g.out.WriteString(t.Name)
	case *ListRef:
		g.out.WriteString("[]")
		g.genStructTypeRef(t.Element)
	default:
		panic(t)
	}

}

func (g *generateGoStructs) genStructMemberDecl(m MemberDecl) {
	switch m := m.(type) {
	case *FieldDecl:
		g.out.WriteString(m.Name)
		g.out.WriteString(" ")
		g.genStructTypeRef(m.T)
		g.out.EndOfLine()
	default:
		panic(m)
	}
}

func (g *generateGoStructs) genStructFile(file *File) {
	// Index
	for _, d := range file.Decls {
		switch d := d.(type) {
		case *EnumDecl:
			g.lut[d.Name] = d
		case *StructDecl:
			g.lut[d.Name] = d
		default:
			panic(d)
		}
	}

	generateGoHeader("ast", []string{"github.com/ncbray/reboot/util"}, g.out)

	for _, d := range file.Decls {
		switch d := d.(type) {
		case *EnumDecl:
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
				}
				for _, m := range v.Members {
					g.genStructMemberDecl(m)
				}
				g.out.Dedent()
				g.out.WriteLine("}")

				g.out.EndOfLine()
				g.out.WriteLine("func (n *" + v.Name + ") is" + d.Name + "() {")
				g.out.Indent()
				g.out.Dedent()
				g.out.WriteLine("}")
			}
		case *StructDecl:
			g.out.EndOfLine()
			g.out.WriteLine("type " + d.Name + " struct {")
			g.out.Indent()
			if g.isAst {
				g.out.WriteLine("Loc util.Location")
			}
			for _, m := range d.Members {
				g.genStructMemberDecl(m)
			}
			g.out.Dedent()
			g.out.WriteLine("}")
		default:
			panic(d)
		}
	}
}

func (g *generateGoStructs) genConvMember(m MemberDecl) {
	switch m := m.(type) {
	case *FieldDecl:
		g.out.WriteString(m.Name)
		g.out.WriteString(": ")
		switch t := m.T.(type) {
		case *TypeName:
			if t.Name == "string" {
				g.out.WriteString("ctx.Get" + m.Name + "().GetText()")
			} else {
				g.out.WriteString("c.Convert" + t.Name + "(ctx.Get" + m.Name + "())")
			}
		case *ListRef:
			elem, ok := t.Element.(*TypeName)
			if !ok {
				panic(t.Element)
			}
			g.out.WriteString("c.Convert" + elem.Name + "List(ctx.Get" + m.Name + "())")
		default:
			panic(t)
		}
		g.out.WriteString(",")
		g.out.EndOfLine()
	default:
		panic(m)
	}
}

func (g *generateGoStructs) genConvFile(file *File) {
	generateGoHeader("ast", []string{
		"github.com/ncbray/reboot/parser",
		"github.com/ncbray/reboot/util",
	}, g.out)

	convCls := "ASTConverter"

	g.out.EndOfLine()
	g.out.WriteLine("type " + convCls + " struct {")
	g.out.Indent()
	g.out.WriteLine("Filename string")
	g.out.Dedent()
	g.out.WriteLine("}")

	for _, d := range file.Decls {
		switch d := d.(type) {
		case *EnumDecl:
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
				for _, m := range v.Members {
					g.genConvMember(m)
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

		case *StructDecl:
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
			for _, m := range d.Members {
				g.genConvMember(m)
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

func GenerateGo(file *File, isAst bool, out io.Writer) {
	g := &generateGoStructs{
		out:   writer.MakeTabbedWriter("\t", out),
		lut:   map[string]Declaration{},
		isAst: isAst,
	}
	g.genStructFile(file)

	if isAst {
		g.out.WriteLine("================================================================================")
		g.genConvFile(file)
	}
}
