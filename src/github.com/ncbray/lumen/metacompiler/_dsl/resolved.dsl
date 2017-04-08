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
}
