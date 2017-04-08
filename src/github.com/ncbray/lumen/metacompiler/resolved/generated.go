package resolved

type Field struct {
	Name string
	Type Type
}

type Variant struct {
	Name   string
	Fields []*Field
}

type Intrinsic struct {
	Name string
	List Type
}

func (n *Intrinsic) isType() {
}

type Enum struct {
	Name     string
	Variants []*Variant
	List     Type
}

func (n *Enum) isType() {
}

type Struct struct {
	Name    string
	Fields  []*Field
	Holders []*Holder
	List    Type
}

func (n *Struct) isType() {
}

type Holder struct {
	Name  string
	Types []*Struct
	List  Type
}

func (n *Holder) isType() {
}

type List struct {
	Element Type
	List    Type
}

func (n *List) isType() {
}

type KeywordArg struct {
	Name  string
	Value ParserBindingExpr
}

type Construct struct {
	Type string
	Args []*KeywordArg
}

func (n *Construct) isParserBindingExpr() {
}

type GetParam struct {
	Param *ParserBindingParam
}

func (n *GetParam) isParserBindingExpr() {
}

type GetLocStart struct {
}

func (n *GetLocStart) isParserBindingExpr() {
}

type ParserBindingExpr interface {
	isParserBindingExpr()
}

type InputList struct {
	Element *ParserBindingGroup
}

func (n *InputList) isParserBindingInput() {
}

type InputSlice struct {
}

func (n *InputSlice) isParserBindingInput() {
}

type ParserBindingInput interface {
	isParserBindingInput()
}

type ParserBindingParam struct {
	Name  string
	Input ParserBindingInput
}

type ParserBindingMapping struct {
	Rule   string
	Params []*ParserBindingParam
	Body   ParserBindingExpr
}

type ParserBindingGroup struct {
	Name     string
	Type     Type
	Mappings []*ParserBindingMapping
	List     *InputList
}

func (n *ParserBindingGroup) isParserBindingInput() {
}

type ParserBinding struct {
	Groups []*ParserBindingGroup
}

type Type interface {
	isType()
}

type File struct {
	Types         []Type
	ParserBinding *ParserBinding
}
