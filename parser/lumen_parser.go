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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 41, 156, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 22, 10, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 59, 10, 2, 12,
	2, 14, 2, 62, 11, 2, 5, 2, 64, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 70,
	10, 2, 12, 2, 14, 2, 73, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 89, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 97, 10, 5, 12, 5, 14, 5, 100, 11, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 121, 10, 7, 12, 7, 14, 7,
	124, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 130, 10, 7, 12, 7, 14, 7, 133,
	11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 142, 10, 7, 12,
	7, 14, 7, 145, 11, 7, 3, 7, 5, 7, 148, 10, 7, 3, 8, 7, 8, 151, 10, 8, 12,
	8, 14, 8, 154, 11, 8, 3, 8, 2, 3, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 8, 3,
	2, 7, 10, 3, 2, 11, 13, 3, 2, 7, 8, 3, 2, 14, 15, 3, 2, 16, 19, 3, 2, 20,
	21, 172, 2, 21, 3, 2, 2, 2, 4, 88, 3, 2, 2, 2, 6, 90, 3, 2, 2, 2, 8, 94,
	3, 2, 2, 2, 10, 103, 3, 2, 2, 2, 12, 147, 3, 2, 2, 2, 14, 152, 3, 2, 2,
	2, 16, 17, 8, 2, 1, 2, 17, 18, 9, 2, 2, 2, 18, 22, 5, 2, 2, 15, 19, 22,
	7, 39, 2, 2, 20, 22, 7, 38, 2, 2, 21, 16, 3, 2, 2, 2, 21, 19, 3, 2, 2,
	2, 21, 20, 3, 2, 2, 2, 22, 71, 3, 2, 2, 2, 23, 24, 12, 14, 2, 2, 24, 25,
	9, 3, 2, 2, 25, 70, 5, 2, 2, 15, 26, 27, 12, 13, 2, 2, 27, 28, 9, 4, 2,
	2, 28, 70, 5, 2, 2, 14, 29, 30, 12, 12, 2, 2, 30, 31, 9, 5, 2, 2, 31, 70,
	5, 2, 2, 13, 32, 33, 12, 11, 2, 2, 33, 34, 9, 6, 2, 2, 34, 70, 5, 2, 2,
	12, 35, 36, 12, 10, 2, 2, 36, 37, 9, 7, 2, 2, 37, 70, 5, 2, 2, 11, 38,
	39, 12, 9, 2, 2, 39, 40, 7, 22, 2, 2, 40, 70, 5, 2, 2, 10, 41, 42, 12,
	8, 2, 2, 42, 43, 7, 23, 2, 2, 43, 70, 5, 2, 2, 9, 44, 45, 12, 7, 2, 2,
	45, 46, 7, 24, 2, 2, 46, 70, 5, 2, 2, 8, 47, 48, 12, 6, 2, 2, 48, 49, 7,
	25, 2, 2, 49, 70, 5, 2, 2, 7, 50, 51, 12, 5, 2, 2, 51, 52, 7, 26, 2, 2,
	52, 70, 5, 2, 2, 6, 53, 54, 12, 17, 2, 2, 54, 63, 7, 3, 2, 2, 55, 60, 5,
	2, 2, 2, 56, 57, 7, 4, 2, 2, 57, 59, 5, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59,
	62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 64, 3, 2, 2,
	2, 62, 60, 3, 2, 2, 2, 63, 55, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65,
	3, 2, 2, 2, 65, 70, 7, 5, 2, 2, 66, 67, 12, 16, 2, 2, 67, 68, 7, 6, 2,
	2, 68, 70, 7, 39, 2, 2, 69, 23, 3, 2, 2, 2, 69, 26, 3, 2, 2, 2, 69, 29,
	3, 2, 2, 2, 69, 32, 3, 2, 2, 2, 69, 35, 3, 2, 2, 2, 69, 38, 3, 2, 2, 2,
	69, 41, 3, 2, 2, 2, 69, 44, 3, 2, 2, 2, 69, 47, 3, 2, 2, 2, 69, 50, 3,
	2, 2, 2, 69, 53, 3, 2, 2, 2, 69, 66, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71,
	69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 3, 3, 2, 2, 2, 73, 71, 3, 2, 2,
	2, 74, 75, 7, 39, 2, 2, 75, 76, 7, 39, 2, 2, 76, 77, 7, 27, 2, 2, 77, 78,
	5, 2, 2, 2, 78, 79, 7, 28, 2, 2, 79, 89, 3, 2, 2, 2, 80, 81, 7, 39, 2,
	2, 81, 82, 7, 27, 2, 2, 82, 83, 5, 2, 2, 2, 83, 84, 7, 28, 2, 2, 84, 89,
	3, 2, 2, 2, 85, 86, 5, 2, 2, 2, 86, 87, 7, 28, 2, 2, 87, 89, 3, 2, 2, 2,
	88, 74, 3, 2, 2, 2, 88, 80, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 89, 5, 3, 2,
	2, 2, 90, 91, 7, 39, 2, 2, 91, 92, 7, 39, 2, 2, 92, 93, 7, 28, 2, 2, 93,
	7, 3, 2, 2, 2, 94, 98, 7, 29, 2, 2, 95, 97, 5, 6, 4, 2, 96, 95, 3, 2, 2,
	2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 101,
	3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 102, 7, 30, 2, 2, 102, 9, 3, 2, 2,
	2, 103, 104, 7, 39, 2, 2, 104, 105, 7, 39, 2, 2, 105, 106, 7, 39, 2, 2,
	106, 107, 7, 28, 2, 2, 107, 11, 3, 2, 2, 2, 108, 109, 7, 31, 2, 2, 109,
	110, 7, 39, 2, 2, 110, 111, 7, 29, 2, 2, 111, 112, 7, 32, 2, 2, 112, 113,
	5, 8, 5, 2, 113, 114, 7, 33, 2, 2, 114, 115, 5, 8, 5, 2, 115, 116, 7, 34,
	2, 2, 116, 117, 5, 8, 5, 2, 117, 118, 7, 35, 2, 2, 118, 122, 7, 29, 2,
	2, 119, 121, 5, 4, 3, 2, 120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122,
	120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 125, 3, 2, 2, 2, 124, 122,
	3, 2, 2, 2, 125, 126, 7, 30, 2, 2, 126, 127, 7, 36, 2, 2, 127, 131, 7,
	29, 2, 2, 128, 130, 5, 4, 3, 2, 129, 128, 3, 2, 2, 2, 130, 133, 3, 2, 2,
	2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 3, 2, 2, 2, 133,
	131, 3, 2, 2, 2, 134, 135, 7, 30, 2, 2, 135, 136, 7, 30, 2, 2, 136, 148,
	3, 2, 2, 2, 137, 138, 7, 37, 2, 2, 138, 139, 7, 39, 2, 2, 139, 143, 7,
	29, 2, 2, 140, 142, 5, 10, 6, 2, 141, 140, 3, 2, 2, 2, 142, 145, 3, 2,
	2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 146, 3, 2, 2, 2,
	145, 143, 3, 2, 2, 2, 146, 148, 7, 30, 2, 2, 147, 108, 3, 2, 2, 2, 147,
	137, 3, 2, 2, 2, 148, 13, 3, 2, 2, 2, 149, 151, 5, 12, 7, 2, 150, 149,
	3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2,
	2, 2, 153, 15, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 14, 21, 60, 63, 69, 71,
	88, 98, 122, 131, 143, 147, 152,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'.'", "'+'", "'-'", "'~'", "'!'", "'*'", "'/'",
	"'%'", "'<<'", "'>>'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&'",
	"'^'", "'|'", "'&&'", "'||'", "'='", "';'", "'{'", "'}'", "'shader'", "'uniform'",
	"'attribute'", "'varying'", "'vs'", "'fs'", "'vertex'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"NUM", "ID", "SINGLE_LINE_COMMENT", "WS",
}

var ruleNames = []string{
	"expr", "statement", "field", "format", "vertexComponent", "topLevelDecl",
	"file",
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
	LumenParserT__34               = 35
	LumenParserNUM                 = 36
	LumenParserID                  = 37
	LumenParserSINGLE_LINE_COMMENT = 38
	LumenParserWS                  = 39
)

// LumenParser rules.
const (
	LumenParserRULE_expr            = 0
	LumenParserRULE_statement       = 1
	LumenParserRULE_field           = 2
	LumenParserRULE_format          = 3
	LumenParserRULE_vertexComponent = 4
	LumenParserRULE_topLevelDecl    = 5
	LumenParserRULE_file            = 6
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LumenParserT__4, LumenParserT__5, LumenParserT__6, LumenParserT__7:
		localctx = NewPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(15)

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
			p.SetState(16)

			var _x = p.expr(13)

			localctx.(*PrefixContext).value = _x
		}

	case LumenParserID:
		localctx = NewGetNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)

			var _m = p.Match(LumenParserID)

			localctx.(*GetNameContext).name = _m
		}

	case LumenParserNUM:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)

			var _m = p.Match(LumenParserNUM)

			localctx.(*NumberContext).raw = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(22)

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
					p.SetState(23)

					var _x = p.expr(13)

					localctx.(*InfixContext).right = _x
				}

			case 2:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(25)

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
					p.SetState(26)

					var _x = p.expr(12)

					localctx.(*InfixContext).right = _x
				}

			case 3:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(28)

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
					p.SetState(29)

					var _x = p.expr(11)

					localctx.(*InfixContext).right = _x
				}

			case 4:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(31)

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
					p.SetState(32)

					var _x = p.expr(10)

					localctx.(*InfixContext).right = _x
				}

			case 5:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				p.SetState(34)

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
					p.SetState(35)

					var _x = p.expr(9)

					localctx.(*InfixContext).right = _x
				}

			case 6:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(37)

					var _m = p.Match(LumenParserT__19)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(38)

					var _x = p.expr(8)

					localctx.(*InfixContext).right = _x
				}

			case 7:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(40)

					var _m = p.Match(LumenParserT__20)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(41)

					var _x = p.expr(7)

					localctx.(*InfixContext).right = _x
				}

			case 8:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(43)

					var _m = p.Match(LumenParserT__21)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(44)

					var _x = p.expr(6)

					localctx.(*InfixContext).right = _x
				}

			case 9:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(46)

					var _m = p.Match(LumenParserT__22)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(47)

					var _x = p.expr(5)

					localctx.(*InfixContext).right = _x
				}

			case 10:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(49)

					var _m = p.Match(LumenParserT__23)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(50)

					var _x = p.expr(4)

					localctx.(*InfixContext).right = _x
				}

			case 11:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallContext).value = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(52)
					p.Match(LumenParserT__0)
				}
				p.SetState(61)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__4)|(1<<LumenParserT__5)|(1<<LumenParserT__6)|(1<<LumenParserT__7))) != 0) || _la == LumenParserNUM || _la == LumenParserID {
					{
						p.SetState(53)

						var _x = p.expr(0)

						localctx.(*CallContext)._expr = _x
					}
					localctx.(*CallContext).args = append(localctx.(*CallContext).args, localctx.(*CallContext)._expr)
					p.SetState(58)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LumenParserT__1 {
						{
							p.SetState(54)
							p.Match(LumenParserT__1)
						}
						{
							p.SetState(55)

							var _x = p.expr(0)

							localctx.(*CallContext)._expr = _x
						}
						localctx.(*CallContext).args = append(localctx.(*CallContext).args, localctx.(*CallContext)._expr)

						p.SetState(60)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(63)
					p.Match(LumenParserT__2)
				}

			case 12:
				localctx = NewGetAttrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GetAttrContext).value = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(65)
					p.Match(LumenParserT__3)
				}
				{
					p.SetState(66)

					var _m = p.Match(LumenParserID)

					localctx.(*GetAttrContext).name = _m
				}

			}

		}
		p.SetState(71)
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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)

			var _m = p.Match(LumenParserID)

			localctx.(*VarDeclContext).t = _m
		}
		{
			p.SetState(73)

			var _m = p.Match(LumenParserID)

			localctx.(*VarDeclContext).name = _m
		}
		{
			p.SetState(74)
			p.Match(LumenParserT__24)
		}
		{
			p.SetState(75)

			var _x = p.expr(0)

			localctx.(*VarDeclContext).value = _x
		}
		{
			p.SetState(76)
			p.Match(LumenParserT__25)
		}

	case 2:
		localctx = NewAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)

			var _m = p.Match(LumenParserID)

			localctx.(*AssignContext).name = _m
		}
		{
			p.SetState(79)
			p.Match(LumenParserT__24)
		}
		{
			p.SetState(80)

			var _x = p.expr(0)

			localctx.(*AssignContext).value = _x
		}
		{
			p.SetState(81)
			p.Match(LumenParserT__25)
		}

	case 3:
		localctx = NewDiscardContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)

			var _x = p.expr(0)

			localctx.(*DiscardContext).value = _x
		}
		{
			p.SetState(84)
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
		p.SetState(88)

		var _m = p.Match(LumenParserID)

		localctx.(*FieldContext).name = _m
	}
	{
		p.SetState(89)

		var _m = p.Match(LumenParserID)

		localctx.(*FieldContext).t = _m
	}
	{
		p.SetState(90)
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
		p.SetState(92)
		p.Match(LumenParserT__26)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LumenParserID {
		{
			p.SetState(93)

			var _x = p.Field()

			localctx.(*FormatContext)._field = _x
		}
		localctx.(*FormatContext).fields = append(localctx.(*FormatContext).fields, localctx.(*FormatContext)._field)

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(LumenParserT__27)
	}

	return localctx
}

// IVertexComponentContext is an interface to support dynamic dispatch.
type IVertexComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetT returns the t token.
	GetT() antlr.Token

	// GetEncoding returns the encoding token.
	GetEncoding() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetT sets the t token.
	SetT(antlr.Token)

	// SetEncoding sets the encoding token.
	SetEncoding(antlr.Token)

	// IsVertexComponentContext differentiates from other interfaces.
	IsVertexComponentContext()
}

type VertexComponentContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	name     antlr.Token
	t        antlr.Token
	encoding antlr.Token
}

func NewEmptyVertexComponentContext() *VertexComponentContext {
	var p = new(VertexComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_vertexComponent
	return p
}

func (*VertexComponentContext) IsVertexComponentContext() {}

func NewVertexComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VertexComponentContext {
	var p = new(VertexComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_vertexComponent

	return p
}

func (s *VertexComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *VertexComponentContext) GetName() antlr.Token { return s.name }

func (s *VertexComponentContext) GetT() antlr.Token { return s.t }

func (s *VertexComponentContext) GetEncoding() antlr.Token { return s.encoding }

func (s *VertexComponentContext) SetName(v antlr.Token) { s.name = v }

func (s *VertexComponentContext) SetT(v antlr.Token) { s.t = v }

func (s *VertexComponentContext) SetEncoding(v antlr.Token) { s.encoding = v }

func (s *VertexComponentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LumenParserID)
}

func (s *VertexComponentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LumenParserID, i)
}

func (s *VertexComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertexComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VertexComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterVertexComponent(s)
	}
}

func (s *VertexComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitVertexComponent(s)
	}
}

func (p *LumenParser) VertexComponent() (localctx IVertexComponentContext) {
	localctx = NewVertexComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LumenParserRULE_vertexComponent)

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
		p.SetState(101)

		var _m = p.Match(LumenParserID)

		localctx.(*VertexComponentContext).name = _m
	}
	{
		p.SetState(102)

		var _m = p.Match(LumenParserID)

		localctx.(*VertexComponentContext).t = _m
	}
	{
		p.SetState(103)

		var _m = p.Match(LumenParserID)

		localctx.(*VertexComponentContext).encoding = _m
	}
	{
		p.SetState(104)
		p.Match(LumenParserT__25)
	}

	return localctx
}

// ITopLevelDeclContext is an interface to support dynamic dispatch.
type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}

type TopLevelDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LumenParserRULE_topLevelDecl
	return p
}

func (*TopLevelDeclContext) IsTopLevelDeclContext() {}

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LumenParserRULE_topLevelDecl

	return p
}

func (s *TopLevelDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDeclContext) CopyFrom(ctx *TopLevelDeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ShaderDeclContext struct {
	*TopLevelDeclContext
	name       antlr.Token
	uniform    IFormatContext
	attribute  IFormatContext
	varying    IFormatContext
	_statement IStatementContext
	vs         []IStatementContext
	fs         []IStatementContext
}

func NewShaderDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShaderDeclContext {
	var p = new(ShaderDeclContext)

	p.TopLevelDeclContext = NewEmptyTopLevelDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopLevelDeclContext))

	return p
}

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

func (s *ShaderDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

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

type VertexDeclContext struct {
	*TopLevelDeclContext
	name             antlr.Token
	_vertexComponent IVertexComponentContext
	components       []IVertexComponentContext
}

func NewVertexDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VertexDeclContext {
	var p = new(VertexDeclContext)

	p.TopLevelDeclContext = NewEmptyTopLevelDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopLevelDeclContext))

	return p
}

func (s *VertexDeclContext) GetName() antlr.Token { return s.name }

func (s *VertexDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *VertexDeclContext) Get_vertexComponent() IVertexComponentContext { return s._vertexComponent }

func (s *VertexDeclContext) Set_vertexComponent(v IVertexComponentContext) { s._vertexComponent = v }

func (s *VertexDeclContext) GetComponents() []IVertexComponentContext { return s.components }

func (s *VertexDeclContext) SetComponents(v []IVertexComponentContext) { s.components = v }

func (s *VertexDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertexDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(LumenParserID, 0)
}

func (s *VertexDeclContext) AllVertexComponent() []IVertexComponentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVertexComponentContext)(nil)).Elem())
	var tst = make([]IVertexComponentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVertexComponentContext)
		}
	}

	return tst
}

func (s *VertexDeclContext) VertexComponent(i int) IVertexComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertexComponentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVertexComponentContext)
}

func (s *VertexDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.EnterVertexDecl(s)
	}
}

func (s *VertexDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LumenListener); ok {
		listenerT.ExitVertexDecl(s)
	}
}

func (p *LumenParser) TopLevelDecl() (localctx ITopLevelDeclContext) {
	localctx = NewTopLevelDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LumenParserRULE_topLevelDecl)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LumenParserT__28:
		localctx = NewShaderDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(LumenParserT__28)
		}
		{
			p.SetState(107)

			var _m = p.Match(LumenParserID)

			localctx.(*ShaderDeclContext).name = _m
		}
		{
			p.SetState(108)
			p.Match(LumenParserT__26)
		}
		{
			p.SetState(109)
			p.Match(LumenParserT__29)
		}
		{
			p.SetState(110)

			var _x = p.Format()

			localctx.(*ShaderDeclContext).uniform = _x
		}
		{
			p.SetState(111)
			p.Match(LumenParserT__30)
		}
		{
			p.SetState(112)

			var _x = p.Format()

			localctx.(*ShaderDeclContext).attribute = _x
		}
		{
			p.SetState(113)
			p.Match(LumenParserT__31)
		}
		{
			p.SetState(114)

			var _x = p.Format()

			localctx.(*ShaderDeclContext).varying = _x
		}
		{
			p.SetState(115)
			p.Match(LumenParserT__32)
		}
		{
			p.SetState(116)
			p.Match(LumenParserT__26)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__4)|(1<<LumenParserT__5)|(1<<LumenParserT__6)|(1<<LumenParserT__7))) != 0) || _la == LumenParserNUM || _la == LumenParserID {
			{
				p.SetState(117)

				var _x = p.Statement()

				localctx.(*ShaderDeclContext)._statement = _x
			}
			localctx.(*ShaderDeclContext).vs = append(localctx.(*ShaderDeclContext).vs, localctx.(*ShaderDeclContext)._statement)

			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(123)
			p.Match(LumenParserT__27)
		}
		{
			p.SetState(124)
			p.Match(LumenParserT__33)
		}
		{
			p.SetState(125)
			p.Match(LumenParserT__26)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__4)|(1<<LumenParserT__5)|(1<<LumenParserT__6)|(1<<LumenParserT__7))) != 0) || _la == LumenParserNUM || _la == LumenParserID {
			{
				p.SetState(126)

				var _x = p.Statement()

				localctx.(*ShaderDeclContext)._statement = _x
			}
			localctx.(*ShaderDeclContext).fs = append(localctx.(*ShaderDeclContext).fs, localctx.(*ShaderDeclContext)._statement)

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(132)
			p.Match(LumenParserT__27)
		}
		{
			p.SetState(133)
			p.Match(LumenParserT__27)
		}

	case LumenParserT__34:
		localctx = NewVertexDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(LumenParserT__34)
		}
		{
			p.SetState(136)

			var _m = p.Match(LumenParserID)

			localctx.(*VertexDeclContext).name = _m
		}
		{
			p.SetState(137)
			p.Match(LumenParserT__26)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LumenParserID {
			{
				p.SetState(138)

				var _x = p.VertexComponent()

				localctx.(*VertexDeclContext)._vertexComponent = _x
			}
			localctx.(*VertexDeclContext).components = append(localctx.(*VertexDeclContext).components, localctx.(*VertexDeclContext)._vertexComponent)

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(144)
			p.Match(LumenParserT__27)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_topLevelDecl returns the _topLevelDecl rule contexts.
	Get_topLevelDecl() ITopLevelDeclContext

	// Set_topLevelDecl sets the _topLevelDecl rule contexts.
	Set_topLevelDecl(ITopLevelDeclContext)

	// GetDecls returns the decls rule context list.
	GetDecls() []ITopLevelDeclContext

	// SetDecls sets the decls rule context list.
	SetDecls([]ITopLevelDeclContext)

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	_topLevelDecl ITopLevelDeclContext
	decls         []ITopLevelDeclContext
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

func (s *FileContext) Get_topLevelDecl() ITopLevelDeclContext { return s._topLevelDecl }

func (s *FileContext) Set_topLevelDecl(v ITopLevelDeclContext) { s._topLevelDecl = v }

func (s *FileContext) GetDecls() []ITopLevelDeclContext { return s.decls }

func (s *FileContext) SetDecls(v []ITopLevelDeclContext) { s.decls = v }

func (s *FileContext) AllTopLevelDecl() []ITopLevelDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem())
	var tst = make([]ITopLevelDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelDeclContext)
		}
	}

	return tst
}

func (s *FileContext) TopLevelDecl(i int) ITopLevelDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelDeclContext)
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
	p.EnterRule(localctx, 12, LumenParserRULE_file)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LumenParserT__28 || _la == LumenParserT__34 {
		{
			p.SetState(147)

			var _x = p.TopLevelDecl()

			localctx.(*FileContext)._topLevelDecl = _x
		}
		localctx.(*FileContext).decls = append(localctx.(*FileContext).decls, localctx.(*FileContext)._topLevelDecl)

		p.SetState(152)
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
