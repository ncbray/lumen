grammar Rommy;

import Unicode;

typeRef
  : name=ID # typeName
  | '[]' element=typeRef # listRef
  ;

memberDecl
  : 'var' name=ID t=typeRef ';' # fieldDecl
  ;

variantDecl
  : name=ID '{' (members+=memberDecl)* '}'
  ;

declaration
  : 'enum' name=ID '{' (variants+=variantDecl)* '}' # enumDecl
  | 'struct' name=ID '{' (members+=memberDecl)* '}' # structDecl
  | 'region' name=ID '{' (decls+=declaration)* '}' # regionDecl
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
