package resolved

import (
	"fmt"

	"github.com/ncbray/lumen/ast"
	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/util"
)

type astConverter struct {
	Logger    log.CompilerLogger
	LocalLut  map[string]*Local
	Locals    []*Local
	Inputs    map[string]*Field
	Outputs   map[string]*Field
	Namespace map[string]TreeValue

	Float     *ScalarType
	Vec2      *VectorType
	Vec3      *VectorType
	Vec4      *VectorType
	Mat2      *MatrixType
	Mat3      *MatrixType
	Mat4      *MatrixType
	Sampler2D *ScalarType
}

func valueInfo(v TreeValue) string {
	switch v := v.(type) {
	case *ExprValue:
		return "an expression"
	case *TypeValue:
		return "type"
	case *FunctionValue:
		return "function " + v.Function.Name
	default:
		panic(v)
	}
}

func (conv *astConverter) requireExpr(src ast.Expr, ok *bool) *ExprValue {
	v, loc := conv.handleExpr(src)
	switch v := v.(type) {
	case *ExprValue:
		if len(v.Type) == 0 {
			conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "tried to use a void expression")
			*ok = false
			return nil
		}
		return v
	case *Poison:
		// Error was reported elsewhere.
		*ok = false
		return nil
	default:
		desc := valueInfo(v)
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "expected an expression but got "+desc)
		*ok = false
		return nil
	}
}

func (conv *astConverter) unwrapExprList(src []*ExprValue) []Expr {
	dst := make([]Expr, len(src))
	for i, v := range src {
		dst[i] = v.Expr
	}
	return dst
}

func (conv *astConverter) checkSwizzleMask(mask string, vecLength int, loc util.Location) bool {
	if len(mask) <= 0 || len(mask) > 4 {
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "swizzle mask must be between 1 and 4 characters")
		return false
	}
	for _, c := range mask {
		var swizzleIndex int
		switch c {
		case 'x', 'r':
			swizzleIndex = 0
		case 'y', 'g':
			swizzleIndex = 1
		case 'z', 'b':
			swizzleIndex = 2
		case 'w', 'a':
			swizzleIndex = 3
		default:
			conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("unknown swizzle character %q", c))
			return false
		}
		if swizzleIndex >= vecLength {
			conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("attempted to access element %d of a %d element vector", swizzleIndex, vecLength))
			return false
		}
		// TODO validate mixing xyzw with rgba.
	}
	return true
}

func (conv *astConverter) handleExpr(src ast.Expr) (TreeValue, util.Location) {
	subExprsOK := true

	switch src := src.(type) {
	case *ast.Number:
		return &ExprValue{
			Expr: &Number{
				Raw: src.Raw,
			},
			Type: []Type{conv.Float},
		}, src.Loc
	case *ast.GetName:
		name := src.Name
		lcl, ok := conv.LocalLut[name]
		if ok {
			return &ExprValue{
				Expr: &GetLocal{
					Local: lcl,
				},
				Type: []Type{lcl.Type},
			}, src.Loc
		}
		inp, ok := conv.Inputs[name]
		if ok {
			return &ExprValue{
				Expr: &GetInput{
					Input: inp,
				},
				Type: []Type{inp.Type},
			}, src.Loc
		}
		tv, ok := conv.Namespace[name]
		if ok {
			return tv, src.Loc
		}

		loc := src.Loc
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("cannot resolve name %q", name))
		return &Poison{}, src.Loc
	case *ast.GetAttr:
		expr := conv.requireExpr(src.Value, &subExprsOK)
		if !subExprsOK {
			return &Poison{}, src.Loc
		}

		vecLength := 0
		switch expr.Type[0] {
		case conv.Vec2:
			vecLength = 2
		case conv.Vec3:
			vecLength = 3
		case conv.Vec4:
			vecLength = 4
		default:
			loc := src.Loc
			conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "cannot get attribute")
			subExprsOK = false
		}
		if !subExprsOK {
			return &Poison{}, src.Loc
		}
		if !conv.checkSwizzleMask(src.Name, vecLength, src.Loc) {
			return &Poison{}, src.Loc
		}
		var swizzType Type
		switch len(src.Name) {
		case 1:
			swizzType = conv.Float
		case 2:
			swizzType = conv.Vec2
		case 3:
			swizzType = conv.Vec3
		case 4:
			swizzType = conv.Vec4
		default:
			panic(src.Name)
		}
		return &ExprValue{
			Expr: &GetAttr{
				Value: expr.Expr,
				Name:  src.Name,
			},
			Type: []Type{swizzType},
		}, src.Loc
	case *ast.Infix:
		left := conv.requireExpr(src.Left, &subExprsOK)
		right := conv.requireExpr(src.Right, &subExprsOK)
		if !subExprsOK {
			return &Poison{}, src.Loc
		}
		return &ExprValue{
			Expr: &Infix{
				Left:  left.Expr,
				Op:    src.Op,
				Right: right.Expr,
			},
			Type: left.Type, // HACK - should actually check types.
		}, src.Loc
	case *ast.Call:
		value, loc := conv.handleExpr(src.Value)
		args := make([]*ExprValue, len(src.Args))
		for i, arg := range src.Args {
			args[i] = conv.requireExpr(arg, &subExprsOK)
		}

		switch value := value.(type) {
		case *TypeValue:
			if !subExprsOK {
				return &Poison{}, loc
			}
			return &ExprValue{
				Expr: &Constructor{
					Type: value.Type,
					Args: conv.unwrapExprList(args),
				},
				Type: []Type{value.Type},
			}, loc
		case *FunctionValue:
			if !subExprsOK {
				return &Poison{}, loc
			}
			return &ExprValue{
				Expr: &CallIntrinsic{
					Function: value.Function,
					Args:     conv.unwrapExprList(args),
				},
				Type: value.Function.Signature.Returns,
			}, loc
		case *Poison:
			return value, loc
		default:
			desc := valueInfo(value)
			conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, desc+" is not callable")
			return &Poison{}, loc
		}
	default:
		panic(src)
	}
}

func (conv *astConverter) resolveTypeRef(name string, loc util.Location) Type {
	v, ok := conv.Namespace[name]
	if !ok {
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("unknown type %q", name))
		return nil
	}
	switch v := v.(type) {
	case *TypeValue:
		return v.Type
	default:
		desc := valueInfo(v)
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, desc+" is not a type")
		return nil
	}
}

func (conv *astConverter) handleBody(src []ast.Statement, logger log.CompilerLogger) ([]*Local, []Statement) {
	conv.Locals = []*Local{}
	conv.LocalLut = map[string]*Local{}

	dst := []Statement{}
	for _, stmt := range src {
		subExprsOK := true
		switch stmt := stmt.(type) {
		case *ast.VarDecl:
			value := conv.requireExpr(stmt.Value, &subExprsOK)
			// TODO type?
			_, ok := conv.LocalLut[stmt.Name]
			if ok {
				loc := stmt.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("attempted to redefine %q", stmt.Name))
			}
			t := conv.resolveTypeRef(stmt.Type, stmt.Loc)
			lcl := &Local{Name: stmt.Name, Type: t}
			conv.Locals = append(conv.Locals, lcl)
			conv.LocalLut[lcl.Name] = lcl
			if !subExprsOK {
				continue
			}
			dst = append(dst, &SetLocal{
				Local: lcl,
				Value: value.Expr,
			})
		case *ast.Assign:
			value := conv.requireExpr(stmt.Value, &subExprsOK)
			lcl, ok := conv.LocalLut[stmt.Name]
			if ok {
				if !subExprsOK {
					continue
				}
				dst = append(dst, &SetLocal{
					Local: lcl,
					Value: value.Expr,
				})
				continue
			}
			outp, ok := conv.Outputs[stmt.Name]
			if ok {
				if !subExprsOK {
					continue
				}
				dst = append(dst, &SetOutput{
					Output: outp,
					Value:  value.Expr,
				})
				continue
			}

			loc := stmt.Loc
			logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("cannot assign to %q", stmt.Name))
		default:
			panic(stmt)
		}
	}
	return conv.Locals, dst
}

func (conv *astConverter) handleFormat(node *ast.Format) *Format {
	fields := []*Field{}
	for _, f := range node.Fields {
		fields = append(fields, &Field{
			Name: f.Name,
			Type: conv.resolveTypeRef(f.Type, f.Loc),
		})
	}
	return &Format{Fields: fields}
}

func (conv *astConverter) declType(name string, t Type) {
	conv.Namespace[name] = &TypeValue{Type: t}
}

func FromAST(src *ast.File, logger log.CompilerLogger) *File {
	c := astConverter{
		Logger:    logger,
		Namespace: map[string]TreeValue{},
	}

	c.Float = &ScalarType{Name: "float"}
	c.declType(c.Float.Name, c.Float)

	c.Vec2 = &VectorType{Name: "vec2", Width: 2, Scalar: c.Float}
	c.declType(c.Vec2.Name, c.Vec2)
	c.Vec3 = &VectorType{Name: "vec3", Width: 3, Scalar: c.Float}
	c.declType(c.Vec3.Name, c.Vec3)
	c.Vec4 = &VectorType{Name: "vec4", Width: 4, Scalar: c.Float}
	c.declType(c.Vec4.Name, c.Vec4)

	c.Mat2 = &MatrixType{Name: "mat2", Width: 2, Height: 2, Scalar: c.Float}
	c.declType(c.Mat2.Name, c.Mat2)
	c.Mat3 = &MatrixType{Name: "mat3", Width: 3, Height: 3, Scalar: c.Float}
	c.declType(c.Mat3.Name, c.Mat3)
	c.Mat4 = &MatrixType{Name: "mat4", Width: 4, Height: 4, Scalar: c.Float}
	c.declType(c.Mat4.Name, c.Mat4)

	c.Sampler2D = &ScalarType{Name: "sampler2D"}
	c.declType(c.Sampler2D.Name, c.Sampler2D)

	f := &IntrinsicFunction{
		Name: "texture2D",
		Signature: &FunctionSignature{
			Params:  []Type{c.Sampler2D, c.Vec2},
			Returns: []Type{c.Vec4},
		},
	}
	c.Namespace[f.Name] = &FunctionValue{Function: f}

	vertex := []*VertexFormat{}
	shaders := []*ShaderProgram{}
	for _, d := range src.Decls {
		switch d := d.(type) {
		case *ast.VertexDecl:
			components := []*VertexComponent{}
			for _, comp := range d.Components {
				t := c.resolveTypeRef(comp.Type, comp.Loc)
				components = append(components, &VertexComponent{
					Name:     comp.Name,
					Type:     t,
					Encoding: comp.Encoding, // TODO validate.
				})
			}
			vertex = append(vertex, &VertexFormat{
				Name:       d.Name,
				Components: components,
			})
		case *ast.ShaderDecl:
			uniform := c.handleFormat(d.Uniform)
			attribute := c.handleFormat(d.Attribute)
			varying := c.handleFormat(d.Varying)

			c.Inputs = map[string]*Field{}
			for _, f := range uniform.Fields {
				c.Inputs[f.Name] = f
			}
			for _, f := range attribute.Fields {
				c.Inputs[f.Name] = f
			}
			c.Outputs = map[string]*Field{}
			c.Outputs["gl_Position"] = &Field{Name: "gl_Position", Type: c.Vec4}
			for _, f := range varying.Fields {
				c.Outputs[f.Name] = f
			}

			locals, body := c.handleBody(d.Vs, logger)
			vs := &Function{
				Name:   "main",
				Locals: locals,
				Body:   body,
			}

			c.Inputs = map[string]*Field{}
			for _, f := range uniform.Fields {
				c.Inputs[f.Name] = f
			}
			for _, f := range varying.Fields {
				c.Inputs[f.Name] = f
			}
			c.Outputs = map[string]*Field{}
			c.Outputs["gl_FragColor"] = &Field{Name: "gl_FragColor", Type: c.Vec4}

			locals, body = c.handleBody(d.Fs, logger)
			fs := &Function{
				Name:   "main",
				Locals: locals,
				Body:   body,
			}

			shaders = append(shaders, &ShaderProgram{
				Name:      d.Name,
				Uniform:   uniform,
				Attribute: attribute,
				Varying:   varying,
				Vs:        vs,
				Fs:        fs,
			})
		default:
			panic(d)
		}
	}
	return &File{
		Vertex:   vertex,
		Programs: shaders,
	}
}
