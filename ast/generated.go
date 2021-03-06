package ast

import (
	"github.com/ncbray/lumen/parser"
	"github.com/ncbray/lumen/util"
)

type Field struct {
	Temp interface{}
	Loc  util.Location
	Name string
	Type string
}

type Format struct {
	Temp   interface{}
	Loc    util.Location
	Fields []*Field
}

type GetName struct {
	Temp interface{}
	Loc  util.Location
	Name string
}

func (n *GetName) isExpr() {
}

type GetAttr struct {
	Temp  interface{}
	Loc   util.Location
	Value Expr
	Name  string
}

func (n *GetAttr) isExpr() {
}

type Number struct {
	Temp interface{}
	Loc  util.Location
	Raw  string
}

func (n *Number) isExpr() {
}

type Prefix struct {
	Temp  interface{}
	Loc   util.Location
	Op    string
	Value Expr
}

func (n *Prefix) isExpr() {
}

type Infix struct {
	Temp  interface{}
	Loc   util.Location
	Left  Expr
	Op    string
	Right Expr
}

func (n *Infix) isExpr() {
}

type Call struct {
	Temp  interface{}
	Loc   util.Location
	Value Expr
	Args  []Expr
}

func (n *Call) isExpr() {
}

type Expr interface {
	isExpr()
}

type VarDecl struct {
	Temp  interface{}
	Loc   util.Location
	Type  string
	Name  string
	Value Expr
}

func (n *VarDecl) isStatement() {
}

type Assign struct {
	Temp  interface{}
	Loc   util.Location
	Name  string
	Value Expr
}

func (n *Assign) isStatement() {
}

type Discard struct {
	Temp  interface{}
	Value Expr
}

func (n *Discard) isStatement() {
}

type Statement interface {
	isStatement()
}

type ShaderDecl struct {
	Temp      interface{}
	Loc       util.Location
	Name      string
	Uniform   *Format
	Attribute *Format
	Varying   *Format
	Vs        []Statement
	Fs        []Statement
}

func (n *ShaderDecl) isTopLevelDecl() {
}

type VertexComponent struct {
	Temp     interface{}
	Loc      util.Location
	Name     string
	Type     string
	Encoding string
}

type VertexDecl struct {
	Temp       interface{}
	Loc        util.Location
	Name       string
	Components []*VertexComponent
}

func (n *VertexDecl) isTopLevelDecl() {
}

type TopLevelDecl interface {
	isTopLevelDecl()
}

type File struct {
	Temp  interface{}
	Decls []TopLevelDecl
}

type ASTConverter struct {
	Filename string
}

func (c *ASTConverter) ConvertField(ctx parser.IFieldContext) *Field {
	switch ctx := ctx.(type) {
	case *parser.FieldContext:
		return &Field{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
			Type: ctx.GetT().GetText(),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertFieldList(src []parser.IFieldContext) []*Field {
	dst := make([]*Field, len(src))
	for i, child := range src {
		dst[i] = c.ConvertField(child)
	}
	return dst
}

func (c *ASTConverter) ConvertFormat(ctx parser.IFormatContext) *Format {
	switch ctx := ctx.(type) {
	case *parser.FormatContext:
		return &Format{
			Loc:    util.GetLocation(c.Filename, ctx.GetStart()),
			Fields: c.ConvertFieldList(ctx.GetFields()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertExpr(ctx parser.IExprContext) Expr {
	switch ctx := ctx.(type) {
	case *parser.GetNameContext:
		return &GetName{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
		}
	case *parser.GetAttrContext:
		return &GetAttr{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Value: c.ConvertExpr(ctx.GetValue()),
			Name:  ctx.GetName().GetText(),
		}
	case *parser.NumberContext:
		return &Number{
			Loc: util.GetLocation(c.Filename, ctx.GetStart()),
			Raw: ctx.GetRaw().GetText(),
		}
	case *parser.PrefixContext:
		return &Prefix{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Op:    ctx.GetOp().GetText(),
			Value: c.ConvertExpr(ctx.GetValue()),
		}
	case *parser.InfixContext:
		return &Infix{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Left:  c.ConvertExpr(ctx.GetLeft()),
			Op:    ctx.GetOp().GetText(),
			Right: c.ConvertExpr(ctx.GetRight()),
		}
	case *parser.CallContext:
		return &Call{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Value: c.ConvertExpr(ctx.GetValue()),
			Args:  c.ConvertExprList(ctx.GetArgs()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertExprList(src []parser.IExprContext) []Expr {
	dst := make([]Expr, len(src))
	for i, child := range src {
		dst[i] = c.ConvertExpr(child)
	}
	return dst
}

func (c *ASTConverter) ConvertStatement(ctx parser.IStatementContext) Statement {
	switch ctx := ctx.(type) {
	case *parser.VarDeclContext:
		return &VarDecl{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Type:  ctx.GetT().GetText(),
			Name:  ctx.GetName().GetText(),
			Value: c.ConvertExpr(ctx.GetValue()),
		}
	case *parser.AssignContext:
		return &Assign{
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
			Name:  ctx.GetName().GetText(),
			Value: c.ConvertExpr(ctx.GetValue()),
		}
	case *parser.DiscardContext:
		return &Discard{
			Value: c.ConvertExpr(ctx.GetValue()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertStatementList(src []parser.IStatementContext) []Statement {
	dst := make([]Statement, len(src))
	for i, child := range src {
		dst[i] = c.ConvertStatement(child)
	}
	return dst
}

func (c *ASTConverter) ConvertVertexComponent(ctx parser.IVertexComponentContext) *VertexComponent {
	switch ctx := ctx.(type) {
	case *parser.VertexComponentContext:
		return &VertexComponent{
			Loc:      util.GetLocation(c.Filename, ctx.GetStart()),
			Name:     ctx.GetName().GetText(),
			Type:     ctx.GetT().GetText(),
			Encoding: ctx.GetEncoding().GetText(),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertVertexComponentList(src []parser.IVertexComponentContext) []*VertexComponent {
	dst := make([]*VertexComponent, len(src))
	for i, child := range src {
		dst[i] = c.ConvertVertexComponent(child)
	}
	return dst
}

func (c *ASTConverter) ConvertTopLevelDecl(ctx parser.ITopLevelDeclContext) TopLevelDecl {
	switch ctx := ctx.(type) {
	case *parser.ShaderDeclContext:
		return &ShaderDecl{
			Loc:       util.GetLocation(c.Filename, ctx.GetStart()),
			Name:      ctx.GetName().GetText(),
			Uniform:   c.ConvertFormat(ctx.GetUniform()),
			Attribute: c.ConvertFormat(ctx.GetAttribute()),
			Varying:   c.ConvertFormat(ctx.GetVarying()),
			Vs:        c.ConvertStatementList(ctx.GetVs()),
			Fs:        c.ConvertStatementList(ctx.GetFs()),
		}
	case *parser.VertexDeclContext:
		return &VertexDecl{
			Loc:        util.GetLocation(c.Filename, ctx.GetStart()),
			Name:       ctx.GetName().GetText(),
			Components: c.ConvertVertexComponentList(ctx.GetComponents()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertTopLevelDeclList(src []parser.ITopLevelDeclContext) []TopLevelDecl {
	dst := make([]TopLevelDecl, len(src))
	for i, child := range src {
		dst[i] = c.ConvertTopLevelDecl(child)
	}
	return dst
}

func (c *ASTConverter) ConvertFile(ctx parser.IFileContext) *File {
	switch ctx := ctx.(type) {
	case *parser.FileContext:
		return &File{
			Decls: c.ConvertTopLevelDeclList(ctx.GetDecls()),
		}
	default:
		panic(ctx)
	}
}
