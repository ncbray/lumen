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
}

func valueInfo(v TreeValue) (string, util.Location) {
	switch v := v.(type) {
	case *ExprValue:
		return "an expression", v.Loc
	case *TypeValue:
		return "type", v.Loc
	case *FunctionValue:
		return "function " + v.Name, v.Loc
	default:
		panic(v)
	}
}

func (conv *astConverter) requireExpr(v TreeValue) Expr {
	switch v := v.(type) {
	case *ExprValue:
		return v.Expr
	default:
		desc, loc := valueInfo(v)
		conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, "expected an expression but got "+desc)
		return nil
	}
}

func (conv *astConverter) handleExpr(src ast.Expr) TreeValue {
	switch src := src.(type) {
	case *ast.Number:
		return &ExprValue{
			Loc: src.Loc,
			Expr: &Number{
				Raw: src.Raw,
			},
		}
	case *ast.GetName:
		name := src.Name
		lcl, ok := conv.LocalLut[name]
		if ok {
			return &ExprValue{
				Loc: src.Loc,
				Expr: &GetLocal{
					Local: lcl,
				},
			}
		}
		_, ok = conv.Inputs[name]
		if ok {
			return &ExprValue{
				Loc: src.Loc,
				Expr: &GetInput{
					Name: name,
				},
			}
		}
		t, ok := conv.Intrinsics[name]
		if ok {
			return &TypeValue{
				Loc:  src.Loc,
				Type: t,
			}
		}

		return &FunctionValue{
			Loc:  src.Loc,
			Name: name,
		}

		//loc := src.Loc
		//conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("cannot resolve name %q", name))
		//return nil
	case *ast.Infix:
		return &ExprValue{
			Expr: &Infix{
				Left:  conv.requireExpr(conv.handleExpr(src.Left)),
				Op:    src.Op,
				Right: conv.requireExpr(conv.handleExpr(src.Right)),
			},
		}
	case *ast.Call:
		value := conv.handleExpr(src.Value)
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
			}
		case *FunctionValue:
			return &ExprValue{
				Expr: &CallIntrinsic{
					Name: value.Name,
					Args: args,
				},
			}
		default:
			desc, loc := valueInfo(value)
			conv.Logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, desc+" is not callable")
			return nil
		}
	default:
		panic(src)
	}
}

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
				Value: conv.requireExpr(conv.handleExpr(stmt.Value)),
			})
		case *ast.Assign:
			dst = append(dst, &SetOutput{
				Name:  stmt.Name,
				Value: conv.requireExpr(conv.handleExpr(stmt.Value)),
			})
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
		Logger: logger,
		Intrinsics: map[string]*IntrinsicType{
			"vec2":      {Name: "vec2"},
			"vec3":      {Name: "vec3"},
			"vec4":      {Name: "vec4"},
			"mat2":      {Name: "mat2"},
			"mat3":      {Name: "mat3"},
			"mat4":      {Name: "mat4"},
			"sampler2D": {Name: "sampler2D"},
		},
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
