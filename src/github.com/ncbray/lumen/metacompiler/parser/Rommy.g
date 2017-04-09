grammar Rommy;

import Unicode;

typeRef
  : name=ID # typeName
  | '[]' element=typeRef # listRef
  ;

memberDecl
  : 'var' name=ID t=typeRef ';' # fieldDecl
  ;

keywordArg
  : name=ID ':' value=parserBindingExpr
  ;

parserBindingExpr
  : name=ID '{' (args+=keywordArg (',' args+=keywordArg)*)? '}' # construct
  | name=ID # nameRef
  ;

paramDecl
  : name=ID t=typeRef
  ;

parserBindingMapping
  : name=ID '(' (params+=paramDecl (',' params+=paramDecl)*)? ')' '=>' body=parserBindingExpr
  ;

parserBindingGroup
  : name=ID ':' t=typeRef '{' (mappings+=parserBindingMapping)* '}'
  ;

declaration
  : 'struct' name=ID '{' (members+=memberDecl)* '}' # structDecl
  | 'holder' name=ID '=' (types+=typeRef ('|' types+=typeRef)*)? ';' # holderDecl
  | 'region' name=ID '{' (decls+=declaration)* '}' # regionDecl
  | 'parser' '{' (groups+=parserBindingGroup)* '}' # parserBindingDecl
  ;

file
  : (decls+=declaration)* EOF
  ;

ID: Letter (Letter | UnicodeDigit) *;
NUM: [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)?;

fragment Letter
  : UnicodeLetter
  | [_]
  ;

SINGLE_LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
WS: [ \t\r\n]+ -> channel(HIDDEN);
