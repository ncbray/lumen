// Generated from Rommy.g by ANTLR 4.6.

package parser // Rommy

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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 21, 155, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 5, 2,
	26, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 36, 10,
	4, 12, 4, 14, 4, 39, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 7, 6, 52, 10, 6, 12, 6, 14, 6, 55, 11, 6, 5, 6, 57,
	10, 6, 3, 6, 3, 6, 5, 6, 61, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 7, 8, 71, 10, 8, 12, 8, 14, 8, 74, 11, 8, 5, 8, 76, 10, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 87, 10, 9,
	12, 9, 14, 9, 90, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10,
	98, 10, 10, 12, 10, 14, 10, 101, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 7, 10, 108, 10, 10, 12, 10, 14, 10, 111, 11, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 120, 10, 10, 12, 10, 14, 10, 123, 11,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 130, 10, 10, 12, 10, 14,
	10, 133, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 139, 10, 10, 12, 10,
	14, 10, 142, 11, 10, 3, 10, 5, 10, 145, 10, 10, 3, 11, 7, 11, 148, 10,
	11, 12, 11, 14, 11, 151, 11, 11, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 2, 2, 162, 2, 25, 3, 2, 2, 2, 4, 27, 3, 2, 2,
	2, 6, 32, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 60, 3, 2, 2, 2, 12, 62, 3,
	2, 2, 2, 14, 65, 3, 2, 2, 2, 16, 81, 3, 2, 2, 2, 18, 144, 3, 2, 2, 2, 20,
	149, 3, 2, 2, 2, 22, 26, 7, 18, 2, 2, 23, 24, 7, 3, 2, 2, 24, 26, 5, 2,
	2, 2, 25, 22, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26, 3, 3, 2, 2, 2, 27, 28,
	7, 4, 2, 2, 28, 29, 7, 18, 2, 2, 29, 30, 5, 2, 2, 2, 30, 31, 7, 5, 2, 2,
	31, 5, 3, 2, 2, 2, 32, 33, 7, 18, 2, 2, 33, 37, 7, 6, 2, 2, 34, 36, 5,
	4, 3, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37,
	38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 41, 7, 7, 2,
	2, 41, 7, 3, 2, 2, 2, 42, 43, 7, 18, 2, 2, 43, 44, 7, 8, 2, 2, 44, 45,
	5, 10, 6, 2, 45, 9, 3, 2, 2, 2, 46, 47, 7, 18, 2, 2, 47, 56, 7, 6, 2, 2,
	48, 53, 5, 8, 5, 2, 49, 50, 7, 9, 2, 2, 50, 52, 5, 8, 5, 2, 51, 49, 3,
	2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54,
	57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 48, 3, 2, 2, 2, 56, 57, 3, 2, 2,
	2, 57, 58, 3, 2, 2, 2, 58, 61, 7, 7, 2, 2, 59, 61, 7, 18, 2, 2, 60, 46,
	3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 11, 3, 2, 2, 2, 62, 63, 7, 18, 2, 2,
	63, 64, 5, 2, 2, 2, 64, 13, 3, 2, 2, 2, 65, 66, 7, 18, 2, 2, 66, 75, 7,
	10, 2, 2, 67, 72, 5, 12, 7, 2, 68, 69, 7, 9, 2, 2, 69, 71, 5, 12, 7, 2,
	70, 68, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3,
	2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 67, 3, 2, 2, 2, 75,
	76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 7, 11, 2, 2, 78, 79, 7, 12,
	2, 2, 79, 80, 5, 10, 6, 2, 80, 15, 3, 2, 2, 2, 81, 82, 7, 18, 2, 2, 82,
	83, 7, 8, 2, 2, 83, 84, 5, 2, 2, 2, 84, 88, 7, 6, 2, 2, 85, 87, 5, 14,
	8, 2, 86, 85, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89,
	3, 2, 2, 2, 89, 91, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 92, 7, 7, 2, 2,
	92, 17, 3, 2, 2, 2, 93, 94, 7, 13, 2, 2, 94, 95, 7, 18, 2, 2, 95, 99, 7,
	6, 2, 2, 96, 98, 5, 6, 4, 2, 97, 96, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99,
	97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2, 2, 101, 99, 3,
	2, 2, 2, 102, 145, 7, 7, 2, 2, 103, 104, 7, 14, 2, 2, 104, 105, 7, 18,
	2, 2, 105, 109, 7, 6, 2, 2, 106, 108, 5, 4, 3, 2, 107, 106, 3, 2, 2, 2,
	108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110,
	112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 145, 7, 7, 2, 2, 113, 114,
	7, 15, 2, 2, 114, 115, 7, 18, 2, 2, 115, 121, 7, 6, 2, 2, 116, 117, 5,
	2, 2, 2, 117, 118, 7, 5, 2, 2, 118, 120, 3, 2, 2, 2, 119, 116, 3, 2, 2,
	2, 120, 123, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122,
	124, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 124, 145, 7, 7, 2, 2, 125, 126,
	7, 16, 2, 2, 126, 127, 7, 18, 2, 2, 127, 131, 7, 6, 2, 2, 128, 130, 5,
	18, 10, 2, 129, 128, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2,
	2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2,
	134, 145, 7, 7, 2, 2, 135, 136, 7, 17, 2, 2, 136, 140, 7, 6, 2, 2, 137,
	139, 5, 16, 9, 2, 138, 137, 3, 2, 2, 2, 139, 142, 3, 2, 2, 2, 140, 138,
	3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2, 142, 140, 3, 2,
	2, 2, 143, 145, 7, 7, 2, 2, 144, 93, 3, 2, 2, 2, 144, 103, 3, 2, 2, 2,
	144, 113, 3, 2, 2, 2, 144, 125, 3, 2, 2, 2, 144, 135, 3, 2, 2, 2, 145,
	19, 3, 2, 2, 2, 146, 148, 5, 18, 10, 2, 147, 146, 3, 2, 2, 2, 148, 151,
	3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2,
	2, 2, 151, 149, 3, 2, 2, 2, 152, 153, 7, 2, 2, 3, 153, 21, 3, 2, 2, 2,
	17, 25, 37, 53, 56, 60, 72, 75, 88, 99, 109, 121, 131, 140, 144, 149,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'[]'", "'var'", "';'", "'{'", "'}'", "':'", "','", "'('", "')'", "'=>'",
	"'enum'", "'struct'", "'holder'", "'region'", "'parser'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "NUM",
	"SINGLE_LINE_COMMENT", "WS",
}

var ruleNames = []string{
	"typeRef", "memberDecl", "variantDecl", "keywordArg", "parserBindingExpr",
	"paramDecl", "parserBindingMapping", "parserBindingGroup", "declaration",
	"file",
}

type RommyParser struct {
	*antlr.BaseParser
}

func NewRommyParser(input antlr.TokenStream) *RommyParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(RommyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rommy.g"

	return this
}

// RommyParser tokens.
const (
	RommyParserEOF                 = antlr.TokenEOF
	RommyParserT__0                = 1
	RommyParserT__1                = 2
	RommyParserT__2                = 3
	RommyParserT__3                = 4
	RommyParserT__4                = 5
	RommyParserT__5                = 6
	RommyParserT__6                = 7
	RommyParserT__7                = 8
	RommyParserT__8                = 9
	RommyParserT__9                = 10
	RommyParserT__10               = 11
	RommyParserT__11               = 12
	RommyParserT__12               = 13
	RommyParserT__13               = 14
	RommyParserT__14               = 15
	RommyParserID                  = 16
	RommyParserNUM                 = 17
	RommyParserSINGLE_LINE_COMMENT = 18
	RommyParserWS                  = 19
)

// RommyParser rules.
const (
	RommyParserRULE_typeRef              = 0
	RommyParserRULE_memberDecl           = 1
	RommyParserRULE_variantDecl          = 2
	RommyParserRULE_keywordArg           = 3
	RommyParserRULE_parserBindingExpr    = 4
	RommyParserRULE_paramDecl            = 5
	RommyParserRULE_parserBindingMapping = 6
	RommyParserRULE_parserBindingGroup   = 7
	RommyParserRULE_declaration          = 8
	RommyParserRULE_file                 = 9
)

// ITypeRefContext is an interface to support dynamic dispatch.
type ITypeRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRefContext differentiates from other interfaces.
	IsTypeRefContext()
}

type TypeRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRefContext() *TypeRefContext {
	var p = new(TypeRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_typeRef
	return p
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_typeRef

	return p
}

func (s *TypeRefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRefContext) CopyFrom(ctx *TypeRefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeNameContext struct {
	*TypeRefContext
	name antlr.Token
}

func NewTypeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeNameContext {
	var p = new(TypeNameContext)

	p.TypeRefContext = NewEmptyTypeRefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeRefContext))

	return p
}

func (s *TypeNameContext) GetName() antlr.Token { return s.name }

func (s *TypeNameContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitTypeName(s)
	}
}

type ListRefContext struct {
	*TypeRefContext
	element ITypeRefContext
}

func NewListRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListRefContext {
	var p = new(ListRefContext)

	p.TypeRefContext = NewEmptyTypeRefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeRefContext))

	return p
}

func (s *ListRefContext) GetElement() ITypeRefContext { return s.element }

func (s *ListRefContext) SetElement(v ITypeRefContext) { s.element = v }

func (s *ListRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListRefContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *ListRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterListRef(s)
	}
}

func (s *ListRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitListRef(s)
	}
}

func (p *RommyParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RommyParserRULE_typeRef)

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

	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RommyParserID:
		localctx = NewTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)

			var _m = p.Match(RommyParserID)

			localctx.(*TypeNameContext).name = _m
		}

	case RommyParserT__0:
		localctx = NewListRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Match(RommyParserT__0)
		}
		{
			p.SetState(22)

			var _x = p.TypeRef()

			localctx.(*ListRefContext).element = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberDeclContext is an interface to support dynamic dispatch.
type IMemberDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberDeclContext differentiates from other interfaces.
	IsMemberDeclContext()
}

type MemberDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberDeclContext() *MemberDeclContext {
	var p = new(MemberDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_memberDecl
	return p
}

func (*MemberDeclContext) IsMemberDeclContext() {}

func NewMemberDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberDeclContext {
	var p = new(MemberDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_memberDecl

	return p
}

func (s *MemberDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberDeclContext) CopyFrom(ctx *MemberDeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MemberDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldDeclContext struct {
	*MemberDeclContext
	name antlr.Token
	t    ITypeRefContext
}

func NewFieldDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldDeclContext {
	var p = new(FieldDeclContext)

	p.MemberDeclContext = NewEmptyMemberDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MemberDeclContext))

	return p
}

func (s *FieldDeclContext) GetName() antlr.Token { return s.name }

func (s *FieldDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *FieldDeclContext) GetT() ITypeRefContext { return s.t }

func (s *FieldDeclContext) SetT(v ITypeRefContext) { s.t = v }

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *FieldDeclContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *FieldDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterFieldDecl(s)
	}
}

func (s *FieldDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitFieldDecl(s)
	}
}

func (p *RommyParser) MemberDecl() (localctx IMemberDeclContext) {
	localctx = NewMemberDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RommyParserRULE_memberDecl)

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

	localctx = NewFieldDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(RommyParserT__1)
	}
	{
		p.SetState(26)

		var _m = p.Match(RommyParserID)

		localctx.(*FieldDeclContext).name = _m
	}
	{
		p.SetState(27)

		var _x = p.TypeRef()

		localctx.(*FieldDeclContext).t = _x
	}
	{
		p.SetState(28)
		p.Match(RommyParserT__2)
	}

	return localctx
}

// IVariantDeclContext is an interface to support dynamic dispatch.
type IVariantDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_memberDecl returns the _memberDecl rule contexts.
	Get_memberDecl() IMemberDeclContext

	// Set_memberDecl sets the _memberDecl rule contexts.
	Set_memberDecl(IMemberDeclContext)

	// GetMembers returns the members rule context list.
	GetMembers() []IMemberDeclContext

	// SetMembers sets the members rule context list.
	SetMembers([]IMemberDeclContext)

	// IsVariantDeclContext differentiates from other interfaces.
	IsVariantDeclContext()
}

type VariantDeclContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	name        antlr.Token
	_memberDecl IMemberDeclContext
	members     []IMemberDeclContext
}

func NewEmptyVariantDeclContext() *VariantDeclContext {
	var p = new(VariantDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_variantDecl
	return p
}

func (*VariantDeclContext) IsVariantDeclContext() {}

func NewVariantDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariantDeclContext {
	var p = new(VariantDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_variantDecl

	return p
}

func (s *VariantDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariantDeclContext) GetName() antlr.Token { return s.name }

func (s *VariantDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *VariantDeclContext) Get_memberDecl() IMemberDeclContext { return s._memberDecl }

func (s *VariantDeclContext) Set_memberDecl(v IMemberDeclContext) { s._memberDecl = v }

func (s *VariantDeclContext) GetMembers() []IMemberDeclContext { return s.members }

func (s *VariantDeclContext) SetMembers(v []IMemberDeclContext) { s.members = v }

func (s *VariantDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *VariantDeclContext) AllMemberDecl() []IMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberDeclContext)(nil)).Elem())
	var tst = make([]IMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberDeclContext)
		}
	}

	return tst
}

func (s *VariantDeclContext) MemberDecl(i int) IMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberDeclContext)
}

func (s *VariantDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariantDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariantDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterVariantDecl(s)
	}
}

func (s *VariantDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitVariantDecl(s)
	}
}

func (p *RommyParser) VariantDecl() (localctx IVariantDeclContext) {
	localctx = NewVariantDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RommyParserRULE_variantDecl)
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
		p.SetState(30)

		var _m = p.Match(RommyParserID)

		localctx.(*VariantDeclContext).name = _m
	}
	{
		p.SetState(31)
		p.Match(RommyParserT__3)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RommyParserT__1 {
		{
			p.SetState(32)

			var _x = p.MemberDecl()

			localctx.(*VariantDeclContext)._memberDecl = _x
		}
		localctx.(*VariantDeclContext).members = append(localctx.(*VariantDeclContext).members, localctx.(*VariantDeclContext)._memberDecl)

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(RommyParserT__4)
	}

	return localctx
}

// IKeywordArgContext is an interface to support dynamic dispatch.
type IKeywordArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IParserBindingExprContext

	// SetValue sets the value rule contexts.
	SetValue(IParserBindingExprContext)

	// IsKeywordArgContext differentiates from other interfaces.
	IsKeywordArgContext()
}

type KeywordArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	value  IParserBindingExprContext
}

func NewEmptyKeywordArgContext() *KeywordArgContext {
	var p = new(KeywordArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_keywordArg
	return p
}

func (*KeywordArgContext) IsKeywordArgContext() {}

func NewKeywordArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordArgContext {
	var p = new(KeywordArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_keywordArg

	return p
}

func (s *KeywordArgContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordArgContext) GetName() antlr.Token { return s.name }

func (s *KeywordArgContext) SetName(v antlr.Token) { s.name = v }

func (s *KeywordArgContext) GetValue() IParserBindingExprContext { return s.value }

func (s *KeywordArgContext) SetValue(v IParserBindingExprContext) { s.value = v }

func (s *KeywordArgContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *KeywordArgContext) ParserBindingExpr() IParserBindingExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParserBindingExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParserBindingExprContext)
}

func (s *KeywordArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterKeywordArg(s)
	}
}

func (s *KeywordArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitKeywordArg(s)
	}
}

func (p *RommyParser) KeywordArg() (localctx IKeywordArgContext) {
	localctx = NewKeywordArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RommyParserRULE_keywordArg)

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
		p.SetState(40)

		var _m = p.Match(RommyParserID)

		localctx.(*KeywordArgContext).name = _m
	}
	{
		p.SetState(41)
		p.Match(RommyParserT__5)
	}
	{
		p.SetState(42)

		var _x = p.ParserBindingExpr()

		localctx.(*KeywordArgContext).value = _x
	}

	return localctx
}

// IParserBindingExprContext is an interface to support dynamic dispatch.
type IParserBindingExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParserBindingExprContext differentiates from other interfaces.
	IsParserBindingExprContext()
}

type ParserBindingExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParserBindingExprContext() *ParserBindingExprContext {
	var p = new(ParserBindingExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_parserBindingExpr
	return p
}

func (*ParserBindingExprContext) IsParserBindingExprContext() {}

func NewParserBindingExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParserBindingExprContext {
	var p = new(ParserBindingExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_parserBindingExpr

	return p
}

func (s *ParserBindingExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ParserBindingExprContext) CopyFrom(ctx *ParserBindingExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ParserBindingExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParserBindingExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NameRefContext struct {
	*ParserBindingExprContext
	name antlr.Token
}

func NewNameRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameRefContext {
	var p = new(NameRefContext)

	p.ParserBindingExprContext = NewEmptyParserBindingExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParserBindingExprContext))

	return p
}

func (s *NameRefContext) GetName() antlr.Token { return s.name }

func (s *NameRefContext) SetName(v antlr.Token) { s.name = v }

func (s *NameRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameRefContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *NameRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterNameRef(s)
	}
}

func (s *NameRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitNameRef(s)
	}
}

type ConstructContext struct {
	*ParserBindingExprContext
	name        antlr.Token
	_keywordArg IKeywordArgContext
	args        []IKeywordArgContext
}

func NewConstructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstructContext {
	var p = new(ConstructContext)

	p.ParserBindingExprContext = NewEmptyParserBindingExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParserBindingExprContext))

	return p
}

func (s *ConstructContext) GetName() antlr.Token { return s.name }

func (s *ConstructContext) SetName(v antlr.Token) { s.name = v }

func (s *ConstructContext) Get_keywordArg() IKeywordArgContext { return s._keywordArg }

func (s *ConstructContext) Set_keywordArg(v IKeywordArgContext) { s._keywordArg = v }

func (s *ConstructContext) GetArgs() []IKeywordArgContext { return s.args }

func (s *ConstructContext) SetArgs(v []IKeywordArgContext) { s.args = v }

func (s *ConstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *ConstructContext) AllKeywordArg() []IKeywordArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeywordArgContext)(nil)).Elem())
	var tst = make([]IKeywordArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeywordArgContext)
		}
	}

	return tst
}

func (s *ConstructContext) KeywordArg(i int) IKeywordArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeywordArgContext)
}

func (s *ConstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterConstruct(s)
	}
}

func (s *ConstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitConstruct(s)
	}
}

func (p *RommyParser) ParserBindingExpr() (localctx IParserBindingExprContext) {
	localctx = NewParserBindingExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RommyParserRULE_parserBindingExpr)
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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)

			var _m = p.Match(RommyParserID)

			localctx.(*ConstructContext).name = _m
		}
		{
			p.SetState(45)
			p.Match(RommyParserT__3)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RommyParserID {
			{
				p.SetState(46)

				var _x = p.KeywordArg()

				localctx.(*ConstructContext)._keywordArg = _x
			}
			localctx.(*ConstructContext).args = append(localctx.(*ConstructContext).args, localctx.(*ConstructContext)._keywordArg)
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == RommyParserT__6 {
				{
					p.SetState(47)
					p.Match(RommyParserT__6)
				}
				{
					p.SetState(48)

					var _x = p.KeywordArg()

					localctx.(*ConstructContext)._keywordArg = _x
				}
				localctx.(*ConstructContext).args = append(localctx.(*ConstructContext).args, localctx.(*ConstructContext)._keywordArg)

				p.SetState(53)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(56)
			p.Match(RommyParserT__4)
		}

	case 2:
		localctx = NewNameRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)

			var _m = p.Match(RommyParserID)

			localctx.(*NameRefContext).name = _m
		}

	}

	return localctx
}

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITypeRefContext

	// SetT sets the t rule contexts.
	SetT(ITypeRefContext)

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	t      ITypeRefContext
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_paramDecl
	return p
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) GetName() antlr.Token { return s.name }

func (s *ParamDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *ParamDeclContext) GetT() ITypeRefContext { return s.t }

func (s *ParamDeclContext) SetT(v ITypeRefContext) { s.t = v }

func (s *ParamDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *ParamDeclContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterParamDecl(s)
	}
}

func (s *ParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitParamDecl(s)
	}
}

func (p *RommyParser) ParamDecl() (localctx IParamDeclContext) {
	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RommyParserRULE_paramDecl)

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
		p.SetState(60)

		var _m = p.Match(RommyParserID)

		localctx.(*ParamDeclContext).name = _m
	}
	{
		p.SetState(61)

		var _x = p.TypeRef()

		localctx.(*ParamDeclContext).t = _x
	}

	return localctx
}

// IParserBindingMappingContext is an interface to support dynamic dispatch.
type IParserBindingMappingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_paramDecl returns the _paramDecl rule contexts.
	Get_paramDecl() IParamDeclContext

	// GetBody returns the body rule contexts.
	GetBody() IParserBindingExprContext

	// Set_paramDecl sets the _paramDecl rule contexts.
	Set_paramDecl(IParamDeclContext)

	// SetBody sets the body rule contexts.
	SetBody(IParserBindingExprContext)

	// GetParams returns the params rule context list.
	GetParams() []IParamDeclContext

	// SetParams sets the params rule context list.
	SetParams([]IParamDeclContext)

	// IsParserBindingMappingContext differentiates from other interfaces.
	IsParserBindingMappingContext()
}

type ParserBindingMappingContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	name       antlr.Token
	_paramDecl IParamDeclContext
	params     []IParamDeclContext
	body       IParserBindingExprContext
}

func NewEmptyParserBindingMappingContext() *ParserBindingMappingContext {
	var p = new(ParserBindingMappingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_parserBindingMapping
	return p
}

func (*ParserBindingMappingContext) IsParserBindingMappingContext() {}

func NewParserBindingMappingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParserBindingMappingContext {
	var p = new(ParserBindingMappingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_parserBindingMapping

	return p
}

func (s *ParserBindingMappingContext) GetParser() antlr.Parser { return s.parser }

func (s *ParserBindingMappingContext) GetName() antlr.Token { return s.name }

func (s *ParserBindingMappingContext) SetName(v antlr.Token) { s.name = v }

func (s *ParserBindingMappingContext) Get_paramDecl() IParamDeclContext { return s._paramDecl }

func (s *ParserBindingMappingContext) GetBody() IParserBindingExprContext { return s.body }

func (s *ParserBindingMappingContext) Set_paramDecl(v IParamDeclContext) { s._paramDecl = v }

func (s *ParserBindingMappingContext) SetBody(v IParserBindingExprContext) { s.body = v }

func (s *ParserBindingMappingContext) GetParams() []IParamDeclContext { return s.params }

func (s *ParserBindingMappingContext) SetParams(v []IParamDeclContext) { s.params = v }

func (s *ParserBindingMappingContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *ParserBindingMappingContext) ParserBindingExpr() IParserBindingExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParserBindingExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParserBindingExprContext)
}

func (s *ParserBindingMappingContext) AllParamDecl() []IParamDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamDeclContext)(nil)).Elem())
	var tst = make([]IParamDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamDeclContext)
		}
	}

	return tst
}

func (s *ParserBindingMappingContext) ParamDecl(i int) IParamDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *ParserBindingMappingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParserBindingMappingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParserBindingMappingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterParserBindingMapping(s)
	}
}

func (s *ParserBindingMappingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitParserBindingMapping(s)
	}
}

func (p *RommyParser) ParserBindingMapping() (localctx IParserBindingMappingContext) {
	localctx = NewParserBindingMappingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RommyParserRULE_parserBindingMapping)
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
		p.SetState(63)

		var _m = p.Match(RommyParserID)

		localctx.(*ParserBindingMappingContext).name = _m
	}
	{
		p.SetState(64)
		p.Match(RommyParserT__7)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RommyParserID {
		{
			p.SetState(65)

			var _x = p.ParamDecl()

			localctx.(*ParserBindingMappingContext)._paramDecl = _x
		}
		localctx.(*ParserBindingMappingContext).params = append(localctx.(*ParserBindingMappingContext).params, localctx.(*ParserBindingMappingContext)._paramDecl)
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RommyParserT__6 {
			{
				p.SetState(66)
				p.Match(RommyParserT__6)
			}
			{
				p.SetState(67)

				var _x = p.ParamDecl()

				localctx.(*ParserBindingMappingContext)._paramDecl = _x
			}
			localctx.(*ParserBindingMappingContext).params = append(localctx.(*ParserBindingMappingContext).params, localctx.(*ParserBindingMappingContext)._paramDecl)

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(75)
		p.Match(RommyParserT__8)
	}
	{
		p.SetState(76)
		p.Match(RommyParserT__9)
	}
	{
		p.SetState(77)

		var _x = p.ParserBindingExpr()

		localctx.(*ParserBindingMappingContext).body = _x
	}

	return localctx
}

// IParserBindingGroupContext is an interface to support dynamic dispatch.
type IParserBindingGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITypeRefContext

	// Get_parserBindingMapping returns the _parserBindingMapping rule contexts.
	Get_parserBindingMapping() IParserBindingMappingContext

	// SetT sets the t rule contexts.
	SetT(ITypeRefContext)

	// Set_parserBindingMapping sets the _parserBindingMapping rule contexts.
	Set_parserBindingMapping(IParserBindingMappingContext)

	// GetMappings returns the mappings rule context list.
	GetMappings() []IParserBindingMappingContext

	// SetMappings sets the mappings rule context list.
	SetMappings([]IParserBindingMappingContext)

	// IsParserBindingGroupContext differentiates from other interfaces.
	IsParserBindingGroupContext()
}

type ParserBindingGroupContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	name                  antlr.Token
	t                     ITypeRefContext
	_parserBindingMapping IParserBindingMappingContext
	mappings              []IParserBindingMappingContext
}

func NewEmptyParserBindingGroupContext() *ParserBindingGroupContext {
	var p = new(ParserBindingGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_parserBindingGroup
	return p
}

func (*ParserBindingGroupContext) IsParserBindingGroupContext() {}

func NewParserBindingGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParserBindingGroupContext {
	var p = new(ParserBindingGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_parserBindingGroup

	return p
}

func (s *ParserBindingGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *ParserBindingGroupContext) GetName() antlr.Token { return s.name }

func (s *ParserBindingGroupContext) SetName(v antlr.Token) { s.name = v }

func (s *ParserBindingGroupContext) GetT() ITypeRefContext { return s.t }

func (s *ParserBindingGroupContext) Get_parserBindingMapping() IParserBindingMappingContext {
	return s._parserBindingMapping
}

func (s *ParserBindingGroupContext) SetT(v ITypeRefContext) { s.t = v }

func (s *ParserBindingGroupContext) Set_parserBindingMapping(v IParserBindingMappingContext) {
	s._parserBindingMapping = v
}

func (s *ParserBindingGroupContext) GetMappings() []IParserBindingMappingContext { return s.mappings }

func (s *ParserBindingGroupContext) SetMappings(v []IParserBindingMappingContext) { s.mappings = v }

func (s *ParserBindingGroupContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *ParserBindingGroupContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *ParserBindingGroupContext) AllParserBindingMapping() []IParserBindingMappingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParserBindingMappingContext)(nil)).Elem())
	var tst = make([]IParserBindingMappingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParserBindingMappingContext)
		}
	}

	return tst
}

func (s *ParserBindingGroupContext) ParserBindingMapping(i int) IParserBindingMappingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParserBindingMappingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParserBindingMappingContext)
}

func (s *ParserBindingGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParserBindingGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParserBindingGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterParserBindingGroup(s)
	}
}

func (s *ParserBindingGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitParserBindingGroup(s)
	}
}

func (p *RommyParser) ParserBindingGroup() (localctx IParserBindingGroupContext) {
	localctx = NewParserBindingGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RommyParserRULE_parserBindingGroup)
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

		var _m = p.Match(RommyParserID)

		localctx.(*ParserBindingGroupContext).name = _m
	}
	{
		p.SetState(80)
		p.Match(RommyParserT__5)
	}
	{
		p.SetState(81)

		var _x = p.TypeRef()

		localctx.(*ParserBindingGroupContext).t = _x
	}
	{
		p.SetState(82)
		p.Match(RommyParserT__3)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RommyParserID {
		{
			p.SetState(83)

			var _x = p.ParserBindingMapping()

			localctx.(*ParserBindingGroupContext)._parserBindingMapping = _x
		}
		localctx.(*ParserBindingGroupContext).mappings = append(localctx.(*ParserBindingGroupContext).mappings, localctx.(*ParserBindingGroupContext)._parserBindingMapping)

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(RommyParserT__4)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CopyFrom(ctx *DeclarationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EnumDeclContext struct {
	*DeclarationContext
	name         antlr.Token
	_variantDecl IVariantDeclContext
	variants     []IVariantDeclContext
}

func NewEnumDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumDeclContext {
	var p = new(EnumDeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *EnumDeclContext) GetName() antlr.Token { return s.name }

func (s *EnumDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *EnumDeclContext) Get_variantDecl() IVariantDeclContext { return s._variantDecl }

func (s *EnumDeclContext) Set_variantDecl(v IVariantDeclContext) { s._variantDecl = v }

func (s *EnumDeclContext) GetVariants() []IVariantDeclContext { return s.variants }

func (s *EnumDeclContext) SetVariants(v []IVariantDeclContext) { s.variants = v }

func (s *EnumDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *EnumDeclContext) AllVariantDecl() []IVariantDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariantDeclContext)(nil)).Elem())
	var tst = make([]IVariantDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariantDeclContext)
		}
	}

	return tst
}

func (s *EnumDeclContext) VariantDecl(i int) IVariantDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariantDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariantDeclContext)
}

func (s *EnumDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterEnumDecl(s)
	}
}

func (s *EnumDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitEnumDecl(s)
	}
}

type HolderDeclContext struct {
	*DeclarationContext
	name     antlr.Token
	_typeRef ITypeRefContext
	types    []ITypeRefContext
}

func NewHolderDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HolderDeclContext {
	var p = new(HolderDeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *HolderDeclContext) GetName() antlr.Token { return s.name }

func (s *HolderDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *HolderDeclContext) Get_typeRef() ITypeRefContext { return s._typeRef }

func (s *HolderDeclContext) Set_typeRef(v ITypeRefContext) { s._typeRef = v }

func (s *HolderDeclContext) GetTypes() []ITypeRefContext { return s.types }

func (s *HolderDeclContext) SetTypes(v []ITypeRefContext) { s.types = v }

func (s *HolderDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HolderDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *HolderDeclContext) AllTypeRef() []ITypeRefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeRefContext)(nil)).Elem())
	var tst = make([]ITypeRefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeRefContext)
		}
	}

	return tst
}

func (s *HolderDeclContext) TypeRef(i int) ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *HolderDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterHolderDecl(s)
	}
}

func (s *HolderDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitHolderDecl(s)
	}
}

type ParserBindingDeclContext struct {
	*DeclarationContext
	_parserBindingGroup IParserBindingGroupContext
	groups              []IParserBindingGroupContext
}

func NewParserBindingDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParserBindingDeclContext {
	var p = new(ParserBindingDeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *ParserBindingDeclContext) Get_parserBindingGroup() IParserBindingGroupContext {
	return s._parserBindingGroup
}

func (s *ParserBindingDeclContext) Set_parserBindingGroup(v IParserBindingGroupContext) {
	s._parserBindingGroup = v
}

func (s *ParserBindingDeclContext) GetGroups() []IParserBindingGroupContext { return s.groups }

func (s *ParserBindingDeclContext) SetGroups(v []IParserBindingGroupContext) { s.groups = v }

func (s *ParserBindingDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParserBindingDeclContext) AllParserBindingGroup() []IParserBindingGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParserBindingGroupContext)(nil)).Elem())
	var tst = make([]IParserBindingGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParserBindingGroupContext)
		}
	}

	return tst
}

func (s *ParserBindingDeclContext) ParserBindingGroup(i int) IParserBindingGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParserBindingGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParserBindingGroupContext)
}

func (s *ParserBindingDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterParserBindingDecl(s)
	}
}

func (s *ParserBindingDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitParserBindingDecl(s)
	}
}

type StructDeclContext struct {
	*DeclarationContext
	name        antlr.Token
	_memberDecl IMemberDeclContext
	members     []IMemberDeclContext
}

func NewStructDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclContext {
	var p = new(StructDeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *StructDeclContext) GetName() antlr.Token { return s.name }

func (s *StructDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *StructDeclContext) Get_memberDecl() IMemberDeclContext { return s._memberDecl }

func (s *StructDeclContext) Set_memberDecl(v IMemberDeclContext) { s._memberDecl = v }

func (s *StructDeclContext) GetMembers() []IMemberDeclContext { return s.members }

func (s *StructDeclContext) SetMembers(v []IMemberDeclContext) { s.members = v }

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *StructDeclContext) AllMemberDecl() []IMemberDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberDeclContext)(nil)).Elem())
	var tst = make([]IMemberDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberDeclContext)
		}
	}

	return tst
}

func (s *StructDeclContext) MemberDecl(i int) IMemberDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberDeclContext)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

type RegionDeclContext struct {
	*DeclarationContext
	name         antlr.Token
	_declaration IDeclarationContext
	decls        []IDeclarationContext
}

func NewRegionDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegionDeclContext {
	var p = new(RegionDeclContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *RegionDeclContext) GetName() antlr.Token { return s.name }

func (s *RegionDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *RegionDeclContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *RegionDeclContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *RegionDeclContext) GetDecls() []IDeclarationContext { return s.decls }

func (s *RegionDeclContext) SetDecls(v []IDeclarationContext) { s.decls = v }

func (s *RegionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegionDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(RommyParserID, 0)
}

func (s *RegionDeclContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *RegionDeclContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *RegionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterRegionDecl(s)
	}
}

func (s *RegionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitRegionDecl(s)
	}
}

func (p *RommyParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RommyParserRULE_declaration)
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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RommyParserT__10:
		localctx = NewEnumDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Match(RommyParserT__10)
		}
		{
			p.SetState(92)

			var _m = p.Match(RommyParserID)

			localctx.(*EnumDeclContext).name = _m
		}
		{
			p.SetState(93)
			p.Match(RommyParserT__3)
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RommyParserID {
			{
				p.SetState(94)

				var _x = p.VariantDecl()

				localctx.(*EnumDeclContext)._variantDecl = _x
			}
			localctx.(*EnumDeclContext).variants = append(localctx.(*EnumDeclContext).variants, localctx.(*EnumDeclContext)._variantDecl)

			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(100)
			p.Match(RommyParserT__4)
		}

	case RommyParserT__11:
		localctx = NewStructDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(RommyParserT__11)
		}
		{
			p.SetState(102)

			var _m = p.Match(RommyParserID)

			localctx.(*StructDeclContext).name = _m
		}
		{
			p.SetState(103)
			p.Match(RommyParserT__3)
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RommyParserT__1 {
			{
				p.SetState(104)

				var _x = p.MemberDecl()

				localctx.(*StructDeclContext)._memberDecl = _x
			}
			localctx.(*StructDeclContext).members = append(localctx.(*StructDeclContext).members, localctx.(*StructDeclContext)._memberDecl)

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(110)
			p.Match(RommyParserT__4)
		}

	case RommyParserT__12:
		localctx = NewHolderDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.Match(RommyParserT__12)
		}
		{
			p.SetState(112)

			var _m = p.Match(RommyParserID)

			localctx.(*HolderDeclContext).name = _m
		}
		{
			p.SetState(113)
			p.Match(RommyParserT__3)
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RommyParserT__0 || _la == RommyParserID {
			{
				p.SetState(114)

				var _x = p.TypeRef()

				localctx.(*HolderDeclContext)._typeRef = _x
			}
			localctx.(*HolderDeclContext).types = append(localctx.(*HolderDeclContext).types, localctx.(*HolderDeclContext)._typeRef)
			{
				p.SetState(115)
				p.Match(RommyParserT__2)
			}

			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(122)
			p.Match(RommyParserT__4)
		}

	case RommyParserT__13:
		localctx = NewRegionDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.Match(RommyParserT__13)
		}
		{
			p.SetState(124)

			var _m = p.Match(RommyParserID)

			localctx.(*RegionDeclContext).name = _m
		}
		{
			p.SetState(125)
			p.Match(RommyParserT__3)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RommyParserT__10)|(1<<RommyParserT__11)|(1<<RommyParserT__12)|(1<<RommyParserT__13)|(1<<RommyParserT__14))) != 0 {
			{
				p.SetState(126)

				var _x = p.Declaration()

				localctx.(*RegionDeclContext)._declaration = _x
			}
			localctx.(*RegionDeclContext).decls = append(localctx.(*RegionDeclContext).decls, localctx.(*RegionDeclContext)._declaration)

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(132)
			p.Match(RommyParserT__4)
		}

	case RommyParserT__14:
		localctx = NewParserBindingDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(133)
			p.Match(RommyParserT__14)
		}
		{
			p.SetState(134)
			p.Match(RommyParserT__3)
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RommyParserID {
			{
				p.SetState(135)

				var _x = p.ParserBindingGroup()

				localctx.(*ParserBindingDeclContext)._parserBindingGroup = _x
			}
			localctx.(*ParserBindingDeclContext).groups = append(localctx.(*ParserBindingDeclContext).groups, localctx.(*ParserBindingDeclContext)._parserBindingGroup)

			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(141)
			p.Match(RommyParserT__4)
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

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// GetDecls returns the decls rule context list.
	GetDecls() []IDeclarationContext

	// SetDecls sets the decls rule context list.
	SetDecls([]IDeclarationContext)

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_declaration IDeclarationContext
	decls        []IDeclarationContext
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RommyParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RommyParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *FileContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *FileContext) GetDecls() []IDeclarationContext { return s.decls }

func (s *FileContext) SetDecls(v []IDeclarationContext) { s.decls = v }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(RommyParserEOF, 0)
}

func (s *FileContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *FileContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RommyListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *RommyParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RommyParserRULE_file)
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
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RommyParserT__10)|(1<<RommyParserT__11)|(1<<RommyParserT__12)|(1<<RommyParserT__13)|(1<<RommyParserT__14))) != 0 {
		{
			p.SetState(144)

			var _x = p.Declaration()

			localctx.(*FileContext)._declaration = _x
		}
		localctx.(*FileContext).decls = append(localctx.(*FileContext).decls, localctx.(*FileContext)._declaration)

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(RommyParserEOF)
	}

	return localctx
}
