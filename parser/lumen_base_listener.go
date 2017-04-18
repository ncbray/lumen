// Generated from Lumen.g by ANTLR 4.6.

package parser // Lumen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLumenListener is a complete listener for a parse tree produced by LumenParser.
type BaseLumenListener struct{}

var _ LumenListener = &BaseLumenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLumenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLumenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLumenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLumenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCall is called when production call is entered.
func (s *BaseLumenListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseLumenListener) ExitCall(ctx *CallContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseLumenListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLumenListener) ExitNumber(ctx *NumberContext) {}

// EnterGetName is called when production getName is entered.
func (s *BaseLumenListener) EnterGetName(ctx *GetNameContext) {}

// ExitGetName is called when production getName is exited.
func (s *BaseLumenListener) ExitGetName(ctx *GetNameContext) {}

// EnterPrefix is called when production prefix is entered.
func (s *BaseLumenListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production prefix is exited.
func (s *BaseLumenListener) ExitPrefix(ctx *PrefixContext) {}

// EnterGetAttr is called when production getAttr is entered.
func (s *BaseLumenListener) EnterGetAttr(ctx *GetAttrContext) {}

// ExitGetAttr is called when production getAttr is exited.
func (s *BaseLumenListener) ExitGetAttr(ctx *GetAttrContext) {}

// EnterInfix is called when production infix is entered.
func (s *BaseLumenListener) EnterInfix(ctx *InfixContext) {}

// ExitInfix is called when production infix is exited.
func (s *BaseLumenListener) ExitInfix(ctx *InfixContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseLumenListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseLumenListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseLumenListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseLumenListener) ExitAssign(ctx *AssignContext) {}

// EnterDiscard is called when production discard is entered.
func (s *BaseLumenListener) EnterDiscard(ctx *DiscardContext) {}

// ExitDiscard is called when production discard is exited.
func (s *BaseLumenListener) ExitDiscard(ctx *DiscardContext) {}

// EnterField is called when production field is entered.
func (s *BaseLumenListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLumenListener) ExitField(ctx *FieldContext) {}

// EnterFormat is called when production format is entered.
func (s *BaseLumenListener) EnterFormat(ctx *FormatContext) {}

// ExitFormat is called when production format is exited.
func (s *BaseLumenListener) ExitFormat(ctx *FormatContext) {}

// EnterVertexComponent is called when production vertexComponent is entered.
func (s *BaseLumenListener) EnterVertexComponent(ctx *VertexComponentContext) {}

// ExitVertexComponent is called when production vertexComponent is exited.
func (s *BaseLumenListener) ExitVertexComponent(ctx *VertexComponentContext) {}

// EnterShaderDecl is called when production shaderDecl is entered.
func (s *BaseLumenListener) EnterShaderDecl(ctx *ShaderDeclContext) {}

// ExitShaderDecl is called when production shaderDecl is exited.
func (s *BaseLumenListener) ExitShaderDecl(ctx *ShaderDeclContext) {}

// EnterVertexDecl is called when production vertexDecl is entered.
func (s *BaseLumenListener) EnterVertexDecl(ctx *VertexDeclContext) {}

// ExitVertexDecl is called when production vertexDecl is exited.
func (s *BaseLumenListener) ExitVertexDecl(ctx *VertexDeclContext) {}

// EnterFile is called when production file is entered.
func (s *BaseLumenListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseLumenListener) ExitFile(ctx *FileContext) {}
