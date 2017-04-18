package resolved

import (
	"fmt"

	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/metacompiler/ast"
	"github.com/ncbray/lumen/util"
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
	case *IntrinsicType:
		if t.List == nil {
			t.List = &List{Element: t}
		}
		return t.List
	case *Struct:
		if t.List == nil {
			t.List = &List{Element: t}
		}
		return t.List
	case *Holder:
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

func canonicalTypeName(t Type) string {
	if t == nil {
		return "<nil>"
	}
	switch t := t.(type) {
	case *IntrinsicType:
		return t.Name
	case *Struct:
		return t.Name
	case *Holder:
		return t.Name
	case *List:
		return "[]" + canonicalTypeName(t.Element)
	default:
		panic(t)
	}
}

func typeRefLoc(t ast.TypeRef) util.Location {
	switch t := t.(type) {
	case *ast.TypeName:
		return t.Loc
	case *ast.ListRef:
		return t.Loc
	default:
		panic(t)
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

func resolveParserInput(t ast.TypeRef, groupLUT map[string]*ParserBindingGroup, logger log.CompilerLogger) ParserBindingInput {
	switch t := t.(type) {
	case *ast.TypeName:
		if t.Name == "string" {
			return &InputSlice{}
		}
		gref, ok := groupLUT[t.Name]
		if !ok {
			loc := t.Loc
			logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Unknown rule %q.", t.Name))
			return nil
		}
		return gref
	case *ast.ListRef:
		switch t := t.Element.(type) {
		case *ast.TypeName:
			gref, ok := groupLUT[t.Name]
			if !ok {
				loc := t.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Unknown rule %q.", t.Name))
				return nil
			}
			if gref.List == nil {
				gref.List = &InputList{
					Element: gref,
				}
			}
			return gref.List
		default:
			panic(t)
		}
	default:
		panic(t)
	}
}

func checkCanHold(dst Type, src Type, loc util.Location, logger log.CompilerLogger) {
	var ok bool
	switch dst := dst.(type) {
	case *IntrinsicType, *List, *Struct:
		ok = dst == src
	case *Holder:
		if dst == src {
			ok = src == dst
			break
		}
		for _, t := range dst.Types {
			if t == src {
				ok = true
				break
			}
		}
	default:
		panic(dst)
	}
	if !ok {
		logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Expected %s, got %s", canonicalTypeName(dst), canonicalTypeName(src)))
	}
}

func resolveBindingExpr(e ast.ParserBindingExpr, paramLUT map[string]*ParserBindingParam, ns *namespace, builtins *namespace, logger log.CompilerLogger) (ParserBindingExpr, Type) {
	switch e := e.(type) {
	case *ast.NameRef:
		if e.Name == "loc_start" {
			return &GetLocStart{}, builtins.ResolveType("location")
		}
		p, ok := paramLUT[e.Name]
		if ok {
			var it Type
			switch i := p.Input.(type) {
			case *InputSlice:
				it = builtins.ResolveType("string")
			case *InputList:
				it = listOfType(i.Element.Type)
			case *ParserBindingGroup:
				it = i.Type
			default:
				panic(i)
			}
			return &GetParam{
				Param: p,
			}, it
		}
		panic(e.Name)
	case *ast.Construct:
		t := ns.ResolveType(e.Name)
		if t == nil {
			loc := e.Loc
			logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Unknown type %q.", e.Name))
			return nil, nil
		}
		st, ok := t.(*Struct)
		if !ok {
			loc := e.Loc
			logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "Can only construct structures.")
			return nil, nil
		}

		fieldLUT := map[string]*Field{}
		for _, f := range st.Fields {
			fieldLUT[f.Name] = f
		}

		args := []*FieldArg{}
		for _, arg := range e.Args {
			value, vt := resolveBindingExpr(arg.Value, paramLUT, ns, builtins, logger)
			f, ok := fieldLUT[arg.Name]
			if !ok {
				loc := e.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Unknown field %s.%s", st.Name, arg.Name))
				continue
			}
			checkCanHold(f.Type, vt, arg.Loc, logger)
			args = append(args, &FieldArg{
				Field: f,
				Value: value,
			})
		}
		return &Construct{
			Type: st,
			Args: args,
		}, st
	default:
		panic(e)
	}
}

func FromAST(src *ast.File, logger log.CompilerLogger) *File {
	builtins := &namespace{
		parent: nil,
		types: map[string]Type{
			"string":   &IntrinsicType{Name: "string"},
			"int":      &IntrinsicType{Name: "int"},
			"bool":     &IntrinsicType{Name: "bool"},
			"location": &IntrinsicType{Name: "location"},
		},
	}
	ns := &namespace{
		parent: builtins,
		types:  map[string]Type{},
	}
	// Generate types.
	types := []Type{}
	var parserBinding *ParserBinding
	for _, d := range src.Decls {
		switch d := d.(type) {
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
		case *ast.HolderDecl:
			if ns.IsDefined(d.Name) {
				loc := d.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Attempted to redefine %q.", d.Name))
				continue
			}
			t := &Holder{
				Name: d.Name,
			}
			ns.Define(d.Name, t)
			types = append(types, t)
			d.Temp = t
		case *ast.ParserBindingDecl:
			if parserBinding != nil {
				loc := d.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "Only one parser binding allowed.")
				continue
			}
			b := &ParserBinding{}
			d.Temp = b
			parserBinding = b
			// TODO
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
		case *ast.StructDecl:
			s := d.Temp.(*Struct)
			s.Fields = resolveMembers("struct "+d.Name, d.Members, ns, logger)
		case *ast.HolderDecl:
			h := d.Temp.(*Holder)
			for _, ref := range d.Types {
				t := resolveType(ref, ns, logger)
				switch t := t.(type) {
				case *Struct:
					h.Types = append(h.Types, t)
					t.Holders = append(t.Holders, h)
				default:
					loc := typeRefLoc(ref)
					logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "Type must be a struct.")
				}
			}
		case *ast.ParserBindingDecl:
			b := d.Temp.(*ParserBinding)
			groups := []*ParserBindingGroup{}
			groupLUT := map[string]*ParserBindingGroup{}
			for _, g := range d.Groups {
				_, ok := groupLUT[g.Name]
				if ok {
					loc := g.Loc
					logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("Attempted to redefine %q.", g.Name))
					continue
				}
				rg := &ParserBindingGroup{
					Name: g.Name,
					Type: resolveType(g.T, ns, logger),
				}
				g.Temp = rg
				groups = append(groups, rg)
			}
			b.Groups = groups
		default:
			panic(d)
		}
	}
	if logger.NumErrors() > 0 {
		return nil
	}

	// Resolve function bodies.
	for _, d := range src.Decls {
		switch d := d.(type) {
		case *ast.StructDecl, *ast.HolderDecl:
		case *ast.ParserBindingDecl:
			//b := d.Temp.(*ParserBinding)

			// Index
			groupLUT := map[string]*ParserBindingGroup{}
			for _, g := range d.Groups {
				rg := g.Temp.(*ParserBindingGroup)
				groupLUT[g.Name] = rg
			}

			// Process
			for _, g := range d.Groups {
				rg := g.Temp.(*ParserBindingGroup)
				mappings := []*ParserBindingMapping{}
				for _, m := range g.Mappings {
					params := []*ParserBindingParam{}
					paramLUT := map[string]*ParserBindingParam{}
					for _, p := range m.Params {
						rp := &ParserBindingParam{
							Name:  p.Name,
							Input: resolveParserInput(p.T, groupLUT, logger),
						}
						params = append(params, rp)
						paramLUT[p.Name] = rp
					}
					// TODO check type
					e, et := resolveBindingExpr(m.Body, paramLUT, ns, builtins, logger)
					checkCanHold(rg.Type, et, m.Loc, logger)
					rm := &ParserBindingMapping{
						Rule:   m.Name,
						Params: params,
						Body:   e,
					}
					mappings = append(mappings, rm)
				}
				rg.Mappings = mappings
			}
		default:
			panic(d)
		}
	}
	if logger.NumErrors() > 0 {
		return nil
	}

	return &File{
		Types:         types,
		ParserBinding: parserBinding,
	}
}
