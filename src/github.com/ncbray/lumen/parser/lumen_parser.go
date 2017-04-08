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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 36, 111, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 5, 2, 16, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 7, 2, 53, 10, 2, 12, 2, 14, 2, 56, 11, 2, 5, 2, 58, 10,
	2, 3, 2, 7, 2, 61, 10, 2, 12, 2, 14, 2, 64, 11, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 80,
	10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 88, 10, 4, 12, 4, 14,
	4, 91, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 97, 10, 4, 12, 4, 14, 4, 100,
	11, 4, 3, 4, 3, 4, 3, 4, 3, 5, 7, 5, 106, 10, 5, 12, 5, 14, 5, 109, 11,
	5, 3, 5, 2, 3, 2, 6, 2, 4, 6, 8, 2, 8, 3, 2, 3, 6, 3, 2, 7, 9, 3, 2, 3,
	4, 3, 2, 10, 11, 3, 2, 12, 15, 3, 2, 16, 17, 126, 2, 15, 3, 2, 2, 2, 4,
	79, 3, 2, 2, 2, 6, 81, 3, 2, 2, 2, 8, 107, 3, 2, 2, 2, 10, 11, 8, 2, 1,
	2, 11, 12, 9, 2, 2, 2, 12, 16, 5, 2, 2, 16, 13, 16, 7, 34, 2, 2, 14, 16,
	7, 33, 2, 2, 15, 10, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 14, 3, 2, 2, 2,
	16, 62, 3, 2, 2, 2, 17, 18, 12, 15, 2, 2, 18, 19, 9, 3, 2, 2, 19, 61, 5,
	2, 2, 16, 20, 21, 12, 14, 2, 2, 21, 22, 9, 4, 2, 2, 22, 61, 5, 2, 2, 15,
	23, 24, 12, 13, 2, 2, 24, 25, 9, 5, 2, 2, 25, 61, 5, 2, 2, 14, 26, 27,
	12, 12, 2, 2, 27, 28, 9, 6, 2, 2, 28, 61, 5, 2, 2, 13, 29, 30, 12, 11,
	2, 2, 30, 31, 9, 7, 2, 2, 31, 61, 5, 2, 2, 12, 32, 33, 12, 10, 2, 2, 33,
	34, 7, 18, 2, 2, 34, 61, 5, 2, 2, 11, 35, 36, 12, 9, 2, 2, 36, 37, 7, 19,
	2, 2, 37, 61, 5, 2, 2, 10, 38, 39, 12, 8, 2, 2, 39, 40, 7, 20, 2, 2, 40,
	61, 5, 2, 2, 9, 41, 42, 12, 7, 2, 2, 42, 43, 7, 21, 2, 2, 43, 61, 5, 2,
	2, 8, 44, 45, 12, 6, 2, 2, 45, 46, 7, 22, 2, 2, 46, 61, 5, 2, 2, 7, 47,
	48, 12, 5, 2, 2, 48, 57, 7, 23, 2, 2, 49, 54, 5, 2, 2, 2, 50, 51, 7, 24,
	2, 2, 51, 53, 5, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2,
	57, 49, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 7,
	25, 2, 2, 60, 17, 3, 2, 2, 2, 60, 20, 3, 2, 2, 2, 60, 23, 3, 2, 2, 2, 60,
	26, 3, 2, 2, 2, 60, 29, 3, 2, 2, 2, 60, 32, 3, 2, 2, 2, 60, 35, 3, 2, 2,
	2, 60, 38, 3, 2, 2, 2, 60, 41, 3, 2, 2, 2, 60, 44, 3, 2, 2, 2, 60, 47,
	3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2,
	63, 3, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 7, 34, 2, 2, 66, 67, 7,
	34, 2, 2, 67, 68, 7, 26, 2, 2, 68, 69, 5, 2, 2, 2, 69, 70, 7, 27, 2, 2,
	70, 80, 3, 2, 2, 2, 71, 72, 7, 34, 2, 2, 72, 73, 7, 26, 2, 2, 73, 74, 5,
	2, 2, 2, 74, 75, 7, 27, 2, 2, 75, 80, 3, 2, 2, 2, 76, 77, 5, 2, 2, 2, 77,
	78, 7, 27, 2, 2, 78, 80, 3, 2, 2, 2, 79, 65, 3, 2, 2, 2, 79, 71, 3, 2,
	2, 2, 79, 76, 3, 2, 2, 2, 80, 5, 3, 2, 2, 2, 81, 82, 7, 28, 2, 2, 82, 83,
	7, 34, 2, 2, 83, 84, 7, 29, 2, 2, 84, 85, 7, 30, 2, 2, 85, 89, 7, 29, 2,
	2, 86, 88, 5, 4, 3, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87,
	3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2,
	92, 93, 7, 31, 2, 2, 93, 94, 7, 32, 2, 2, 94, 98, 7, 29, 2, 2, 95, 97,
	5, 4, 3, 2, 96, 95, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2,
	98, 99, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 102,
	7, 31, 2, 2, 102, 103, 7, 31, 2, 2, 103, 7, 3, 2, 2, 2, 104, 106, 5, 6,
	4, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2,
	107, 108, 3, 2, 2, 2, 108, 9, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 11, 15,
	54, 57, 60, 62, 79, 89, 98, 107,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'~'", "'!'", "'*'", "'/'", "'%'", "'<<'", "'>>'", "'<'",
	"'>'", "'<='", "'>='", "'=='", "'!='", "'&'", "'^'", "'|'", "'&&'", "'||'",
	"'('", "','", "')'", "'='", "';'", "'shader'", "'{'", "'vs'", "'}'", "'fs'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "NUM", "ID", "SINGLE_LINE_COMMENT",
	"WS",
}

var ruleNames = []string{
	"expr", "statement", "shaderDecl", "file",
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
	LumenParserNUM                 = 31
	LumenParserID                  = 32
	LumenParserSINGLE_LINE_COMMENT = 33
	LumenParserWS                  = 34
)

// LumenParser rules.
const (
	LumenParserRULE_expr       = 0
	LumenParserRULE_statement  = 1
	LumenParserRULE_shaderDecl = 2
	LumenParserRULE_file       = 3
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LumenParserT__0, LumenParserT__1, LumenParserT__2, LumenParserT__3:
		localctx = NewPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(9)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PrefixContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__0)|(1<<LumenParserT__1)|(1<<LumenParserT__2)|(1<<LumenParserT__3))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PrefixContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(10)

			var _x = p.expr(14)

			localctx.(*PrefixContext).value = _x
		}

	case LumenParserID:
		localctx = NewGetNameContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)

			var _m = p.Match(LumenParserID)

			localctx.(*GetNameContext).name = _m
		}

	case LumenParserNUM:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)

			var _m = p.Match(LumenParserNUM)

			localctx.(*NumberContext).raw = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(15)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				p.SetState(16)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__4)|(1<<LumenParserT__5)|(1<<LumenParserT__6))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(17)

					var _x = p.expr(14)

					localctx.(*InfixContext).right = _x
				}

			case 2:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(18)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(19)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == LumenParserT__0 || _la == LumenParserT__1) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(20)

					var _x = p.expr(13)

					localctx.(*InfixContext).right = _x
				}

			case 3:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(22)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == LumenParserT__7 || _la == LumenParserT__8) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(23)

					var _x = p.expr(12)

					localctx.(*InfixContext).right = _x
				}

			case 4:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(25)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LumenParserT__9)|(1<<LumenParserT__10)|(1<<LumenParserT__11)|(1<<LumenParserT__12))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(26)

					var _x = p.expr(11)

					localctx.(*InfixContext).right = _x
				}

			case 5:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(28)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*InfixContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == LumenParserT__13 || _la == LumenParserT__14) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*InfixContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(29)

					var _x = p.expr(10)

					localctx.(*InfixContext).right = _x
				}

			case 6:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(31)

					var _m = p.Match(LumenParserT__15)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(32)

					var _x = p.expr(9)

					localctx.(*InfixContext).right = _x
				}

			case 7:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(34)

					var _m = p.Match(LumenParserT__16)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(35)

					var _x = p.expr(8)

					localctx.(*InfixContext).right = _x
				}

			case 8:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(37)

					var _m = p.Match(LumenParserT__17)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(38)

					var _x = p.expr(7)

					localctx.(*InfixContext).right = _x
				}

			case 9:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(40)

					var _m = p.Match(LumenParserT__18)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(41)

					var _x = p.expr(6)

					localctx.(*InfixContext).right = _x
				}

			case 10:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*InfixContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(43)

					var _m = p.Match(LumenParserT__19)

					localctx.(*InfixContext).op = _m
				}
				{
					p.SetState(44)

					var _x = p.expr(5)

					localctx.(*InfixContext).right = _x
				}

			case 11:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallContext).value = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LumenParserRULE_expr)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(46)
					p.Match(LumenParserT__20)
				}
				p.SetState(55)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(LumenParserT__0-1))|(1<<(LumenParserT__1-1))|(1<<(LumenParserT__2-1))|(1<<(LumenParserT__3-1))|(1<<(LumenParserNUM-1))|(1<<(LumenParserID-1)))) != 0 {
					{
						p.SetState(47)

						var _x = p.expr(0)

						localctx.(*CallContext)._expr = _x
					}
					localctx.(*CallContext).args = append(localctx.(*CallContext).args, localctx.(*CallContext)._expr)
					p.SetState(52)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LumenParserT__21 {
						{
							p.SetState(48)
							p.Match(LumenParserT__21)
						}
						{
							p.SetState(49)

							var _x = p.expr(0)

							localctx.(*CallContext)._expr = _x
						}
						localctx.(*CallContext).args = append(localctx.(*CallContext).args, localctx.(*CallContext)._expr)

						p.SetState(54)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(57)
					p.Match(LumenParserT__22)
				}

			}

		}
		p.SetState(62)
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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)

			var _m = p.Match(LumenParserID)

			localctx.(*VarDeclContext).t = _m
		}
		{
			p.SetState(64)

			var _m = p.Match(LumenParserID)

			localctx.(*VarDeclContext).name = _m
		}
		{
			p.SetState(65)
			p.Match(LumenParserT__23)
		}
		{
			p.SetState(66)

			var _x = p.expr(0)

			localctx.(*VarDeclContext).value = _x
		}
		{
			p.SetState(67)
			p.Match(LumenParserT__24)
		}

	case 2:
		localctx = NewAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)

			var _m = p.Match(LumenParserID)

			localctx.(*AssignContext).name = _m
		}
		{
			p.SetState(70)
			p.Match(LumenParserT__23)
		}
		{
			p.SetState(71)

			var _x = p.expr(0)

			localctx.(*AssignContext).value = _x
		}
		{
			p.SetState(72)
			p.Match(LumenParserT__24)
		}

	case 3:
		localctx = NewDiscardContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)

			var _x = p.expr(0)

			localctx.(*DiscardContext).value = _x
		}
		{
			p.SetState(75)
			p.Match(LumenParserT__24)
		}

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

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

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

func (s *ShaderDeclContext) Get_statement() IStatementContext { return s._statement }

func (s *ShaderDeclContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *ShaderDeclContext) GetVs() []IStatementContext { return s.vs }

func (s *ShaderDeclContext) GetFs() []IStatementContext { return s.fs }

func (s *ShaderDeclContext) SetVs(v []IStatementContext) { s.vs = v }

func (s *ShaderDeclContext) SetFs(v []IStatementContext) { s.fs = v }

func (s *ShaderDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(LumenParserID, 0)
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
	p.EnterRule(localctx, 4, LumenParserRULE_shaderDecl)
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
		p.SetState(79)
		p.Match(LumenParserT__25)
	}
	{
		p.SetState(80)

		var _m = p.Match(LumenParserID)

		localctx.(*ShaderDeclContext).name = _m
	}
	{
		p.SetState(81)
		p.Match(LumenParserT__26)
	}
	{
		p.SetState(82)
		p.Match(LumenParserT__27)
	}
	{
		p.SetState(83)
		p.Match(LumenParserT__26)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(LumenParserT__0-1))|(1<<(LumenParserT__1-1))|(1<<(LumenParserT__2-1))|(1<<(LumenParserT__3-1))|(1<<(LumenParserNUM-1))|(1<<(LumenParserID-1)))) != 0 {
		{
			p.SetState(84)

			var _x = p.Statement()

			localctx.(*ShaderDeclContext)._statement = _x
		}
		localctx.(*ShaderDeclContext).vs = append(localctx.(*ShaderDeclContext).vs, localctx.(*ShaderDeclContext)._statement)

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(LumenParserT__28)
	}
	{
		p.SetState(91)
		p.Match(LumenParserT__29)
	}
	{
		p.SetState(92)
		p.Match(LumenParserT__26)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(LumenParserT__0-1))|(1<<(LumenParserT__1-1))|(1<<(LumenParserT__2-1))|(1<<(LumenParserT__3-1))|(1<<(LumenParserNUM-1))|(1<<(LumenParserID-1)))) != 0 {
		{
			p.SetState(93)

			var _x = p.Statement()

			localctx.(*ShaderDeclContext)._statement = _x
		}
		localctx.(*ShaderDeclContext).fs = append(localctx.(*ShaderDeclContext).fs, localctx.(*ShaderDeclContext)._statement)

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(LumenParserT__28)
	}
	{
		p.SetState(100)
		p.Match(LumenParserT__28)
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
	p.EnterRule(localctx, 6, LumenParserRULE_file)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LumenParserT__25 {
		{
			p.SetState(102)

			var _x = p.ShaderDecl()

			localctx.(*FileContext)._shaderDecl = _x
		}
		localctx.(*FileContext).shaders = append(localctx.(*FileContext).shaders, localctx.(*FileContext)._shaderDecl)

		p.SetState(107)
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
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
