package resolved

type Intrinsic struct {
	Name string
}

func (n *Intrinsic) isType() {
}

type Type interface {
	isType()
}

type Local struct {
	Name string
	Type Type
}

type GetInput struct {
	Name string
}

func (n *GetInput) isExpr() {
}

type GetLocal struct {
	Local *Local
}

func (n *GetLocal) isExpr() {
}

type Number struct {
	Raw string
}

func (n *Number) isExpr() {
}

type Prefix struct {
	Op    string
	Value Expr
}

func (n *Prefix) isExpr() {
}

type Infix struct {
	Left  Expr
	Op    string
	Right Expr
}

func (n *Infix) isExpr() {
}

type Constructor struct {
	Type Type
	Args []Expr
}

func (n *Constructor) isExpr() {
}

type CallIntrinsic struct {
	Name string
	Args []Expr
}

func (n *CallIntrinsic) isExpr() {
}

type Expr interface {
	isExpr()
}

type SetOutput struct {
	Name  string
	Value Expr
}

func (n *SetOutput) isStatement() {
}

type SetLocal struct {
	Local *Local
	Value Expr
}

func (n *SetLocal) isStatement() {
}

type Discard struct {
	Value Expr
}

func (n *Discard) isStatement() {
}

type Statement interface {
	isStatement()
}

type Function struct {
	Name   string
	Locals []*Local
	Body   []Statement
}

type ShaderProgram struct {
	Name string
	Fs   *Function
	Vs   *Function
}

type File struct {
	Programs []*ShaderProgram
}
