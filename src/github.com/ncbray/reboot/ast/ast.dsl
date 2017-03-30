enum TypeRef {
  TypeName {
    var Name string;
  }
  ListRef {
    var Element TypeRef;
  }
}

enum MemberDecl {
  FieldDecl {
    var Name string;
    var T TypeRef;
  }
}

struct VariantDecl {
  var Name string;
  var Members []MemberDecl;
}

enum Declaration {
  EnumDecl {
    var Name string;
    var Variants []VariantDecl;
  }
  StructDecl {
    var Name string;
    var Members []MemberDecl;
  }
}

struct File {
  var Decls []Declaration;
}
