grammar Lumen;

// TODO paren
expr
  : op=('+'|'-'|'~'|'!') value=expr # prefix
  | left=expr op=('*'|'/'|'%') right=expr # infix
  | left=expr op=('+'|'-') right=expr # infix
  | left=expr op=('<<'|'>>') right=expr # infix
  | left=expr op=('<'|'>'|'<='|'>=') right=expr # infix
  | left=expr op=('=='|'!=') right=expr # infix
  | left=expr op='&' right=expr # infix
  | left=expr op='^' right=expr # infix
  | left=expr op='|' right=expr # infix
  | left=expr op='&&' right=expr # infix
  | left=expr op='||' right=expr # infix
  | value=expr '(' (args+=expr (',' args+=expr)*)? ')' # call
  | raw=ID # getName
  | raw=NUM # number
  ;

statement
  : t=ID name=ID '=' value=expr ';' # varDecl
  | name=ID '=' value=expr ';' # assign
  | value=expr ';' # discard
  ;

shaderDecl
  : 'shader' name=ID '{'
      'vs' '{' (vs+=statement)* '}'
      'fs' '{' (fs+=statement)* '}'
    '}'
  ;

file
  : (shaders+=shaderDecl)*
  ;

NUM: [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)?;

ID: Letter (Letter | DecimalDigit) *;

fragment Letter: [a-zA-Z_];

fragment DecimalDigit: [0-9];

SINGLE_LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
WS: [ \t\r\n]+ -> channel(HIDDEN);
