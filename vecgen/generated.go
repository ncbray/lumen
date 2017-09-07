package vecgen

type VectorType struct {
	Temp  interface{}
	Width int
	Name  string
}

func (n *VectorType) isType() {
}

type MatrixType struct {
	Temp  interface{}
	Width int
	Name  string
}

func (n *MatrixType) isType() {
}

type Type interface {
	isType()
}

type Object struct {
	Temp         interface{}
	Type         Type
	Name         string
	CachedGetter []string
	CachedSetter []string
}

type Atom struct {
	Temp interface{}
	Text string
}

func (n *Atom) isExpr() {
}

type Neg struct {
	Temp interface{}
	Expr Expr
}

func (n *Neg) isExpr() {
}

type Add struct {
	Temp  interface{}
	Left  Expr
	Right Expr
}

func (n *Add) isExpr() {
}

type Sub struct {
	Temp  interface{}
	Left  Expr
	Right Expr
}

func (n *Sub) isExpr() {
}

type Mul struct {
	Temp  interface{}
	Left  Expr
	Right Expr
}

func (n *Mul) isExpr() {
}

type Div struct {
	Temp  interface{}
	Left  Expr
	Right Expr
}

func (n *Div) isExpr() {
}

type GetComponent struct {
	Temp   interface{}
	Object *Object
	Index  int
}

func (n *GetComponent) isExpr() {
}

type Expr interface {
	isExpr()
}

type SetComponent struct {
	Temp   interface{}
	Expr   Expr
	Object *Object
	Index  int
}
