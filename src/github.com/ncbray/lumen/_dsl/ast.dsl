enum Expr {
  GetName {
    var Name string;
  }
  Number {
    var Raw string;
  }
  Prefix {
    var Op string;
    var Value Expr;
  }
  Infix {
    var Left Expr;
    var Op string;
    var Right Expr;
  }
  Call {
    var Value Expr;
    var Args []Expr;
  }
}

enum Statement {
  VarDecl {
    var T string;
    var Name string;
    var Value Expr;
  }
  Assign {
    var Name string;
    var Value Expr;
  }
  Discard {
    var Value Expr;
  }
}

struct ShaderDecl {
  var Name string;
  var Fs []Statement;
  var Vs []Statement;
}

struct File {
  var Shaders []ShaderDecl;
}
