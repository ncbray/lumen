// Generated from Lumen.g by ANTLR 4.6.

package parser // Lumen

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 40, 137, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 20, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 57, 10, 2, 12, 2, 14, 2, 60,
	11, 2, 5, 2, 62, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 68, 10, 2, 12, 2,
	14, 2, 71, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 87, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 7, 5, 95, 10, 5, 12, 5, 14, 5, 98, 11, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6,
	114, 10, 6, 12, 6, 14, 6, 117, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 123,
	10, 6, 12, 6, 14, 6, 126, 11, 6, 3, 6, 3, 6, 3, 6, 3, 7, 7, 7, 132, 10,
	7, 12, 7, 14, 7, 135, 11, 7, 3, 7, 2, 3, 2, 8, 2, 4, 6, 8, 10, 12, 2, 8,
	3, 2, 7, 10, 3, 2, 11, 13, 3, 2, 7, 8, 3, 2, 14, 15, 3, 2, 16, 19, 3, 2,
	20, 21, 152, 2, 19, 3, 2, 2, 2, 4, 86, 3, 2, 2, 2, 6, 88, 3, 2, 2, 2, 8,
	92, 3, 2, 2, 2, 10, 101, 3, 2, 2, 2, 12, 133, 3, 2, 2, 2, 14, 15, 8, 2,
	1, 2, 15, 16, 9, 2, 2, 2, 16, 20, 5, 2, 2, 15, 17, 20, 7, 38, 2, 2, 18,
	20, 7, 37, 2, 2, 19, 14, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 18, 3, 2,
	2, 2, 20, 69, 3, 2, 2, 2, 21, 22, 12, 14, 2, 2, 22, 23, 9, 3, 2, 2, 23,
	68, 5, 2, 2, 15, 24, 25, 12, 13, 2, 2, 25, 26, 9, 4, 2, 2, 26, 68, 5, 2,
	2, 14, 27, 28, 12, 12, 2, 2, 28, 29, 9, 5, 2, 2, 29, 68, 5, 2, 2, 13, 30,
	31, 12, 11, 2, 2, 31, 32, 9, 6, 2, 2, 32, 68, 5, 2, 2, 12, 33, 34, 12,
	10, 2, 2, 34, 35, 9, 7, 2, 2, 35, 68, 5, 2, 2, 11, 36, 37, 12, 9, 2, 2,
	37, 38, 7, 22, 2, 2, 38, 68, 5, 2, 2, 10, 39, 40, 12, 8, 2, 2, 40, 41,
	7, 23, 2, 2, 41, 68, 5, 2, 2, 9, 42, 43, 12, 7, 2, 2, 43, 44, 7, 24, 2,
	2, 44, 68, 5, 2, 2, 8, 45, 46, 12, 6, 2, 2, 46, 47, 7, 25, 2, 2, 47, 68,
	5, 2, 2, 7, 48, 49, 12, 5, 2, 2, 49, 50, 7, 26, 2, 2, 50, 68, 5, 2, 2,
	6, 51, 52, 12, 17, 2, 2, 52, 61, 7, 3, 2, 2, 53, 58, 5, 2, 2, 2, 54, 55,
	7, 4, 2, 2, 55, 57, 5, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2,
	58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3,
	2, 2, 2, 61, 53, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63,
	68, 7, 5, 2, 2, 64, 65, 12, 16, 2, 2, 65, 66, 7, 6, 2, 2, 66, 68, 7, 38,
	2, 2, 67, 21, 3, 2, 2, 2, 67, 24, 3, 2, 2, 2, 67, 27, 3, 2, 2, 2, 67, 30,
	3, 2, 2, 2, 67, 33, 3, 2, 2, 2, 67, 36, 3, 2, 2, 2, 67, 39, 3, 2, 2, 2,
	67, 42, 3, 2, 2, 2, 67, 45, 3, 2, 2, 2, 67, 48, 3, 2, 2, 2, 67, 51, 3,
	2, 2, 2, 67, 64, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69,
	70, 3, 2, 2, 2, 70, 3, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 7, 38, 2,
	2, 73, 74, 7, 38, 2, 2, 74, 75, 7, 27, 2, 2, 75, 76, 5, 2, 2, 2, 76, 77,
	7, 28, 2, 2, 77, 87, 3, 2, 2, 2, 78, 79, 7, 38, 2, 2, 79, 80, 7, 27, 2,
	2, 80, 81, 5, 2, 2, 2, 81, 82, 7, 28, 2, 2, 82, 87, 3, 2, 2, 2, 83, 84,
	5, 2, 2, 2, 84, 85, 7, 28, 2, 2, 85, 87, 3, 2, 2, 2, 86, 72, 3, 2, 2, 2,
	86, 78, 3, 2, 2, 2, 86, 83, 3, 2, 2, 2, 87, 5, 3, 2, 2, 2, 88, 89, 7, 38,
	2, 2, 89, 90, 7, 38, 2, 2, 90, 91, 7, 28, 2, 2, 91, 7, 3, 2, 2, 2, 92,
	96, 7, 29, 2, 2, 93, 95, 5, 6, 4, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2,
	2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96,
	3, 2, 2, 2, 99, 100, 7, 30, 2, 2, 100, 9, 3, 2, 2, 2, 101, 102, 7, 31,
	2, 2, 102, 103, 7, 38, 2, 2, 103, 104, 7, 29, 2, 2, 104, 105, 7, 32, 2,
	2, 105, 106, 5, 8, 5, 2, 106, 107, 7, 33, 2, 2, 107, 108, 5, 8, 5, 2, 108,
	109, 7, 34, 2, 2, 109, 110, 5, 8, 5, 2, 110, 111, 7, 35, 2, 2, 111, 115,
	7, 29, 2, 2, 112, 114, 5, 4, 3, 2, 113, 112, 3, 2, 2, 2, 114, 117, 3, 2,
	2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 3, 2, 2, 2,
	117, 115, 3, 2, 2, 2, 118, 119, 7, 30, 2, 2, 119, 120, 7, 36, 2, 2, 120,
	124, 7, 29, 2, 2, 121, 123, 5, 4, 3, 2, 122, 121, 3, 2, 2, 2, 123, 126,
	3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2,
	2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 7, 30, 2, 2, 128, 129, 7, 30, 2,
	2, 129, 11, 3, 2, 2, 2, 130, 132, 5, 10, 6, 2, 131, 130, 3, 2, 2, 2, 132,
	135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 13, 3,
	2, 2, 2, 135, 133, 3, 2, 2, 2, 12, 19, 58, 61, 67, 69, 86, 96, 115, 124,
	133,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'.'", "'+'", "'-'", "'~'", "'!'", "'*'", "'/'",
	"'%'", "'<<'", "'>>'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&'",
	"'^'", "'|'", "'&&'", "'||'", "'='", "';'", "'{'", "'}'", "'shader'", "'uniform'",
	"'attribute'", "'varying'", "'vs'", "'fs'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NUM",
	"ID", "SINGLE_LINE_COMMENT", "WS",
}

var ruleNames = []string{
	"expr", "statement", "field", "format", "shaderDecl", "file",
}

type LumenParser struct {
	*antlr.BaseParser
}

func NewLumenParser(input antlr.TokenStream) *LumenParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(LumenParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Lumen.g"

	return this
}

// LumenParser tokens.
const (
	LumenParserEOF                 = antlr.TokenEOF
	LumenParserT__0                = 1
	LumenParserT__1                = 2
	LumenParserT__2                = 3
	LumenParserT__3                = 4
	LumenParserT__4                = 5
	LumenParserT__5                = 6
	LumenParserT__6                = 7
	LumenParserT__7                = 8
	LumenParserT__8                = 9
	LumenParserT__9                = 10
	LumenParserT__10               = 11
	LumenParserT__11               = 12
	LumenParserT__12               = 13
	LumenParserT__13               = 14
	LumenParserT__14               = 15
	LumenParserT__15               = 16
	LumenParserT__16               = 17
	LumenParserT__17               = 18
	LumenParserT__18               = 19
	LumenParserT__19               = 20
	LumenParserT__20               = 21
	LumenParserT__21               = 22
	LumenParserT__22               = 23
	LumenParserT__23               = 24
	LumenParserT__24               = 25
	LumenParserT__25               = 26
	LumenParserT__26               = 27
	LumenParserT__27               = 28
	LumenParserT__28               = 29
	LumenParserT__29               = 30
	LumenParserT__30               = 31
	LumenParserT__31               = 32
	LumenParserT__32               = 33
	LumenParserT__33               = 34
	LumenParserNUM                 = 35
	LumenParserID                  = 36
	LumenParserSINGLE_LINE_COMMENT = 37
	LumenParserWS                  = 38
)

// LumenParser rules.
const (
	LumenParserRULE_expr       = 0
	LumenParserRULE_statement  = 1
	LumenParserRULE_field      = 2
	LumenParserRULE_format     = 3
	LumenParserRULE_shaderDecl = 4
	LumenParserRULE_file       = 5
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallContext struct {
	*ExprContext
	value IExprContext
	_expr IExprContext
	args  []IExprContext
}

func NewCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallContext {
	var p = new(CallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallContext) GetValue() IExprContext { return s.value }

func (s *CallContext) Get_expr() IExprContext { return s._expr }

func (s *CallContext) SetValue(v IExprContext) { s.value = v }

func (s *CallContext) Set_expr(v IExprContext) { s._expr = v }

func (s *CallContext) GetArgs() []IExprContext { return s.args }

func (s *CallContext) SetArgs(v []IExprContext) { s.args = v }

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CallContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitCall(s)
	}
}

type NumberContext struct {
	*ExprContext
	raw antlr.Token
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRaw() antlr.Token { return s.raw }

func (s *NumberContext) SetRaw(v antlr.Token) { s.raw = v }

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUM() antlr.TerminalNode {
	return s.GetToken(LumenParserNUM, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitNumber(s)
	}
}

type GetNameContext struct {
	*ExprContext
	name antlr.Token
}

func NewGetNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetNameContext {
	var p = new(GetNameContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetNameContext) GetName() antlr.Token { return s.name }

func (s *GetNameContext) SetName(v antlr.Token) { s.name = v }

func (s *GetNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetNameContext) ID() antlr.TerminalNode {
	return s.GetToken(LumenParserID, 0)
}

func (s *GetNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterGetName(s)
	}
}

func (s *GetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitGetName(s)
	}
}

type PrefixContext struct {
	*ExprContext
	op    antlr.Token
	value IExprContext
}

func NewPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixContext {
	var p = new(PrefixContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PrefixContext) GetOp() antlr.Token { return s.op }

func (s *PrefixContext) SetOp(v antlr.Token) { s.op = v }

func (s *PrefixContext) GetValue() IExprContext { return s.value }

func (s *PrefixContext) SetValue(v IExprContext) { s.value = v }

func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitPrefix(s)
	}
}

type GetAttrContext struct {
	*ExprContext
	value IExprContext
	name  antlr.Token
}

func NewGetAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetAttrContext {
	var p = new(GetAttrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetAttrContext) GetName() antlr.Token { return s.name }

func (s *GetAttrContext) SetName(v antlr.Token) { s.name = v }

func (s *GetAttrContext) GetValue() IExprContext { return s.value }

func (s *GetAttrContext) SetValue(v IExprContext) { s.value = v }

func (s *GetAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetAttrContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GetAttrContext) ID() antlr.TerminalNode {
	return s.GetToken(LumenParserID, 0)
}

func (s *GetAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterGetAttr(s)
	}
}

func (s *GetAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitGetAttr(s)
	}
}

type InfixContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewInfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InfixContext {
	var p = new(InfixContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InfixContext) GetOp() antlr.Token { return s.op }

func (s *InfixContext) SetOp(v antlr.Token) { s.op = v }

func (s *InfixContext) GetLeft() IExprContext { return s.left }

func (s *InfixContext) GetRight() IExprContext { return s.right }

func (s *InfixContext) SetLeft(v IExprContext) { s.left = v }

func (s *InfixContext) SetRight(v IExprContext) { s.right = v }

func (s *InfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfixContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *InfixContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterInfix(s)
	}
}

func (s *InfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitInfix(s)
	}
}

func (p *LumenParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LumenParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, LumenParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LumenParserT__4, LumenParserT__5, LumenParserT__6, LumenParserT__7:
		localctx = NewPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(13)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PrefixContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__4)|(1<<LumenParserT__5)|(1<<LumenParserT__6)|(1<<LumenParserT__7))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PrefixContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(14)

			var _x = p.expr(13)

			localctx.(*PrefixContext).value = _x
		}

	case LumenParserID:
		localctx = NewGetNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)

			var _m = p.Match(LumenParserID)

			localctx.(*GetNameContext).name = _m
		}

	case LumenParserNUM:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)

			var _m = p.Match(LumenParserNUM)

			localctx.(*NumberContext).raw = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(20)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__8)|(1<<LumenParserT__9)|(1<<LumenParserT__10))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(21)

					var _x = p.expr(13)

					localctx.(*InfixContext).right = _x
				}

			case 2:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(23)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == LumenParserT__4 || _la == LumenParserT__5) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(24)

					var _x = p.expr(12)

					localctx.(*InfixContext).right = _x
				}

			case 3:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(25)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(26)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == LumenParserT__11 || _la == LumenParserT__12) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(27)

					var _x = p.expr(11)

					localctx.(*InfixContext).right = _x
				}

			case 4:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(28)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(29)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__13)|(1<<LumenParserT__14)|(1<<LumenParserT__15)|(1<<LumenParserT__16))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(30)

					var _x = p.expr(10)

					localctx.(*InfixContext).right = _x
				}

			case 5:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(31)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				p.SetState(32)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == LumenParserT__17 || _la == LumenParserT__18) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(33)

					var _x = p.expr(9)

					localctx.(*InfixContext).right = _x
				}

			case 6:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(35)

					var _m = p.Match(LumenParserT__19)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(36)

					var _x = p.expr(8)

					localctx.(*InfixContext).right = _x
				}

			case 7:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(38)

					var _m = p.Match(LumenParserT__20)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(39)

					var _x = p.expr(7)

					localctx.(*InfixContext).right = _x
				}

			case 8:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(41)

					var _m = p.Match(LumenParserT__21)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(42)

					var _x = p.expr(6)

					localctx.(*InfixContext).right = _x
				}

			case 9:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(44)

					var _m = p.Match(LumenParserT__22)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(45)

					var _x = p.expr(5)

					localctx.(*InfixContext).right = _x
				}

			case 10:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(47)

					var _m = p.Match(LumenParserT__23)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(48)

					var _x = p.expr(4)

					localctx.(*InfixContext).right = _x
				}

			case 11:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallContext).value = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(50)
					p.Match(LumenParserT__0)
				}
				p.SetState(59)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(LumenParserT__4-5))|(1<<(LumenParserT__5-5))|(1<<(LumenParserT__6-5))|(1<<(LumenParserT__7-5))|(1<<(LumenParserNUM-5))|(1<<(LumenParserID-5)))) != 0 {
					{
						p.SetState(51)

						var _x = p.expr(0)

						localctx.(*CallContext)._expr = _x
					}
					localctx.(*CallContext).args = append(localctx.(*CallContext).args, localctx.(*CallContext)._expr)
					p.SetState(56)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LumenParserT__1 {
						{
							p.SetState(52)
							p.Match(LumenParserT__1)
						}
						{
							p.SetState(53)

							var _x = p.expr(0)

							localctx.(*CallContext)._expr = _x
						}
						localctx.(*CallContext).args = append(localctx.(*CallContext).args, localctx.(*CallContext)._expr)

						p.SetState(58)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(61)
					p.Match(LumenParserT__2)
				}

			case 12:
				localctx = NewGetAttrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GetAttrContext).value = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(63)
					p.Match(LumenParserT__3)
				}
				{
					p.SetState(64)

					var _m = p.Match(LumenParserID)

					localctx.(*GetAttrContext).name = _m
				}

			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DiscardContext struct {
	*StatementContext
	value IExprContext
}

func NewDiscardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DiscardContext {
	var p = new(DiscardContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *DiscardContext) GetValue() IExprContext { return s.value }

func (s *DiscardContext) SetValue(v IExprContext) { s.value = v }

func (s *DiscardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiscardContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DiscardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterDiscard(s)
	}
}

func (s *DiscardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitDiscard(s)
	}
}

type VarDeclContext struct {
	*StatementContext
	t     antlr.Token
	name  antlr.Token
	value IExprContext
}

func NewVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclContext {
	var p = new(VarDeclContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *VarDeclContext) GetT() antlr.Token { return s.t }

func (s *VarDeclContext) GetName() antlr.Token { return s.name }

func (s *VarDeclContext) SetT(v antlr.Token) { s.t = v }

func (s *VarDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *VarDeclContext) GetValue() IExprContext { return s.value }

func (s *VarDeclContext) SetValue(v IExprContext) { s.value = v }

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LumenParserID)
}

func (s *VarDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LumenParserID, i)
}

func (s *VarDeclContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

type AssignContext struct {
	*StatementContext
	name  antlr.Token
	value IExprContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignContext) GetName() antlr.Token { return s.name }

func (s *AssignContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignContext) GetValue() IExprContext { return s.value }

func (s *AssignContext) SetValue(v IExprContext) { s.value = v }

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(LumenParserID, 0)
}

func (s *AssignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *LumenParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LumenParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)

			var _m = p.Match(LumenParserID)

			localctx.(*VarDeclContext).t = _m
		}
		{
			p.SetState(71)

			var _m = p.Match(LumenParserID)

			localctx.(*VarDeclContext).name = _m
		}
		{
			p.SetState(72)
			p.Match(LumenParserT__24)
		}
		{
			p.SetState(73)

			var _x = p.expr(0)

			localctx.(*VarDeclContext).value = _x
		}
		{
			p.SetState(74)
			p.Match(LumenParserT__25)
		}

	case 2:
		localctx = NewAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)

			var _m = p.Match(LumenParserID)

			localctx.(*AssignContext).name = _m
		}
		{
			p.SetState(77)
			p.Match(LumenParserT__24)
		}
		{
			p.SetState(78)

			var _x = p.expr(0)

			localctx.(*AssignContext).value = _x
		}
		{
			p.SetState(79)
			p.Match(LumenParserT__25)
		}

	case 3:
		localctx = NewDiscardContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)

			var _x = p.expr(0)

			localctx.(*DiscardContext).value = _x
		}
		{
			p.SetState(82)
			p.Match(LumenParserT__25)
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetT returns the t token.
	GetT() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetT sets the t token.
	SetT(antlr.Token)

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	t      antlr.Token
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) GetName() antlr.Token { return s.name }

func (s *FieldContext) GetT() antlr.Token { return s.t }

func (s *FieldContext) SetName(v antlr.Token) { s.name = v }

func (s *FieldContext) SetT(v antlr.Token) { s.t = v }

func (s *FieldContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LumenParserID)
}

func (s *FieldContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LumenParserID, i)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *LumenParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LumenParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)

		var _m = p.Match(LumenParserID)

		localctx.(*FieldContext).name = _m
	}
	{
		p.SetState(87)

		var _m = p.Match(LumenParserID)

		localctx.(*FieldContext).t = _m
	}
	{
		p.SetState(88)
		p.Match(LumenParserT__25)
	}

	return localctx
}

// IFormatContext is an interface to support dynamic dispatch.
type IFormatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_field returns the _field rule contexts.
	Get_field() IFieldContext

	// Set_field sets the _field rule contexts.
	Set_field(IFieldContext)

	// GetFields returns the fields rule context list.
	GetFields() []IFieldContext

	// SetFields sets the fields rule context list.
	SetFields([]IFieldContext)

	// IsFormatContext differentiates from other interfaces.
	IsFormatContext()
}

type FormatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_field IFieldContext
	fields []IFieldContext
}

func NewEmptyFormatContext() *FormatContext {
	var p = new(FormatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_format
	return p
}

func (*FormatContext) IsFormatContext() {}

func NewFormatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatContext {
	var p = new(FormatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_format

	return p
}

func (s *FormatContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatContext) Get_field() IFieldContext { return s._field }

func (s *FormatContext) Set_field(v IFieldContext) { s._field = v }

func (s *FormatContext) GetFields() []IFieldContext { return s.fields }

func (s *FormatContext) SetFields(v []IFieldContext) { s.fields = v }

func (s *FormatContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *FormatContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FormatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterFormat(s)
	}
}

func (s *FormatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitFormat(s)
	}
}

func (p *LumenParser) Format() (localctx IFormatContext) {
	localctx = NewFormatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LumenParserRULE_format)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(LumenParserT__26)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LumenParserID {
		{
			p.SetState(91)

			var _x = p.Field()

			localctx.(*FormatContext)._field = _x
		}
		localctx.(*FormatContext).fields = append(localctx.(*FormatContext).fields, localctx.(*FormatContext)._field)

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Match(LumenParserT__27)
	}

	return localctx
}

// IShaderDeclContext is an interface to support dynamic dispatch.
type IShaderDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetUniform returns the uniform rule contexts.
	GetUniform() IFormatContext

	// GetAttribute returns the attribute rule contexts.
	GetAttribute() IFormatContext

	// GetVarying returns the varying rule contexts.
	GetVarying() IFormatContext

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// SetUniform sets the uniform rule contexts.
	SetUniform(IFormatContext)

	// SetAttribute sets the attribute rule contexts.
	SetAttribute(IFormatContext)

	// SetVarying sets the varying rule contexts.
	SetVarying(IFormatContext)

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetVs returns the vs rule context list.
	GetVs() []IStatementContext

	// GetFs returns the fs rule context list.
	GetFs() []IStatementContext

	// SetVs sets the vs rule context list.
	SetVs([]IStatementContext)

	// SetFs sets the fs rule context list.
	SetFs([]IStatementContext)

	// IsShaderDeclContext differentiates from other interfaces.
	IsShaderDeclContext()
}

type ShaderDeclContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	name       antlr.Token
	uniform    IFormatContext
	attribute  IFormatContext
	varying    IFormatContext
	_statement IStatementContext
	vs         []IStatementContext
	fs         []IStatementContext
}

func NewEmptyShaderDeclContext() *ShaderDeclContext {
	var p = new(ShaderDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_shaderDecl
	return p
}

func (*ShaderDeclContext) IsShaderDeclContext() {}

func NewShaderDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShaderDeclContext {
	var p = new(ShaderDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_shaderDecl

	return p
}

func (s *ShaderDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ShaderDeclContext) GetName() antlr.Token { return s.name }

func (s *ShaderDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *ShaderDeclContext) GetUniform() IFormatContext { return s.uniform }

func (s *ShaderDeclContext) GetAttribute() IFormatContext { return s.attribute }

func (s *ShaderDeclContext) GetVarying() IFormatContext { return s.varying }

func (s *ShaderDeclContext) Get_statement() IStatementContext { return s._statement }

func (s *ShaderDeclContext) SetUniform(v IFormatContext) { s.uniform = v }

func (s *ShaderDeclContext) SetAttribute(v IFormatContext) { s.attribute = v }

func (s *ShaderDeclContext) SetVarying(v IFormatContext) { s.varying = v }

func (s *ShaderDeclContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *ShaderDeclContext) GetVs() []IStatementContext { return s.vs }

func (s *ShaderDeclContext) GetFs() []IStatementContext { return s.fs }

func (s *ShaderDeclContext) SetVs(v []IStatementContext) { s.vs = v }

func (s *ShaderDeclContext) SetFs(v []IStatementContext) { s.fs = v }

func (s *ShaderDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(LumenParserID, 0)
}

func (s *ShaderDeclContext) AllFormat() []IFormatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormatContext)(nil)).Elem())
	var tst = make([]IFormatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormatContext)
		}
	}

	return tst
}

func (s *ShaderDeclContext) Format(i int) IFormatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormatContext)
}

func (s *ShaderDeclContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ShaderDeclContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ShaderDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShaderDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShaderDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterShaderDecl(s)
	}
}

func (s *ShaderDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitShaderDecl(s)
	}
}

func (p *LumenParser) ShaderDecl() (localctx IShaderDeclContext) {
	localctx = NewShaderDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LumenParserRULE_shaderDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(LumenParserT__28)
	}
	{
		p.SetState(100)

		var _m = p.Match(LumenParserID)

		localctx.(*ShaderDeclContext).name = _m
	}
	{
		p.SetState(101)
		p.Match(LumenParserT__26)
	}
	{
		p.SetState(102)
		p.Match(LumenParserT__29)
	}
	{
		p.SetState(103)

		var _x = p.Format()

		localctx.(*ShaderDeclContext).uniform = _x
	}
	{
		p.SetState(104)
		p.Match(LumenParserT__30)
	}
	{
		p.SetState(105)

		var _x = p.Format()

		localctx.(*ShaderDeclContext).attribute = _x
	}
	{
		p.SetState(106)
		p.Match(LumenParserT__31)
	}
	{
		p.SetState(107)

		var _x = p.Format()

		localctx.(*ShaderDeclContext).varying = _x
	}
	{
		p.SetState(108)
		p.Match(LumenParserT__32)
	}
	{
		p.SetState(109)
		p.Match(LumenParserT__26)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(LumenParserT__4-5))|(1<<(LumenParserT__5-5))|(1<<(LumenParserT__6-5))|(1<<(LumenParserT__7-5))|(1<<(LumenParserNUM-5))|(1<<(LumenParserID-5)))) != 0 {
		{
			p.SetState(110)

			var _x = p.Statement()

			localctx.(*ShaderDeclContext)._statement = _x
		}
		localctx.(*ShaderDeclContext).vs = append(localctx.(*ShaderDeclContext).vs, localctx.(*ShaderDeclContext)._statement)

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
		p.Match(LumenParserT__27)
	}
	{
		p.SetState(117)
		p.Match(LumenParserT__33)
	}
	{
		p.SetState(118)
		p.Match(LumenParserT__26)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(LumenParserT__4-5))|(1<<(LumenParserT__5-5))|(1<<(LumenParserT__6-5))|(1<<(LumenParserT__7-5))|(1<<(LumenParserNUM-5))|(1<<(LumenParserID-5)))) != 0 {
		{
			p.SetState(119)

			var _x = p.Statement()

			localctx.(*ShaderDeclContext)._statement = _x
		}
		localctx.(*ShaderDeclContext).fs = append(localctx.(*ShaderDeclContext).fs, localctx.(*ShaderDeclContext)._statement)

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(LumenParserT__27)
	}
	{
		p.SetState(126)
		p.Match(LumenParserT__27)
	}

	return localctx
}

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_shaderDecl returns the _shaderDecl rule contexts.
	Get_shaderDecl() IShaderDeclContext

	// Set_shaderDecl sets the _shaderDecl rule contexts.
	Set_shaderDecl(IShaderDeclContext)

	// GetShaders returns the shaders rule context list.
	GetShaders() []IShaderDeclContext

	// SetShaders sets the shaders rule context list.
	SetShaders([]IShaderDeclContext)

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_shaderDecl IShaderDeclContext
	shaders     []IShaderDeclContext
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Get_shaderDecl() IShaderDeclContext { return s._shaderDecl }

func (s *FileContext) Set_shaderDecl(v IShaderDeclContext) { s._shaderDecl = v }

func (s *FileContext) GetShaders() []IShaderDeclContext { return s.shaders }

func (s *FileContext) SetShaders(v []IShaderDeclContext) { s.shaders = v }

func (s *FileContext) AllShaderDecl() []IShaderDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShaderDeclContext)(nil)).Elem())
	var tst = make([]IShaderDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShaderDeclContext)
		}
	}

	return tst
}

func (s *FileContext) ShaderDecl(i int) IShaderDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShaderDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShaderDeclContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *LumenParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LumenParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LumenParserT__28 {
		{
			p.SetState(128)

			var _x = p.ShaderDecl()

			localctx.(*FileContext)._shaderDecl = _x
		}
		localctx.(*FileContext).shaders = append(localctx.(*FileContext).shaders, localctx.(*FileContext)._shaderDecl)

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *LumenParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LumenParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
