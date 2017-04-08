struct Field {
  var Name string;
  var Type Type;
}

struct Variant {
  var Name string;
  var Fields []Field;
}

struct Intrinsic {
  var Name string;
  var List Type;
}

struct Enum {
  var Name string;
  var Variants []Variant;
  var List Type;
}

struct Struct {
  var Name string;
  var Fields []Field;
  var Holders []Holder;
  var List Type;
}

struct Holder {
  var Name string;
  var Types []Struct;
  var List Type;
}

struct List {
  var Element Type;
  var List Type;
}

struct KeywordArg {
  var Name string;
  var Value ParserBindingExpr;
}

struct Construct {
  // TODO directly ref Type object once enums are eliminated.
  var Type string;
  var Args []KeywordArg;
}

struct GetParam {
  var Param ParserBindingParam;
}

struct GetLocStart {
}

holder ParserBindingExpr {
  Construct;
  GetParam;
  GetLocStart;
}

struct InputList {
  var Element ParserBindingGroup;
}

struct InputSlice {
}

holder ParserBindingInput {
  ParserBindingGroup;
  InputList;
  InputSlice;
}

struct ParserBindingParam {
  var Name string;
  var Input ParserBindingInput;
}

struct ParserBindingMapping {
  var Rule string;
  var Params []ParserBindingParam;
  var Body ParserBindingExpr;
}

struct ParserBindingGroup {
  var Name string;
  var Type Type;
  var Mappings []ParserBindingMapping;
  var List InputList;
}

struct ParserBinding {
  var Groups []ParserBindingGroup;
}

holder Type {
  Intrinsic;
  Enum;
  Struct;
  Holder;
  List;
}

// Index of what should be output
struct File {
  var Types []Type;
  var ParserBinding ParserBinding;
}
