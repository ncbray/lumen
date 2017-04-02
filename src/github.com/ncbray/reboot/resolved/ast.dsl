struct Field {
  var Name string;
  var Type Type;
}

struct Variant {
  var Name string;
  var Fields []Field;
}

enum Type {
  Intrinsic {
    var Name string;
    var List Type;
  }
  Enum {
    var Name string;
    var Variants []Variant;
    var List Type;
  }
  Struct {
    var Name string;
    var Fields []Field;
    var List Type;
  }
  List {
    var Element Type;
    var List Type;
  }
}

// Index of what should be output
struct File {
  var Types []Type;
}
