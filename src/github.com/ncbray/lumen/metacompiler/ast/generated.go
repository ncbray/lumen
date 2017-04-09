package ast

import (
	"github.com/ncbray/lumen/metacompiler/parser"
	"github.com/ncbray/lumen/util"
)

type TypeName struct {
	Loc  util.Location
	Temp interface{}
	Name string
}

func (n *TypeName) isTypeRef() {
}

type ListRef struct {
	Loc     util.Location
	Temp    interface{}
	Element TypeRef
}

func (n *ListRef) isTypeRef() {
}

type TypeRef interface {
	isTypeRef()
}

type FieldDecl struct {
	Loc  util.Location
	Temp interface{}
	Name string
	T    TypeRef
}

func (n *FieldDecl) isMemberDecl() {
}

type MemberDecl interface {
	isMemberDecl()
}

type KeywordArg struct {
	Loc   util.Location
	Temp  interface{}
	Name  string
	Value ParserBindingExpr
}

type Construct struct {
	Loc  util.Location
	Temp interface{}
	Name string
	Args []*KeywordArg
}

func (n *Construct) isParserBindingExpr() {
}

type NameRef struct {
	Loc  util.Location
	Temp interface{}
	Name string
}

func (n *NameRef) isParserBindingExpr() {
}

type ParserBindingExpr interface {
	isParserBindingExpr()
}

type ParamDecl struct {
	Loc  util.Location
	Temp interface{}
	Name string
	T    TypeRef
}

type ParserBindingMapping struct {
	Loc    util.Location
	Temp   interface{}
	Name   string
	Params []*ParamDecl
	Body   ParserBindingExpr
}

type ParserBindingGroup struct {
	Loc      util.Location
	Temp     interface{}
	Name     string
	T        TypeRef
	Mappings []*ParserBindingMapping
}

type StructDecl struct {
	Loc     util.Location
	Temp    interface{}
	Name    string
	Members []MemberDecl
}

func (n *StructDecl) isDeclaration() {
}

type HolderDecl struct {
	Loc   util.Location
	Temp  interface{}
	Name  string
	Types []TypeRef
}

func (n *HolderDecl) isDeclaration() {
}

type ParserBindingDecl struct {
	Loc    util.Location
	Temp   interface{}
	Groups []*ParserBindingGroup
}

func (n *ParserBindingDecl) isDeclaration() {
}

type Declaration interface {
	isDeclaration()
}

type File struct {
	Loc   util.Location
	Temp  interface{}
	Decls []Declaration
}

type ASTConverter struct {
	Filename string
}

func (c *ASTConverter) ConvertTypeRef(ctx parser.ITypeRefContext) TypeRef {
	switch ctx := ctx.(type) {
	case *parser.TypeNameContext:
		return &TypeName{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
		}
	case *parser.ListRefContext:
		return &ListRef{
			Loc:     util.GetLocation(c.Filename, ctx.GetStart()),
			Element: c.ConvertTypeRef(ctx.GetElement()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertTypeRefList(src []parser.ITypeRefContext) []TypeRef {
	dst := make([]TypeRef, len(src))
	for i, child := range src {
		dst[i] = c.ConvertTypeRef(child)
	}
	return dst
}

func (c *ASTConverter) ConvertMemberDecl(ctx parser.IMemberDeclContext) MemberDecl {
	switch ctx := ctx.(type) {
	case *parser.FieldDeclContext:
		return &FieldDecl{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
			T:    c.ConvertTypeRef(ctx.GetT()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertMemberDeclList(src []parser.IMemberDeclContext) []MemberDecl {
	dst := make([]MemberDecl, len(src))
	for i, child := range src {
		dst[i] = c.ConvertMemberDecl(child)
	}
	return dst
}

func (c *ASTConverter) ConvertKeywordArg(ctx parser.IKeywordArgContext) *KeywordArg {
	switch ctx := ctx.(type) {
	case *parser.KeywordArgContext:
		return &KeywordArg{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Name:  ctx.GetName().GetText(),
			Value: c.ConvertParserBindingExpr(ctx.GetValue()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertKeywordArgList(src []parser.IKeywordArgContext) []*KeywordArg {
	dst := make([]*KeywordArg, len(src))
	for i, child := range src {
		dst[i] = c.ConvertKeywordArg(child)
	}
	return dst
}

func (c *ASTConverter) ConvertParserBindingExpr(ctx parser.IParserBindingExprContext) ParserBindingExpr {
	switch ctx := ctx.(type) {
	case *parser.ConstructContext:
		return &Construct{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
			Args: c.ConvertKeywordArgList(ctx.GetArgs()),
		}
	case *parser.NameRefContext:
		return &NameRef{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertParamDecl(ctx parser.IParamDeclContext) *ParamDecl {
	switch ctx := ctx.(type) {
	case *parser.ParamDeclContext:
		return &ParamDecl{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
			T:    c.ConvertTypeRef(ctx.GetT()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertParamDeclList(src []parser.IParamDeclContext) []*ParamDecl {
	dst := make([]*ParamDecl, len(src))
	for i, child := range src {
		dst[i] = c.ConvertParamDecl(child)
	}
	return dst
}

func (c *ASTConverter) ConvertParserBindingMapping(ctx parser.IParserBindingMappingContext) *ParserBindingMapping {
	switch ctx := ctx.(type) {
	case *parser.ParserBindingMappingContext:
		return &ParserBindingMapping{
			Loc:    util.GetLocation(c.Filename, ctx.GetStart()),
			Name:   ctx.GetName().GetText(),
			Params: c.ConvertParamDeclList(ctx.GetParams()),
			Body:   c.ConvertParserBindingExpr(ctx.GetBody()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertParserBindingMappingList(src []parser.IParserBindingMappingContext) []*ParserBindingMapping {
	dst := make([]*ParserBindingMapping, len(src))
	for i, child := range src {
		dst[i] = c.ConvertParserBindingMapping(child)
	}
	return dst
}

func (c *ASTConverter) ConvertParserBindingGroup(ctx parser.IParserBindingGroupContext) *ParserBindingGroup {
	switch ctx := ctx.(type) {
	case *parser.ParserBindingGroupContext:
		return &ParserBindingGroup{
			Loc:      util.GetLocation(c.Filename, ctx.GetStart()),
			Name:     ctx.GetName().GetText(),
			T:        c.ConvertTypeRef(ctx.GetT()),
			Mappings: c.ConvertParserBindingMappingList(ctx.GetMappings()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertParserBindingGroupList(src []parser.IParserBindingGroupContext) []*ParserBindingGroup {
	dst := make([]*ParserBindingGroup, len(src))
	for i, child := range src {
		dst[i] = c.ConvertParserBindingGroup(child)
	}
	return dst
}

func (c *ASTConverter) ConvertDeclaration(ctx parser.IDeclarationContext) Declaration {
	switch ctx := ctx.(type) {
	case *parser.StructDeclContext:
		return &StructDecl{
			Loc:     util.GetLocation(c.Filename, ctx.GetStart()),
			Name:    ctx.GetName().GetText(),
			Members: c.ConvertMemberDeclList(ctx.GetMembers()),
		}
	case *parser.HolderDeclContext:
		return &HolderDecl{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Name:  ctx.GetName().GetText(),
			Types: c.ConvertTypeRefList(ctx.GetTypes()),
		}
	case *parser.ParserBindingDeclContext:
		return &ParserBindingDecl{
			Loc:    util.GetLocation(c.Filename, ctx.GetStart()),
			Groups: c.ConvertParserBindingGroupList(ctx.GetGroups()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertDeclarationList(src []parser.IDeclarationContext) []Declaration {
	dst := make([]Declaration, len(src))
	for i, child := range src {
		dst[i] = c.ConvertDeclaration(child)
	}
	return dst
}

func (c *ASTConverter) ConvertFile(ctx parser.IFileContext) *File {
	switch ctx := ctx.(type) {
	case *parser.FileContext:
		return &File{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Decls: c.ConvertDeclarationList(ctx.GetDecls()),
		}
	default:
		panic(ctx)
	}
}
