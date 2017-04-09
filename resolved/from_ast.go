package resolved

import (
	"fmt"

	"github.com/ncbray/lumen/ast"
	"github.com/ncbray/lumen/log"
)

type astConverter struct {
	Logger     log.CompilerLogger
	LocalLut   map[string]*Local
	Locals     []*Local
	Intrinsics map[string]*IntrinsicType
}

func (conv *astConverter) handleExpr(src ast.Expr, logger log.CompilerLogger) Expr {
	switch src := src.(type) {
	case *ast.Number:
		return &Number{
			Raw: src.Raw,
		}
	case *ast.GetName:
		lcl, ok := conv.LocalLut[src.Name]
		if ok {
			return &GetLocal{
				Local: lcl,
			}
		}
		return &GetInput{
			Name: src.Name,
		}
	case *ast.Infix:
		return &Infix{
			Left:  conv.handleExpr(src.Left, logger),
			Op:    src.Op,
			Right: conv.handleExpr(src.Right, logger),
		}
	case *ast.Call:
		conv.handleExpr(src.Value, logger)
		args := make([]Expr, len(src.Args))
		for i, arg := range src.Args {
			args[i] = conv.handleExpr(arg, logger)
		}
		return &CallIntrinsic{
			Name: "foo",
			Args: args,
		}
	default:
		panic(src)
	}
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
			t, ok := conv.Intrinsics[stmt.T]
			if !ok {
				loc := stmt.Loc
				logger.ErrorAtLocation(loc.File, loc.Line, loc.Column, fmt.Sprintf("unknown type %q", stmt.T))
			}
			lcl := &Local{Name: stmt.Name, Type: t}
			conv.Locals = append(conv.Locals, lcl)
			conv.LocalLut[lcl.Name] = lcl
			dst = append(dst, &SetLocal{
				Local: lcl,
				Value: conv.handleExpr(stmt.Value, logger),
			})
		case *ast.Assign:
			dst = append(dst, &SetOutput{
				Name:  stmt.Name,
				Value: conv.handleExpr(stmt.Value, logger),
			})
		default:
			panic(stmt)
		}
	}
	return conv.Locals, dst
}

func FromAST(src *ast.File, logger log.CompilerLogger) *File {
	c := astConverter{
		Logger: logger,
		Intrinsics: map[string]*IntrinsicType{
			"vec2": {Name: "vec2"},
			"vec3": {Name: "vec3"},
			"vec4": {Name: "vec4"},
		},
	}

	shaders := []*ShaderProgram{}
	for _, s := range src.Shaders {
		locals, body := c.handleBody(s.Vs, logger)
		vs := &Function{
			Name:   "main",
			Locals: locals,
			Body:   body,
		}

		locals, body = c.handleBody(s.Fs, logger)
		fs := &Function{
			Name:   "main",
			Locals: locals,
			Body:   body,
		}

		shaders = append(shaders, &ShaderProgram{
			Name: s.Name,
			Vs:   vs,
			Fs:   fs,
		})
	}
	return &File{
		Programs: shaders,
	}
}
