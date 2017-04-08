package ast

import (
	"github.com/ncbray/lumen/parser"
	"github.com/ncbray/lumen/util"
)

type Expr interface {
	isExpr()
}

type GetName struct {
	Loc  util.Location
	Temp interface{}
	Name string
}

func (n *GetName) isExpr() {
}

type Number struct {
	Loc  util.Location
	Temp interface{}
	Raw  string
}

func (n *Number) isExpr() {
}

type Prefix struct {
	Loc   util.Location
	Temp  interface{}
	Op    string
	Value Expr
}

func (n *Prefix) isExpr() {
}

type Infix struct {
	Loc   util.Location
	Temp  interface{}
	Left  Expr
	Op    string
	Right Expr
}

func (n *Infix) isExpr() {
}

type Call struct {
	Loc   util.Location
	Temp  interface{}
	Value Expr
	Args  []Expr
}

func (n *Call) isExpr() {
}

type Statement interface {
	isStatement()
}

type VarDecl struct {
	Loc   util.Location
	Temp  interface{}
	T     string
	Name  string
	Value Expr
}

func (n *VarDecl) isStatement() {
}

type Assign struct {
	Loc   util.Location
	Temp  interface{}
	Name  string
	Value Expr
}

func (n *Assign) isStatement() {
}

type Discard struct {
	Loc   util.Location
	Temp  interface{}
	Value Expr
}

func (n *Discard) isStatement() {
}

type ShaderDecl struct {
	Loc  util.Location
	Temp interface{}
	Name string
	Fs   []Statement
	Vs   []Statement
}

type File struct {
	Loc     util.Location
	Temp    interface{}
	Shaders []*ShaderDecl
}

type ASTConverter struct {
	Filename string
}

func (c *ASTConverter) ConvertExpr(ctx parser.IExprContext) Expr {
	switch ctx := ctx.(type) {
	case *parser.GetNameContext:
		return &GetName{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
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
			T:     ctx.GetT().GetText(),
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
			Loc:   util.GetLocation(c.Filename, ctx.GetStart()),
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

func (c *ASTConverter) ConvertShaderDecl(ctx parser.IShaderDeclContext) *ShaderDecl {
	switch ctx := ctx.(type) {
	case *parser.ShaderDeclContext:
		return &ShaderDecl{
			Loc:  util.GetLocation(c.Filename, ctx.GetStart()),
			Name: ctx.GetName().GetText(),
			Fs:   c.ConvertStatementList(ctx.GetFs()),
			Vs:   c.ConvertStatementList(ctx.GetVs()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertShaderDeclList(src []parser.IShaderDeclContext) []*ShaderDecl {
	dst := make([]*ShaderDecl, len(src))
	for i, child := range src {
		dst[i] = c.ConvertShaderDecl(child)
	}
	return dst
}

func (c *ASTConverter) ConvertFile(ctx parser.IFileContext) *File {
	switch ctx := ctx.(type) {
	case *parser.FileContext:
		return &File{
			Loc:     util.GetLocation(c.Filename, ctx.GetStart()),
			Shaders: c.ConvertShaderDeclList(ctx.GetShaders()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertFileList(src []parser.IFileContext) []*File {
	dst := make([]*File, len(src))
	for i, child := range src {
		dst[i] = c.ConvertFile(child)
	}
	return dst
}
