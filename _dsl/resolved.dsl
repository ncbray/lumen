struct IntrinsicType {
  var Name string;
}

holder Type = IntrinsicType;

struct Field {
  var Name string;
  var Type Type;
}

struct Format {
  var Fields []Field;
}

struct Local {
  var Name string;
  var Type Type;
}

struct GetInput {
  var Input Field;
}

struct GetLocal {
  var Local Local;
}

struct GetAttr {
  var Value Expr;
  var Name string;
}

struct Number {
  var Raw string;
}

struct Prefix {
  var Op string;
  var Value Expr;
}

struct Infix {
  var Left Expr;
  var Op string;
  var Right Expr;
}

struct Constructor {
  var Type Type;
  var Args []Expr;
}

struct CallIntrinsic {
  var Name string;
  var Args []Expr;
}

holder Expr = GetInput | GetLocal | GetAttr | Number | Prefix | Infix | Constructor | CallIntrinsic;

struct ExprValue {
  var Loc location;
  var Expr Expr;
  var Type Type;
}

struct TypeValue {
  var Loc location;
  var Type Type;
}

struct FunctionValue {
  var Loc location;
  var Name string;
}

// TODO holder of holders?
holder TreeValue = ExprValue | TypeValue | FunctionValue;

struct SetOutput {
  var Output Field;
  var Value Expr;
}

struct SetLocal {
  var Local Local;
  var Value Expr;
}

struct Discard {
  var Value Expr;
}

holder Statement = SetOutput | SetLocal | Discard;

struct Function {
  var Name string;
  var Locals []Local;
  var Body []Statement;
}

struct ShaderProgram {
  var Name string;
  var Uniform Format;
  var Attribute Format;
  var Varying Format;
  var Fs Function;
  var Vs Function;
}

struct File {
  var Programs []ShaderProgram;
}
