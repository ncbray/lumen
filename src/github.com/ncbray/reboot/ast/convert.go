package ast

import (
	"github.com/ncbray/reboot/parser"
	"github.com/ncbray/reboot/util"
)

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

func (c *ASTConverter) ConvertVariantDecl(ctx parser.IVariantDeclContext) *VariantDecl {
	switch ctx := ctx.(type) {
	case *parser.VariantDeclContext:
		return &VariantDecl{
			Loc:     util.GetLocation(c.Filename, ctx.GetStart()),
			Name:    ctx.GetName().GetText(),
			Members: c.ConvertMemberDeclList(ctx.GetMembers()),
		}
	default:
		panic(ctx)
	}
}

func (c *ASTConverter) ConvertVariantDeclList(src []parser.IVariantDeclContext) []*VariantDecl {
	dst := make([]*VariantDecl, len(src))
	for i, child := range src {
		dst[i] = c.ConvertVariantDecl(child)
	}
	return dst
}

func (c *ASTConverter) ConvertDeclaration(ctx parser.IDeclarationContext) Declaration {
	switch ctx := ctx.(type) {
	case *parser.EnumDeclContext:
		return &EnumDecl{
			Loc:      util.GetLocation(c.Filename, ctx.GetStart()),
			Name:     ctx.GetName().GetText(),
			Variants: c.ConvertVariantDeclList(ctx.GetVariants()),
		}
	case *parser.StructDeclContext:
		return &StructDecl{
			Loc:     util.GetLocation(c.Filename, ctx.GetStart()),
			Name:    ctx.GetName().GetText(),
			Members: c.ConvertMemberDeclList(ctx.GetMembers()),
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

func (c *ASTConverter) ConvertFileList(src []parser.IFileContext) []*File {
	dst := make([]*File, len(src))
	for i, child := range src {
		dst[i] = c.ConvertFile(child)
	}
	return dst
}