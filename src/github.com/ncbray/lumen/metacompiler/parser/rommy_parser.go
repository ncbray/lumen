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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 14, 73, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 5, 2, 16, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7,
	4, 26, 10, 4, 12, 4, 14, 4, 29, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 7, 5, 37, 10, 5, 12, 5, 14, 5, 40, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 7, 5, 47, 10, 5, 12, 5, 14, 5, 50, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 7, 5, 57, 10, 5, 12, 5, 14, 5, 60, 11, 5, 3, 5, 5, 5, 63, 10, 5, 3,
	6, 7, 6, 66, 10, 6, 12, 6, 14, 6, 69, 11, 6, 3, 6, 3, 6, 3, 6, 2, 2, 7,
	2, 4, 6, 8, 10, 2, 2, 75, 2, 15, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 22,
	3, 2, 2, 2, 8, 62, 3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12, 16, 7, 11, 2, 2,
	13, 14, 7, 3, 2, 2, 14, 16, 5, 2, 2, 2, 15, 12, 3, 2, 2, 2, 15, 13, 3,
	2, 2, 2, 16, 3, 3, 2, 2, 2, 17, 18, 7, 4, 2, 2, 18, 19, 7, 11, 2, 2, 19,
	20, 5, 2, 2, 2, 20, 21, 7, 5, 2, 2, 21, 5, 3, 2, 2, 2, 22, 23, 7, 11, 2,
	2, 23, 27, 7, 6, 2, 2, 24, 26, 5, 4, 3, 2, 25, 24, 3, 2, 2, 2, 26, 29,
	3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2,
	29, 27, 3, 2, 2, 2, 30, 31, 7, 7, 2, 2, 31, 7, 3, 2, 2, 2, 32, 33, 7, 8,
	2, 2, 33, 34, 7, 11, 2, 2, 34, 38, 7, 6, 2, 2, 35, 37, 5, 6, 4, 2, 36,
	35, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2,
	2, 39, 41, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 63, 7, 7, 2, 2, 42, 43,
	7, 9, 2, 2, 43, 44, 7, 11, 2, 2, 44, 48, 7, 6, 2, 2, 45, 47, 5, 4, 3, 2,
	46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3,
	2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 63, 7, 7, 2, 2, 52,
	53, 7, 10, 2, 2, 53, 54, 7, 11, 2, 2, 54, 58, 7, 6, 2, 2, 55, 57, 5, 8,
	5, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59,
	3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 63, 7, 7, 2, 2,
	62, 32, 3, 2, 2, 2, 62, 42, 3, 2, 2, 2, 62, 52, 3, 2, 2, 2, 63, 9, 3, 2,
	2, 2, 64, 66, 5, 8, 5, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2,
	70, 71, 7, 2, 2, 3, 71, 11, 3, 2, 2, 2, 9, 15, 27, 38, 48, 58, 62, 67,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'[]'", "'var'", "';'", "'{'", "'}'", "'enum'", "'struct'", "'region'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "ID", "NUM", "SINGLE_LINE_COMMENT",
	"WS",
}

var ruleNames = []string{
	"typeRef", "memberDecl", "variantDecl", "declaration", "file",
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
	RommyParserID                  = 9
	RommyParserNUM                 = 10
	RommyParserSINGLE_LINE_COMMENT = 11
	RommyParserWS                  = 12
)

// RommyParser rules.
const (
	RommyParserRULE_typeRef     = 0
	RommyParserRULE_memberDecl  = 1
	RommyParserRULE_variantDecl = 2
	RommyParserRULE_declaration = 3
	RommyParserRULE_file        = 4
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

	p.SetState(13)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RommyParserID:
		localctx = NewTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)

			var _m = p.Match(RommyParserID)

			localctx.(*TypeNameContext).name = _m
		}

	case RommyParserT__0:
		localctx = NewListRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Match(RommyParserT__0)
		}
		{
			p.SetState(12)

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
		p.SetState(15)
		p.Match(RommyParserT__1)
	}
	{
		p.SetState(16)

		var _m = p.Match(RommyParserID)

		localctx.(*FieldDeclContext).name = _m
	}
	{
		p.SetState(17)

		var _x = p.TypeRef()

		localctx.(*FieldDeclContext).t = _x
	}
	{
		p.SetState(18)
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
		p.SetState(20)

		var _m = p.Match(RommyParserID)

		localctx.(*VariantDeclContext).name = _m
	}
	{
		p.SetState(21)
		p.Match(RommyParserT__3)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RommyParserT__1 {
		{
			p.SetState(22)

			var _x = p.MemberDecl()

			localctx.(*VariantDeclContext)._memberDecl = _x
		}
		localctx.(*VariantDeclContext).members = append(localctx.(*VariantDeclContext).members, localctx.(*VariantDeclContext)._memberDecl)

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
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
	p.EnterRule(localctx, 6, RommyParserRULE_declaration)
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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RommyParserT__5:
		localctx = NewEnumDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Match(RommyParserT__5)
		}
		{
			p.SetState(31)

			var _m = p.Match(RommyParserID)

			localctx.(*EnumDeclContext).name = _m
		}
		{
			p.SetState(32)
			p.Match(RommyParserT__3)
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RommyParserID {
			{
				p.SetState(33)

				var _x = p.VariantDecl()

				localctx.(*EnumDeclContext)._variantDecl = _x
			}
			localctx.(*EnumDeclContext).variants = append(localctx.(*EnumDeclContext).variants, localctx.(*EnumDeclContext)._variantDecl)

			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(39)
			p.Match(RommyParserT__4)
		}

	case RommyParserT__6:
		localctx = NewStructDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(RommyParserT__6)
		}
		{
			p.SetState(41)

			var _m = p.Match(RommyParserID)

			localctx.(*StructDeclContext).name = _m
		}
		{
			p.SetState(42)
			p.Match(RommyParserT__3)
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RommyParserT__1 {
			{
				p.SetState(43)

				var _x = p.MemberDecl()

				localctx.(*StructDeclContext)._memberDecl = _x
			}
			localctx.(*StructDeclContext).members = append(localctx.(*StructDeclContext).members, localctx.(*StructDeclContext)._memberDecl)

			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(49)
			p.Match(RommyParserT__4)
		}

	case RommyParserT__7:
		localctx = NewRegionDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(RommyParserT__7)
		}
		{
			p.SetState(51)

			var _m = p.Match(RommyParserID)

			localctx.(*RegionDeclContext).name = _m
		}
		{
			p.SetState(52)
			p.Match(RommyParserT__3)
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RommyParserT__5)|(1<<RommyParserT__6)|(1<<RommyParserT__7))) != 0 {
			{
				p.SetState(53)

				var _x = p.Declaration()

				localctx.(*RegionDeclContext)._declaration = _x
			}
			localctx.(*RegionDeclContext).decls = append(localctx.(*RegionDeclContext).decls, localctx.(*RegionDeclContext)._declaration)

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(59)
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
	p.EnterRule(localctx, 8, RommyParserRULE_file)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RommyParserT__5)|(1<<RommyParserT__6)|(1<<RommyParserT__7))) != 0 {
		{
			p.SetState(62)

			var _x = p.Declaration()

			localctx.(*FileContext)._declaration = _x
		}
		localctx.(*FileContext).decls = append(localctx.(*FileContext).decls, localctx.(*FileContext)._declaration)

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(RommyParserEOF)
	}

	return localctx
}
