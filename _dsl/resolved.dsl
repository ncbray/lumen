struct ScalarType {
  var Name string;
}

struct VectorType {
  var Name string;
  var Width int;
  var Scalar ScalarType;
}

struct MatrixType {
  var Name string;
  var Width int;
  var Height int;
  var Scalar ScalarType;
}

holder Type = ScalarType | VectorType | MatrixType;

struct Field {
  var Name string;
  var Type Type;
  var Marked bool; // For analysis.
}

struct Format {
  var Fields []Field;
}

struct Local {
  var Name string;
  var Type Type;
}

struct GetInput {
  var Input Field;
}

struct GetLocal {
  var Local Local;
}

struct GetAttr {
  var Value Expr;
  var Name string;
}

struct Number {
  var Raw string;
}

struct Prefix {
  var Op string;
  var Value Expr;
}

struct Infix {
  var Left Expr;
  var Op string;
  var Right Expr;
}

struct Constructor {
  var Type Type;
  var Args []Expr;
}

struct FunctionSignature {
  var Params []Type;
  var Returns []Type;
}

struct IntrinsicFunction {
  var Name string;
  var Signature FunctionSignature;
}

struct CallIntrinsic {
  var Function IntrinsicFunction;
  var Args []Expr;
}

holder Expr = GetInput | GetLocal | GetAttr | Number | Prefix | Infix | Constructor | CallIntrinsic;

struct ExprValue {
  var Expr Expr;
  var Type []Type;
}

struct TypeValue {
  var Type Type;
}

struct FunctionValue {
  var Function IntrinsicFunction;
}

struct Poison {
}

// TODO holder of holders?
holder TreeValue = ExprValue | TypeValue | FunctionValue | Poison;

struct SetOutput {
  var Output Field;
  var Value Expr;
}

struct SetLocal {
  var Local Local;
  var Value Expr;
}

struct Discard {
  var Value Expr;
}

holder Statement = SetOutput | SetLocal | Discard;

struct Function {
  var Name string;
  var Locals []Local;
  var Body []Statement;
}

struct ShaderProgram {
  var Name string;
  var Uniform Format;
  var Attribute Format;
  var Varying Format;
  var Fs Function;
  var Vs Function;
}

struct VertexComponent {
  var Name string;
  var Type Type;
  var Encoding string; // TODO turn into enumeration?

  var ByteOffset int;
  var ByteSize int;
}

struct VertexFormat {
  var Name string;
  var Components []VertexComponent;

  var ByteSize int;
}

struct File {
  var Vertex []VertexFormat;
  var Programs []ShaderProgram;
}
