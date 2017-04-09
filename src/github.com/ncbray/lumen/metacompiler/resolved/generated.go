package resolved

type Field struct {
	Name string
	Type Type
}

type IntrinsicType struct {
	Name string
	List *List
}

func (n *IntrinsicType) isType() {
}

type Struct struct {
	Name    string
	Fields  []*Field
	Holders []*Holder
	List    *List
}

func (n *Struct) isType() {
}

type Holder struct {
	Name  string
	Types []*Struct
	List  *List
}

func (n *Holder) isType() {
}

type List struct {
	Element Type
	List    *List
}

func (n *List) isType() {
}

type Type interface {
	isType()
}

type KeywordArg struct {
	Name  string
	Value ParserBindingExpr
}

type Construct struct {
	Type *Struct
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

type File struct {
	Types         []Type
	ParserBinding *ParserBinding
}
