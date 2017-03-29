grammar Rommy;

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

typeDecl
  : 'enum' name=ID '{' (variants+=variantDecl)* '}' # enumDecl
  | 'struct' name=ID '{' (members+=memberDecl)* '}' # structDecl
  ;

file
  : (decls+=typeDecl)* EOF
  ;

ID: [a-zA-Z][a-zA-Z0-9]*;
NUM: [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)?;
WS: [ \t\r\n] -> channel(HIDDEN);
