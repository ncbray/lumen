// Generated from Lumen.g by ANTLR 4.6.

package parser // Lumen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LumenListener is a complete listener for a parse tree produced by LumenParser.
type LumenListener interface {
	antlr.ParseTreeListener

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterGetName is called when entering the getName production.
	EnterGetName(c *GetNameContext)

	// EnterPrefix is called when entering the prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterGetAttr is called when entering the getAttr production.
	EnterGetAttr(c *GetAttrContext)

	// EnterInfix is called when entering the infix production.
	EnterInfix(c *InfixContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterDiscard is called when entering the discard production.
	EnterDiscard(c *DiscardContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFormat is called when entering the format production.
	EnterFormat(c *FormatContext)

	// EnterShaderDecl is called when entering the shaderDecl production.
	EnterShaderDecl(c *ShaderDeclContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitGetName is called when exiting the getName production.
	ExitGetName(c *GetNameContext)

	// ExitPrefix is called when exiting the prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitGetAttr is called when exiting the getAttr production.
	ExitGetAttr(c *GetAttrContext)

	// ExitInfix is called when exiting the infix production.
	ExitInfix(c *InfixContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitDiscard is called when exiting the discard production.
	ExitDiscard(c *DiscardContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFormat is called when exiting the format production.
	ExitFormat(c *FormatContext)

	// ExitShaderDecl is called when exiting the shaderDecl production.
	ExitShaderDecl(c *ShaderDeclContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)
}
