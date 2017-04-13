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

	Float     *IntrinsicType
	Vec2      *IntrinsicType
	Vec3      *IntrinsicType
	Vec4      *IntrinsicType
	Mat2      *IntrinsicType
	Mat3      *IntrinsicType
	Mat4      *IntrinsicType
	Sampler2D *IntrinsicType
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

func (conv *astConverter) requireExpr(v TreeValue, loc util.Location) Expr {
	switch v := v.(type) {
	case *ExprValue:
		return v.Expr
	case *Poison:
		// Error was reported elsewhere.
		return nil
	default:
		desc := valueInfo(v)
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "expected an expression but got "+desc)
		return nil
	}
}

func (conv *astConverter) handleExpr(src ast.Expr) (TreeValue, util.Location) {
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
		expr := conv.requireExpr(conv.handleExpr(src.Value))
		return &ExprValue{
			Expr: &GetAttr{
				Value: expr,
				Name:  src.Name,
			},
		}, src.Loc
	case *ast.Infix:
		left := conv.requireExpr(conv.handleExpr(src.Left))
		right := conv.requireExpr(conv.handleExpr(src.Right))
		return &ExprValue{
			Expr: &Infix{
				Left:  left,
				Op:    src.Op,
				Right: right,
			},
		}, src.Loc
	case *ast.Call:
		value, loc := conv.handleExpr(src.Value)
		args := make([]Expr, len(src.Args))
		for i, arg := range src.Args {
			args[i] = conv.requireExpr(conv.handleExpr(arg))
		}

		switch value := value.(type) {
		case *TypeValue:
			return &ExprValue{
				Expr: &Constructor{
					Type: value.Type,
					Args: args,
				},
			}, loc
		case *FunctionValue:
			return &ExprValue{
				Expr: &CallIntrinsic{
					Function: value.Function,
					Args:     args,
				},
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
		switch stmt := stmt.(type) {
		case *ast.VarDecl:
			value := conv.requireExpr(conv.handleExpr(stmt.Value))
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
			dst = append(dst, &SetLocal{
				Local: lcl,
				Value: value,
			})
		case *ast.Assign:
			value := conv.requireExpr(conv.handleExpr(stmt.Value))
			lcl, ok := conv.LocalLut[stmt.Name]
			if ok {
				dst = append(dst, &SetLocal{
					Local: lcl,
					Value: value,
				})
				continue
			}
			outp, ok := conv.Outputs[stmt.Name]
			if ok {
				dst = append(dst, &SetOutput{
					Output: outp,
					Value:  value,
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

func (conv *astConverter) declIntrinsicType(t *IntrinsicType) *IntrinsicType {
	conv.Namespace[t.Name] = &TypeValue{Type: t}
	return t
}

func FromAST(src *ast.File, logger log.CompilerLogger) *File {
	c := astConverter{
		Logger:    logger,
		Namespace: map[string]TreeValue{},
	}

	c.Float = c.declIntrinsicType(&IntrinsicType{Name: "float"})
	c.Vec2 = c.declIntrinsicType(&IntrinsicType{Name: "vec2"})
	c.Vec3 = c.declIntrinsicType(&IntrinsicType{Name: "vec3"})
	c.Vec4 = c.declIntrinsicType(&IntrinsicType{Name: "vec4"})
	c.Mat2 = c.declIntrinsicType(&IntrinsicType{Name: "mat2"})
	c.Mat3 = c.declIntrinsicType(&IntrinsicType{Name: "mat3"})
	c.Mat4 = c.declIntrinsicType(&IntrinsicType{Name: "mat4"})
	c.Sampler2D = c.declIntrinsicType(&IntrinsicType{Name: "sampler2D"})

	f := &IntrinsicFunction{
		Name: "texture2D",
		Signature: &FunctionSignature{
			Params:  []Type{c.Sampler2D, c.Vec2},
			Returns: []Type{c.Vec4},
		},
	}
	c.Namespace[f.Name] = &FunctionValue{Function: f}

	shaders := []*ShaderProgram{}
	for _, s := range src.Shaders {
		uniform := c.handleFormat(s.Uniform)
		attribute := c.handleFormat(s.Attribute)
		varying := c.handleFormat(s.Varying)

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

		locals, body := c.handleBody(s.Vs, logger)
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

		locals, body = c.handleBody(s.Fs, logger)
		fs := &Function{
			Name:   "main",
			Locals: locals,
			Body:   body,
		}

		shaders = append(shaders, &ShaderProgram{
			Name:      s.Name,
			Uniform:   uniform,
			Attribute: attribute,
			Varying:   varying,
			Vs:        vs,
			Fs:        fs,
		})
	}
	return &File{
		Programs: shaders,
	}
}
