package ast

import (
	"github.com/ncbray/reboot/util"
)

type TypeRef interface {
	isTypeRef()
}

type TypeName struct {
	Loc  util.Location
	Name string
}

func (n *TypeName) isTypeRef() {
}

type ListRef struct {
	Loc     util.Location
	Element TypeRef
}

func (n *ListRef) isTypeRef() {
}

type MemberDecl interface {
	isMemberDecl()
}

type FieldDecl struct {
	Loc  util.Location
	Name string
	T    TypeRef
}

func (n *FieldDecl) isMemberDecl() {
}

type VariantDecl struct {
	Loc     util.Location
	Name    string
	Members []MemberDecl
}

type Declaration interface {
	isDeclaration()
}

type EnumDecl struct {
	Loc      util.Location
	Name     string
	Variants []*VariantDecl
}

func (n *EnumDecl) isDeclaration() {
}

type StructDecl struct {
	Loc     util.Location
	Name    string
	Members []MemberDecl
}

func (n *StructDecl) isDeclaration() {
}

type File struct {
	Loc   util.Location
	Decls []Declaration
}
