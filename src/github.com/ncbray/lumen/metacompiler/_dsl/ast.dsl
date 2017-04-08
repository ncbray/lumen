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


struct KeywordArg {
  var Name string;
  var Value ParserBindingExpr;
}

enum ParserBindingExpr {
  Construct {
    var Name string;
    var Args []KeywordArg;
  }
  NameRef {
    var Name string;
  }
}

struct ParamDecl {
  var Name string;
  var T TypeRef;
}

struct ParserBindingMapping {
  var Name string;
  var Params []ParamDecl;
  var Body ParserBindingExpr;
}

struct ParserBindingGroup {
  var Name string;
  var T TypeRef;
  var Mappings []ParserBindingMapping;
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
  ParserBindingDecl {
    var Groups []ParserBindingGroup;
  }
}

// Compilation unit
struct File {
  var Decls []Declaration;
}


parser {
  typeRef:TypeRef {
    typeName(name string) => TypeName{Loc: loc_start, Name: name}
    listRef(element typeRef) => ListRef{Loc: loc_start, Element: element}
  }
  memberDecl:MemberDecl {
    fieldDecl(name string, t typeRef) => FieldDecl{Loc: loc_start, Name: name, T: t}
  }
  variantDecl:VariantDecl {
    default(name string, members []memberDecl) => VariantDecl{Loc: loc_start, Name: name, Members: members}
  }
  keywordArg:KeywordArg {
    default(name string, value parserBindingExpr) => KeywordArg{Loc: loc_start, Name: name, Value: value}
  }
  parserBindingExpr:ParserBindingExpr {
    construct(name string, args []keywordArg) => Construct{Loc: loc_start, Name: name, Args: args}
    nameRef(name string) => NameRef{Loc: loc_start, Name: name}
  }
  paramDecl:ParamDecl {
    default(name string, t typeRef) => ParamDecl{Loc: loc_start, Name: name, T: t}
  }
  parserBindingMapping:ParserBindingMapping {
    default(name string, params []paramDecl, body parserBindingExpr) => ParserBindingMapping{Loc: loc_start, Name: name, Params: params, Body: body}
  }
  parserBindingGroup:ParserBindingGroup {
    default(name string, t typeRef, mappings []parserBindingMapping) => ParserBindingGroup{Loc: loc_start, Name: name, T: t, Mappings: mappings}
  }
  declaration:Declaration {
    enumDecl(name string, variants []variantDecl) => EnumDecl{Loc: loc_start, Name: name, Variants: variants}
    structDecl(name string, members []memberDecl) => StructDecl{Loc: loc_start, Name: name, Members: members}
    holderDecl(name string, types []typeRef) => HolderDecl{Loc: loc_start, Name: name, Types: types}
    parserBindingDecl(groups []parserBindingGroup) => ParserBindingDecl{Loc: loc_start, Groups: groups}
  }
  file:File {
    default(decls []declaration) => File{Loc: loc_start, Decls: decls}
  }
}
