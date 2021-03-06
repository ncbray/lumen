struct Field {
  var Loc location;
  var Name string;
  var Type string;
}

struct Format {
  var Loc location;
  var Fields []Field;
}

struct GetName {
  var Loc location;
  var Name string;
}

struct GetAttr {
  var Loc location;
  var Value Expr;
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

holder Expr = GetName | GetAttr | Number | Prefix | Infix | Call;


struct VarDecl {
  var Loc location;
  var Type string;
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
  var Uniform Format;
  var Attribute Format;
  var Varying Format;
  var Vs []Statement;
  var Fs []Statement;
}

struct VertexComponent {
  var Loc location;
  var Name string;
  var Type string;
  var Encoding string;
}

struct VertexDecl {
  var Loc location;
  var Name string;
  var Components []VertexComponent;
}

holder TopLevelDecl = ShaderDecl | VertexDecl;

struct File {
  var Decls []TopLevelDecl;
}

parser {
  field:Field {
    default(name string, t string) => Field{Loc: loc_start, Name: name, Type: t}
  }
  format:Format {
    default(fields []field) => Format{Loc: loc_start, Fields: fields}
  }
  expr:Expr {
    getName(name string) => GetName{Loc: loc_start, Name: name}
    getAttr(value expr, name string) => GetAttr{Loc: loc_start, Value: value, Name: name}
    number(raw string) => Number{Loc: loc_start, Raw: raw}
    prefix(op string, value expr) => Prefix{Loc: loc_start, Op: op, Value: value}
    infix(left expr, op string, right expr) => Infix{Loc: loc_start, Left: left, Op: op, Right: right}
    call(value expr, args []expr) => Call{Loc: loc_start, Value: value, Args: args}
  }
  statement:Statement {
    varDecl(t string, name string, value expr) => VarDecl{Loc: loc_start, Type: t, Name: name, Value: value}
    assign(name string, value expr) => Assign{Loc: loc_start, Name: name, Value: value}
    discard(value expr) => Discard{Value: value}
  }
  vertexComponent:VertexComponent {
    default(name string, t string, encoding string) => VertexComponent{Loc: loc_start, Name: name, Type: t, Encoding: encoding}
  }
  topLevelDecl:TopLevelDecl {
    shaderDecl(name string, uniform format, attribute format, varying format, vs []statement, fs []statement) => ShaderDecl{
        Loc: loc_start,
        Name: name,
        Uniform: uniform,
        Attribute: attribute,
        Varying: varying,
        Vs: vs,
        Fs: fs}
    vertexDecl(name string, components []vertexComponent) => VertexDecl{Loc: loc_start, Name: name, Components: components}
  }
  file:File {
    default(decls []topLevelDecl) => File{Decls: decls}
  }
}
