enum Type {
  Intrinsic {
    var Name string;
  }
}

struct Local {
  var Name string;
  var Type Type;
}

enum Expr {
  GetInput {
    var Name string;
  }
  GetLocal {
    var Local Local;
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
  Constructor {
    var Type Type;
    var Args []Expr;
  }
  CallIntrinsic {
    var Name string;
    var Args []Expr;
  }
}

enum Statement {
  SetOutput {
    var Name string;
    var Value Expr;
  }
  SetLocal {
    var Local Local;
    var Value Expr;
  }
  Discard {
    var Value Expr;
  }
}

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
