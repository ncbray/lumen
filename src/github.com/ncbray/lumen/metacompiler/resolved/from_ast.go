package resolved

import (
	"fmt"

	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/metacompiler/ast"
)

type namespace struct {
	parent *namespace
	types  map[string]Type
}

func (ns *namespace) IsDefined(name string) bool {
	_, ok := ns.types[name]
	return ok
}

func (ns *namespace) Define(name string, t Type) {
	ns.types[name] = t
}

func (ns *namespace) ResolveType(name string) Type {
	t, ok := ns.types[name]
	if !ok && ns.parent != nil {
		return ns.parent.ResolveType(name)
	}
	return t
}

func listOfType(t Type) Type {
	switch t := t.(type) {
	case *Intrinsic:
		if t.List == nil {
			t.List = &List{Element: t}
		}
		return t.List
	case *Enum:
		if t.List == nil {
			t.List = &List{Element: t}
		}
		return t.List
	case *Struct:
		if t.List == nil {
			t.List = &List{Element: t}
		}
		return t.List
	case *List:
		if t.List == nil {
			t.List = &List{Element: t}
		}
		return t.List
	default:
		panic(t)
	}
}

func resolveType(ref ast.TypeRef, ns *namespace, logger log.CompilerLogger) Type {
	switch ref := ref.(type) {
	case *ast.TypeName:
		t := ns.ResolveType(ref.Name)
		if t == nil {
			loc := ref.Loc
			logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Unknown type %q.", ref.Name))
		}
		return t
	case *ast.ListRef:
		t := resolveType(ref.Element, ns, logger)
		if t == nil {
			return nil
		}
		return listOfType(t)
	default:
		panic(ref)
	}
}

func resolveMembers(scopeName string, members []ast.MemberDecl, ns *namespace, logger log.CompilerLogger) []*Field {
	existing := map[string]bool{}
	fields := []*Field{}
	for _, m := range members {
		switch m := m.(type) {
		case *ast.FieldDecl:
			if existing[m.Name] {
				loc := m.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Attempted to redefine %q in %s.", m.Name, scopeName))
			}
			existing[m.Name] = true
			t := resolveType(m.T, ns, logger)
			fields = append(fields, &Field{Name: m.Name, Type: t})
		default:
			panic(m)
		}
	}
	return fields
}

func FromAST(src *ast.File, logger log.CompilerLogger) *File {
	builtins := &namespace{
		parent: nil,
		types: map[string]Type{
			"string": &Intrinsic{Name: "string"},
		},
	}
	ns := &namespace{
		parent: builtins,
		types:  map[string]Type{},
	}
	// Generate types.
	types := []Type{}
	for _, d := range src.Decls {
		switch d := d.(type) {
		case *ast.EnumDecl:
			if ns.IsDefined(d.Name) {
				loc := d.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Attempted to redefine %q.", d.Name))
				continue
			}
			t := &Enum{
				Name: d.Name,
			}
			ns.Define(d.Name, t)
			types = append(types, t)
			d.Temp = t

			// TODO check variant names.
			variants := make([]*Variant, len(d.Variants))
			for i, v := range d.Variants {
				vt := &Variant{
					Name: v.Name,
				}
				variants[i] = vt
				v.Temp = vt
			}
			t.Variants = variants
		case *ast.StructDecl:
			if ns.IsDefined(d.Name) {
				loc := d.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Attempted to redefine %q.", d.Name))
				continue
			}
			t := &Struct{
				Name: d.Name,
			}
			ns.Define(d.Name, t)
			types = append(types, t)
			d.Temp = t
		default:
			panic(d)
		}
	}
	if logger.NumErrors() > 0 {
		return nil
	}

	// Resolve fields
	for _, d := range src.Decls {
		switch d := d.(type) {
		case *ast.EnumDecl:
			//e := d.Temp.(*Enum)
			for _, v := range d.Variants {
				vt := v.Temp.(*Variant)
				vt.Fields = resolveMembers("variant "+d.Name+"::"+v.Name, v.Members, ns, logger)
			}

		case *ast.StructDecl:
			s := d.Temp.(*Struct)
			s.Fields = resolveMembers("struct "+d.Name, d.Members, ns, logger)
		default:
			panic(d)
		}
	}
	if logger.NumErrors() > 0 {
		return nil
	}

	return &File{
		Types: types,
	}
}
