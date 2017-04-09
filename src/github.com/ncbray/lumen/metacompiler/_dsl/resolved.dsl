struct Field {
  var Name string;
  var Type Type;
}

struct IntrinsicType {
  var Name string;
  var List List;
}

struct Struct {
  var Name string;
  var Fields []Field;
  var Holders []Holder;
  var List List;
}

struct Holder {
  var Name string;
  var Types []Struct;
  var List List;
}

struct List {
  var Element Type;
  var List List;
}

holder Type = IntrinsicType | Struct | Holder | List;

struct KeywordArg {
  var Name string;
  var Value ParserBindingExpr;
}

struct Construct {
  var Type Struct;
  var Args []KeywordArg;
}

struct GetParam {
  var Param ParserBindingParam;
}

struct GetLocStart {
}

holder ParserBindingExpr = Construct | GetParam | GetLocStart;

struct InputList {
  var Element ParserBindingGroup;
}

struct InputSlice {
}

holder ParserBindingInput = ParserBindingGroup | InputList | InputSlice;

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

// Index of what should be output
struct File {
  var Types []Type;
  var ParserBinding ParserBinding;
}
