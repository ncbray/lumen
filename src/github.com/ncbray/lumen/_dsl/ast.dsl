struct GetName {
  var Loc location;
  var Name string;
}

struct Number {
  var Loc location;
  var Raw string;
}

struct Prefix {
  var Loc location;
  var Op string;
  var Value Expr;
}

struct Infix {
  var Loc location;
  var Left Expr;
  var Op string;
  var Right Expr;
}

struct Call {
  var Loc location;
  var Value Expr;
  var Args []Expr;
}

holder Expr = GetName | Number | Prefix | Infix | Call;


struct VarDecl {
  var Loc location;
  var T string;
  var Name string;
  var Value Expr;
}

struct Assign {
  var Loc location;
  var Name string;
  var Value Expr;
}

struct Discard {
  var Value Expr;
}

holder Statement = VarDecl | Assign | Discard;

struct ShaderDecl {
  var Loc location;
  var Name string;
  var Vs []Statement;
  var Fs []Statement;
}

struct File {
  var Shaders []ShaderDecl;
}

parser {
  expr:Expr {
    getName(name string) => GetName{Loc: loc_start, Name: name}
    number(raw string) => Number{Loc: loc_start, Raw: raw}
    prefix(op string, value expr) => Prefix{Loc: loc_start, Op: op, Value: value}
    infix(left expr, op string, right expr) => Infix{Loc: loc_start, Left: left, Op: op, Right: right}
    call(value expr, args []expr) => Call{Loc: loc_start, Value: value, Args: args}
  }
  statement:Statement {
    varDecl(t string, name string, value expr) => VarDecl{Loc: loc_start, T: t, Name: name, Value: value}
    assign(name string, value expr) => Assign{Loc: loc_start, Name: name, Value: value}
    discard(value expr) => Discard{Value: value}
  }
  shaderDecl:ShaderDecl {
    default(name string, vs []statement, fs []statement) => ShaderDecl{Loc: loc_start, Name: name, Vs: vs, Fs: fs}
  }
  file:File {
    default(shaders []shaderDecl) => File{Shaders: shaders}
  }
}
