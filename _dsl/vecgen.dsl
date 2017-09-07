struct VectorType {
  var Width int;
  var Name string;
}

struct MatrixType {
  var Width int;
  var Name string;
}

holder Type = VectorType | MatrixType;

struct Object {
  var Type Type;
  var Name string;
  var CachedGetter []string;
  var CachedSetter []string;
}

struct Atom {
  var Text string;
}

struct Neg {
  var Expr Expr;
}

struct Add {
  var Left Expr;
  var Right Expr;
}

struct Sub {
  var Left Expr;
  var Right Expr;
}

struct Mul {
  var Left Expr;
  var Right Expr;
}

struct Div {
  var Left Expr;
  var Right Expr;
}

struct GetComponent {
  var Object Object;
  var Index int;
}

holder Expr = Atom | GetComponent | Neg | Add | Sub | Mul | Div;

struct SetComponent {
  var Expr Expr;
  var Object Object;
  var Index int;
}
