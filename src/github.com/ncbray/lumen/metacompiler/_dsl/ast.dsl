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

// Top-level declarations
enum Declaration {
  EnumDecl {
    var Name string;
    var Variants []VariantDecl;
  }
  StructDecl {
    var Name string;
    var Members []MemberDecl;
  }
  HolderDecl {
    var Name string;
    var Types []TypeRef;
  }
}

// Compilation unit
struct File {
  var Decls []Declaration;
}
