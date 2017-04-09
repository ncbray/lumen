package resolved

import (
	"github.com/ncbray/lumen/util"
)

type IntrinsicType struct {
	Temp interface{}
	Name string
}

func (n *IntrinsicType) isType() {
}

type Type interface {
	isType()
}

type Field struct {
	Temp interface{}
	Name string
	Type Type
}

type Format struct {
	Temp   interface{}
	Fields []*Field
}

type Local struct {
	Temp interface{}
	Name string
	Type Type
}

type GetInput struct {
	Temp  interface{}
	Input *Field
}

func (n *GetInput) isExpr() {
}

type GetLocal struct {
	Temp  interface{}
	Local *Local
}

func (n *GetLocal) isExpr() {
}

type Number struct {
	Temp interface{}
	Raw  string
}

func (n *Number) isExpr() {
}

type Prefix struct {
	Temp  interface{}
	Op    string
	Value Expr
}

func (n *Prefix) isExpr() {
}

type Infix struct {
	Temp  interface{}
	Left  Expr
	Op    string
	Right Expr
}

func (n *Infix) isExpr() {
}

type Constructor struct {
	Temp interface{}
	Type Type
	Args []Expr
}

func (n *Constructor) isExpr() {
}

type CallIntrinsic struct {
	Temp interface{}
	Name string
	Args []Expr
}

func (n *CallIntrinsic) isExpr() {
}

type Expr interface {
	isExpr()
}

type ExprValue struct {
	Temp interface{}
	Loc  util.Location
	Expr Expr
	Type Type
}

func (n *ExprValue) isTreeValue() {
}

type TypeValue struct {
	Temp interface{}
	Loc  util.Location
	Type Type
}

func (n *TypeValue) isTreeValue() {
}

type FunctionValue struct {
	Temp interface{}
	Loc  util.Location
	Name string
}

func (n *FunctionValue) isTreeValue() {
}

type TreeValue interface {
	isTreeValue()
}

type SetOutput struct {
	Temp   interface{}
	Output *Field
	Value  Expr
}

func (n *SetOutput) isStatement() {
}

type SetLocal struct {
	Temp  interface{}
	Local *Local
	Value Expr
}

func (n *SetLocal) isStatement() {
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

type Function struct {
	Temp   interface{}
	Name   string
	Locals []*Local
	Body   []Statement
}

type ShaderProgram struct {
	Temp      interface{}
	Name      string
	Uniform   *Format
	Attribute *Format
	Varying   *Format
	Fs        *Function
	Vs        *Function
}

type File struct {
	Temp     interface{}
	Programs []*ShaderProgram
}
