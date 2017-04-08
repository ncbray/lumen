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

	// EnterVariantDecl is called when entering the variantDecl production.
	EnterVariantDecl(c *VariantDeclContext)

	// EnterEnumDecl is called when entering the enumDecl production.
	EnterEnumDecl(c *EnumDeclContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterHolderDecl is called when entering the holderDecl production.
	EnterHolderDecl(c *HolderDeclContext)

	// EnterRegionDecl is called when entering the regionDecl production.
	EnterRegionDecl(c *RegionDeclContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitListRef is called when exiting the listRef production.
	ExitListRef(c *ListRefContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitVariantDecl is called when exiting the variantDecl production.
	ExitVariantDecl(c *VariantDeclContext)

	// ExitEnumDecl is called when exiting the enumDecl production.
	ExitEnumDecl(c *EnumDeclContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitHolderDecl is called when exiting the holderDecl production.
	ExitHolderDecl(c *HolderDeclContext)

	// ExitRegionDecl is called when exiting the regionDecl production.
	ExitRegionDecl(c *RegionDeclContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)
}
