// Generated from Rommy.g by ANTLR 4.6.

package parser // Rommy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRommyListener is a complete listener for a parse tree produced by RommyParser.
type BaseRommyListener struct{}

var _ RommyListener = &BaseRommyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRommyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRommyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRommyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRommyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseRommyListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseRommyListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterListRef is called when production listRef is entered.
func (s *BaseRommyListener) EnterListRef(ctx *ListRefContext) {}

// ExitListRef is called when production listRef is exited.
func (s *BaseRommyListener) ExitListRef(ctx *ListRefContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseRommyListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseRommyListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterVariantDecl is called when production variantDecl is entered.
func (s *BaseRommyListener) EnterVariantDecl(ctx *VariantDeclContext) {}

// ExitVariantDecl is called when production variantDecl is exited.
func (s *BaseRommyListener) ExitVariantDecl(ctx *VariantDeclContext) {}

// EnterKeywordArg is called when production keywordArg is entered.
func (s *BaseRommyListener) EnterKeywordArg(ctx *KeywordArgContext) {}

// ExitKeywordArg is called when production keywordArg is exited.
func (s *BaseRommyListener) ExitKeywordArg(ctx *KeywordArgContext) {}

// EnterConstruct is called when production construct is entered.
func (s *BaseRommyListener) EnterConstruct(ctx *ConstructContext) {}

// ExitConstruct is called when production construct is exited.
func (s *BaseRommyListener) ExitConstruct(ctx *ConstructContext) {}

// EnterNameRef is called when production nameRef is entered.
func (s *BaseRommyListener) EnterNameRef(ctx *NameRefContext) {}

// ExitNameRef is called when production nameRef is exited.
func (s *BaseRommyListener) ExitNameRef(ctx *NameRefContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseRommyListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseRommyListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterParserBindingMapping is called when production parserBindingMapping is entered.
func (s *BaseRommyListener) EnterParserBindingMapping(ctx *ParserBindingMappingContext) {}

// ExitParserBindingMapping is called when production parserBindingMapping is exited.
func (s *BaseRommyListener) ExitParserBindingMapping(ctx *ParserBindingMappingContext) {}

// EnterParserBindingGroup is called when production parserBindingGroup is entered.
func (s *BaseRommyListener) EnterParserBindingGroup(ctx *ParserBindingGroupContext) {}

// ExitParserBindingGroup is called when production parserBindingGroup is exited.
func (s *BaseRommyListener) ExitParserBindingGroup(ctx *ParserBindingGroupContext) {}

// EnterEnumDecl is called when production enumDecl is entered.
func (s *BaseRommyListener) EnterEnumDecl(ctx *EnumDeclContext) {}

// ExitEnumDecl is called when production enumDecl is exited.
func (s *BaseRommyListener) ExitEnumDecl(ctx *EnumDeclContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseRommyListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseRommyListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterHolderDecl is called when production holderDecl is entered.
func (s *BaseRommyListener) EnterHolderDecl(ctx *HolderDeclContext) {}

// ExitHolderDecl is called when production holderDecl is exited.
func (s *BaseRommyListener) ExitHolderDecl(ctx *HolderDeclContext) {}

// EnterRegionDecl is called when production regionDecl is entered.
func (s *BaseRommyListener) EnterRegionDecl(ctx *RegionDeclContext) {}

// ExitRegionDecl is called when production regionDecl is exited.
func (s *BaseRommyListener) ExitRegionDecl(ctx *RegionDeclContext) {}

// EnterParserBindingDecl is called when production parserBindingDecl is entered.
func (s *BaseRommyListener) EnterParserBindingDecl(ctx *ParserBindingDeclContext) {}

// ExitParserBindingDecl is called when production parserBindingDecl is exited.
func (s *BaseRommyListener) ExitParserBindingDecl(ctx *ParserBindingDeclContext) {}

// EnterFile is called when production file is entered.
func (s *BaseRommyListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseRommyListener) ExitFile(ctx *FileContext) {}
