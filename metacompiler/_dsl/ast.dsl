struct TypeName {
  var Loc location;
  var Name string;
}

struct ListRef {
  var Loc location;
  var Element TypeRef;
}

holder TypeRef = TypeName | ListRef;

struct FieldDecl {
  var Loc location;
  var Name string;
  var T TypeRef;
}

holder MemberDecl = FieldDecl;

struct KeywordArg {
  var Loc location;
  var Name string;
  var Value ParserBindingExpr;
}

struct Construct {
  var Loc location;
  var Name string;
  var Args []KeywordArg;
}

struct NameRef {
  var Loc location;
  var Name string;
}

holder ParserBindingExpr = Construct | NameRef;

struct ParamDecl {
  var Loc location;
  var Name string;
  var T TypeRef;
}

struct ParserBindingMapping {
  var Loc location;
  var Name string;
  var Params []ParamDecl;
  var Body ParserBindingExpr;
}

struct ParserBindingGroup {
  var Loc location;
  var Name string;
  var T TypeRef;
  var Mappings []ParserBindingMapping;
}

struct StructDecl {
  var Loc location;
  var Name string;
  var Members []MemberDecl;
}

struct HolderDecl {
  var Loc location;
  var Name string;
  var Types []TypeRef;
}

struct ParserBindingDecl {
  var Loc location;
  var Groups []ParserBindingGroup;
}

// Top-level declarations
holder Declaration = StructDecl | HolderDecl | ParserBindingDecl;

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
    structDecl(name string, members []memberDecl) => StructDecl{Loc: loc_start, Name: name, Members: members}
    holderDecl(name string, types []typeRef) => HolderDecl{Loc: loc_start, Name: name, Types: types}
    parserBindingDecl(groups []parserBindingGroup) => ParserBindingDecl{Loc: loc_start, Groups: groups}
  }
  file:File {
    default(decls []declaration) => File{Decls: decls}
  }
}
