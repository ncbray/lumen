package resolved

import (
	"io"
	"strings"

	"github.com/ncbray/compilerutil/writer"
)

type generateGoStructs struct {
	out   *writer.TabbedWriter
	isAst bool
}

func (g *generateGoStructs) genStructTypeRef(t Type) {
	switch t := t.(type) {
	case *IntrinsicType:
		switch t.Name {
		case "location":
			g.out.WriteString("util.Location")
		default:
			g.out.WriteString(t.Name)
		}
	case *Struct:
		g.out.WriteString("*")
		g.out.WriteString(t.Name)
	case *Holder:
		g.out.WriteString(t.Name)
	case *List:
		g.out.WriteString("[]")
		g.genStructTypeRef(t.Element)
	default:
		panic(t)
	}

}

func (g *generateGoStructs) genStructField(f *Field, largestName int) {
	g.out.WriteString(f.Name)
	g.out.WriteString(strings.Repeat(" ", largestName-len(f.Name)+1))
	g.genStructTypeRef(f.Type)
	g.out.EndOfLine()
}

func (g *generateGoStructs) genStructFile(file *File) {
	for _, d := range file.Types {
		switch d := d.(type) {
		case *Holder:
			g.out.EndOfLine()
			g.out.WriteLine("type " + d.Name + " interface {")
			g.out.Indent()
			g.out.WriteLine("is" + d.Name + "()")
			g.out.Dedent()
			g.out.WriteLine("}")
		case *Struct:
			g.out.EndOfLine()
			g.out.WriteLine("type " + d.Name + " struct {")
			g.out.Indent()

			largestName := 4 // Temp
			for _, f := range d.Fields {
				l := len(f.Name)
				if l > largestName {
					largestName = l
				}
			}

			// HACK to make translation simpler.
			g.out.WriteLine("Temp" + strings.Repeat(" ", largestName-4+1) + "interface{}")

			for _, f := range d.Fields {
				g.genStructField(f, largestName)
			}
			g.out.Dedent()
			g.out.WriteLine("}")

			for _, h := range d.Holders {
				g.out.EndOfLine()
				g.out.WriteLine("func (n *" + d.Name + ") is" + h.Name + "() {")
				g.out.Indent()
				g.out.Dedent()
				g.out.WriteLine("}")
			}
		default:
			panic(d)
		}
	}
}

func capitalize(name string) string {
	return strings.ToUpper(name[0:1]) + name[1:]
}

func (g *generateGoStructs) genBindingExpr(e ParserBindingExpr) {
	switch e := e.(type) {
	case *GetLocStart:
		g.out.WriteString("util.GetLocation(c.Filename, ctx.GetStart())")
	case *GetParam:
		pname := capitalize(e.Param.Name)
		switch i := e.Param.Input.(type) {
		case *InputSlice:
			g.out.WriteString("ctx.Get" + pname + "().GetText()")
		case *ParserBindingGroup:
			g.out.WriteString("c.Convert" + capitalize(i.Name) + "(ctx.Get" + pname + "())")
		case *InputList:
			g.out.WriteString("c.Convert" + capitalize(i.Element.Name) + "List(ctx.Get" + pname + "())")
		default:
			panic(i)
		}
	case *Construct:
		g.out.WriteString("&" + e.Type.Name + "{")
		if len(e.Args) > 0 {
			g.out.EndOfLine()
			g.out.Indent()
			longestName := 0
			for _, a := range e.Args {
				l := len(a.Field.Name)
				if l > longestName {
					longestName = l
				}
			}

			for _, a := range e.Args {
				fn := a.Field.Name
				g.out.WriteString(fn + ":" + strings.Repeat(" ", longestName-len(fn)+1))
				g.genBindingExpr(a.Value)
				g.out.WriteString(",")
				g.out.EndOfLine()
			}
			g.out.Dedent()
		}
		g.out.WriteString("}")
	default:
		panic(e)
	}
}

func (g *generateGoStructs) genParserBinding(bindings *ParserBinding) {
	convCls := "ASTConverter"

	g.out.EndOfLine()
	g.out.WriteLine("type " + convCls + " struct {")
	g.out.Indent()
	g.out.WriteLine("Filename string")
	g.out.Dedent()
	g.out.WriteLine("}")

	for _, group := range bindings.Groups {
		g.out.EndOfLine()
		cname := capitalize(group.Name)
		g.out.WriteString("func (c *" + convCls + ") Convert" + cname + "(ctx parser.I" + cname + "Context) ")
		g.genStructTypeRef(group.Type)
		g.out.WriteString(" {")
		g.out.EndOfLine()
		g.out.Indent()

		g.out.WriteLine("switch ctx := ctx.(type) {")
		for _, m := range group.Mappings {
			rule := cname
			if m.Rule != "default" {
				rule = capitalize(m.Rule)
			}
			g.out.WriteLine("case *parser." + rule + "Context:")
			g.out.Indent()
			g.out.WriteString("return ")
			g.genBindingExpr(m.Body)
			g.out.EndOfLine()
			g.out.Dedent()
		}
		g.out.WriteLine("default:")
		g.out.Indent()
		g.out.WriteLine("panic(ctx)")
		g.out.Dedent()
		g.out.WriteLine("}")

		g.out.Dedent()
		g.out.WriteLine("}")

		if group.List != nil {
			g.out.EndOfLine()
			g.out.WriteString("func (c *" + convCls + ") Convert" + cname + "List(src []parser.I" + cname + "Context) []")
			g.genStructTypeRef(group.Type)
			g.out.WriteString(" {")
			g.out.EndOfLine()
			g.out.Indent()

			g.out.WriteString("dst := make([]")
			g.genStructTypeRef(group.Type)
			g.out.WriteString(", len(src))")
			g.out.EndOfLine()

			g.out.WriteLine("for i, child := range src {")
			g.out.Indent()
			g.out.WriteLine("dst[i] = c.Convert" + cname + "(child)")
			g.out.Dedent()
			g.out.WriteLine("}")

			g.out.WriteLine("return dst")
			g.out.Dedent()
			g.out.WriteLine("}")
		}
	}
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

func GenerateGo(pkg string, file *File, parserPackage string, out io.Writer) {
	isAst := file.ParserBinding != nil

	g := &generateGoStructs{
		out:   writer.MakeTabbedWriter("\t", out),
		isAst: isAst,
	}

	usesLocation := false
	for _, t := range file.Types {
		switch t := t.(type) {
		case *Struct:
			for _, f := range t.Fields {
				switch ft := f.Type.(type) {
				case *IntrinsicType:
					if ft.Name == "location" {
						usesLocation = true
					}
				}
			}
		}
	}

	imports := []string{}
	if isAst {
		imports = append(imports, parserPackage)
	}
	if usesLocation {
		imports = append(imports, "github.com/ncbray/lumen/util")
	}
	generateGoHeader(pkg, imports, g.out)

	g.genStructFile(file)
	if file.ParserBinding != nil {
		g.genParserBinding(file.ParserBinding)
	}
}
