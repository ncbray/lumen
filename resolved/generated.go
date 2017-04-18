package resolved

type ScalarType struct {
	Temp interface{}
	Name string
}

func (n *ScalarType) isType() {
}

type VectorType struct {
	Temp   interface{}
	Name   string
	Width  int
	Scalar *ScalarType
}

func (n *VectorType) isType() {
}

type MatrixType struct {
	Temp   interface{}
	Name   string
	Width  int
	Height int
	Scalar *ScalarType
}

func (n *MatrixType) isType() {
}

type Type interface {
	isType()
}

type Field struct {
	Temp   interface{}
	Name   string
	Type   Type
	Marked bool
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

type GetAttr struct {
	Temp  interface{}
	Value Expr
	Name  string
}

func (n *GetAttr) isExpr() {
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

type FunctionSignature struct {
	Temp    interface{}
	Params  []Type
	Returns []Type
}

type IntrinsicFunction struct {
	Temp      interface{}
	Name      string
	Signature *FunctionSignature
}

type CallIntrinsic struct {
	Temp     interface{}
	Function *IntrinsicFunction
	Args     []Expr
}

func (n *CallIntrinsic) isExpr() {
}

type Expr interface {
	isExpr()
}

type ExprValue struct {
	Temp interface{}
	Expr Expr
	Type []Type
}

func (n *ExprValue) isTreeValue() {
}

type TypeValue struct {
	Temp interface{}
	Type Type
}

func (n *TypeValue) isTreeValue() {
}

type FunctionValue struct {
	Temp     interface{}
	Function *IntrinsicFunction
}

func (n *FunctionValue) isTreeValue() {
}

type Poison struct {
	Temp interface{}
}

func (n *Poison) isTreeValue() {
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

type VertexComponent struct {
	Temp       interface{}
	Name       string
	Type       Type
	Encoding   string
	ByteOffset int
	ByteSize   int
}

type VertexFormat struct {
	Temp       interface{}
	Name       string
	Components []*VertexComponent
	ByteSize   int
}

type File struct {
	Temp     interface{}
	Vertex   []*VertexFormat
	Programs []*ShaderProgram
}
