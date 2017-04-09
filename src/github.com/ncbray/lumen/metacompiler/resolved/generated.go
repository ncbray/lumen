package resolved

type Field struct {
	Temp interface{}
	Name string
	Type Type
}

type IntrinsicType struct {
	Temp interface{}
	Name string
	List *List
}

func (n *IntrinsicType) isType() {
}

type Struct struct {
	Temp    interface{}
	Name    string
	Fields  []*Field
	Holders []*Holder
	List    *List
}

func (n *Struct) isType() {
}

type Holder struct {
	Temp  interface{}
	Name  string
	Types []*Struct
	List  *List
}

func (n *Holder) isType() {
}

type List struct {
	Temp    interface{}
	Element Type
	List    *List
}

func (n *List) isType() {
}

type Type interface {
	isType()
}

type FieldArg struct {
	Temp  interface{}
	Field *Field
	Value ParserBindingExpr
}

type Construct struct {
	Temp interface{}
	Type *Struct
	Args []*FieldArg
}

func (n *Construct) isParserBindingExpr() {
}

type GetParam struct {
	Temp  interface{}
	Param *ParserBindingParam
}

func (n *GetParam) isParserBindingExpr() {
}

type GetLocStart struct {
	Temp interface{}
}

func (n *GetLocStart) isParserBindingExpr() {
}

type ParserBindingExpr interface {
	isParserBindingExpr()
}

type InputList struct {
	Temp    interface{}
	Element *ParserBindingGroup
}

func (n *InputList) isParserBindingInput() {
}

type InputSlice struct {
	Temp interface{}
}

func (n *InputSlice) isParserBindingInput() {
}

type ParserBindingInput interface {
	isParserBindingInput()
}

type ParserBindingParam struct {
	Temp  interface{}
	Name  string
	Input ParserBindingInput
}

type ParserBindingMapping struct {
	Temp   interface{}
	Rule   string
	Params []*ParserBindingParam
	Body   ParserBindingExpr
}

type ParserBindingGroup struct {
	Temp     interface{}
	Name     string
	Type     Type
	Mappings []*ParserBindingMapping
	List     *InputList
}

func (n *ParserBindingGroup) isParserBindingInput() {
}

type ParserBinding struct {
	Temp   interface{}
	Groups []*ParserBindingGroup
}

type File struct {
	Temp          interface{}
	Types         []Type
	ParserBinding *ParserBinding
}
