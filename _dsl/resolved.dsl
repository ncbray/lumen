struct IntrinsicType {
  var Name string;
}

holder Type = IntrinsicType;

struct Local {
  var Name string;
  var Type Type;
}

struct GetInput {
  var Name string;
}

struct GetLocal {
  var Local Local;
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

holder Expr = GetInput | GetLocal | Number | Prefix | Infix | Constructor | CallIntrinsic;

struct SetOutput {
  var Name string;
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
  var Fs Function;
  var Vs Function;
}

struct File {
  var Programs []ShaderProgram;
}
