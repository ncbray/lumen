// Generated from Rommy.g by ANTLR 4.6.

package parser // Rommy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RommyListener is a complete listener for a parse tree produced by RommyParser.
type RommyListener interface {
	antlr.ParseTreeListener

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterListRef is called when entering the listRef production.
	EnterListRef(c *ListRefContext)

	// EnterFieldDecl is called when entering the fieldDecl production.
	EnterFieldDecl(c *FieldDeclContext)

	// EnterKeywordArg is called when entering the keywordArg production.
	EnterKeywordArg(c *KeywordArgContext)

	// EnterConstruct is called when entering the construct production.
	EnterConstruct(c *ConstructContext)

	// EnterNameRef is called when entering the nameRef production.
	EnterNameRef(c *NameRefContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterParserBindingMapping is called when entering the parserBindingMapping production.
	EnterParserBindingMapping(c *ParserBindingMappingContext)

	// EnterParserBindingGroup is called when entering the parserBindingGroup production.
	EnterParserBindingGroup(c *ParserBindingGroupContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterHolderDecl is called when entering the holderDecl production.
	EnterHolderDecl(c *HolderDeclContext)

	// EnterRegionDecl is called when entering the regionDecl production.
	EnterRegionDecl(c *RegionDeclContext)

	// EnterParserBindingDecl is called when entering the parserBindingDecl production.
	EnterParserBindingDecl(c *ParserBindingDeclContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitListRef is called when exiting the listRef production.
	ExitListRef(c *ListRefContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitKeywordArg is called when exiting the keywordArg production.
	ExitKeywordArg(c *KeywordArgContext)

	// ExitConstruct is called when exiting the construct production.
	ExitConstruct(c *ConstructContext)

	// ExitNameRef is called when exiting the nameRef production.
	ExitNameRef(c *NameRefContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitParserBindingMapping is called when exiting the parserBindingMapping production.
	ExitParserBindingMapping(c *ParserBindingMappingContext)

	// ExitParserBindingGroup is called when exiting the parserBindingGroup production.
	ExitParserBindingGroup(c *ParserBindingGroupContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitHolderDecl is called when exiting the holderDecl production.
	ExitHolderDecl(c *HolderDeclContext)

	// ExitRegionDecl is called when exiting the regionDecl production.
	ExitRegionDecl(c *RegionDeclContext)

	// ExitParserBindingDecl is called when exiting the parserBindingDecl production.
	ExitParserBindingDecl(c *ParserBindingDeclContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)
}
