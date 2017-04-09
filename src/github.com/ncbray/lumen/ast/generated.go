package ast

import (
	"github.com/ncbray/lumen/parser"
	"github.com/ncbray/lumen/util"
)

type GetName struct {
	Temp interface{}
	Loc  util.Location
	Name string
}

func (n *GetName) isExpr() {
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
	T     string
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
	Loc   util.Location
	Value Expr
}

func (n *Discard) isStatement() {
}

type Statement interface {
	isStatement()
}

type ShaderDecl struct {
	Temp interface{}
	Loc  util.Location
	Name string
	Vs   []Statement
	Fs   []Statement
}

type File struct {
	Temp    interface{}
	Loc     util.Location
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
			Vs:   c.ConvertStatementList(ctx.GetVs()),
			Fs:   c.ConvertStatementList(ctx.GetFs()),
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
