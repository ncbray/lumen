package resolved

import (
	"fmt"

	"github.com/ncbray/lumen/ast"
	"github.com/ncbray/lumen/log"
	"github.com/ncbray/lumen/util"
)

type astConverter struct {
	Logger     log.CompilerLogger
	LocalLut   map[string]*Local
	Locals     []*Local
	Inputs     map[string]*Field
	Outputs    map[string]*Field
	Intrinsics map[string]*IntrinsicType
	Namespace  map[string]TreeValue
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
		}, src.Loc
	case *ast.GetName:
		name := src.Name
		lcl, ok := conv.LocalLut[name]
		if ok {
			return &ExprValue{
				Expr: &GetLocal{
					Local: lcl,
				},
			}, src.Loc
		}
		inp, ok := conv.Inputs[name]
		if ok {
			return &ExprValue{
				Expr: &GetInput{
					Input: inp,
				},
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
		return &ExprValue{
			Expr: &GetAttr{
				Value: conv.requireExpr(conv.handleExpr(src.Value)),
				Name:  src.Name,
			},
		}, src.Loc
	case *ast.Infix:
		return &ExprValue{
			Expr: &Infix{
				Left:  conv.requireExpr(conv.handleExpr(src.Left)),
				Op:    src.Op,
				Right: conv.requireExpr(conv.handleExpr(src.Right)),
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

// TODO eliminate this method.
func (conv *astConverter) resolveTypeRef(name string, loc util.Location) Type {
	t, ok := conv.Intrinsics[name]
	if !ok {
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("unknown type %q", name))
	}
	return t
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

func FromAST(src *ast.File, logger log.CompilerLogger) *File {
	c := astConverter{
		Logger:     logger,
		Intrinsics: map[string]*IntrinsicType{},
		Namespace:  map[string]TreeValue{},
	}

	for _, name := range []string{"vec2", "vec3", "vec4", "mat2", "mat3", "mat4", "sampler2D"} {
		it := &IntrinsicType{Name: name}
		c.Intrinsics[name] = it
		c.Namespace[name] = &TypeValue{Type: it}
	}

	for _, name := range []string{"texture2D"} {
		f := &IntrinsicFunction{Name: name}
		c.Namespace[name] = &FunctionValue{Function: f}
	}

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
		c.Outputs["gl_Position"] = &Field{Name: "gl_Position", Type: c.Intrinsics["vec4"]}
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
		c.Outputs["gl_FragColor"] = &Field{Name: "gl_FragColor", Type: c.Intrinsics["vec4"]}

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
