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

type Type interface {
	isType()
}

type File struct {
	Types []Type
}
