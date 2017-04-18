grammar Lumen;

// TODO paren
expr
  : value=expr '(' (args+=expr (',' args+=expr)*)? ')' # call
  | value=expr '.' name=ID # getAttr
  | op=('+'|'-'|'~'|'!') value=expr # prefix
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
  | name=ID # getName
  | raw=NUM # number
  ;

statement
  : t=ID name=ID '=' value=expr ';' # varDecl
  | name=ID '=' value=expr ';' # assign
  | value=expr ';' # discard
  ;

field
  : name=ID t=ID ';'
  ;

format
  : '{' (fields+=field)* '}'
  ;

vertexComponent
  : name=ID t=ID encoding=ID ';'
  ;

topLevelDecl
  : 'shader' name=ID '{'
      'uniform' uniform=format
      'attribute' attribute=format
      'varying' varying=format
      'vs' '{' (vs+=statement)* '}'
      'fs' '{' (fs+=statement)* '}'
    '}' # shaderDecl
  | 'vertex' name=ID '{' (components+=vertexComponent)* '}' # vertexDecl
  ;

file
  : (decls+=topLevelDecl)*
  ;

NUM: [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)?;

ID: Letter (Letter | DecimalDigit) *;

fragment Letter: [a-zA-Z_];

fragment DecimalDigit: [0-9];

SINGLE_LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
WS: [ \t\r\n]+ -> channel(HIDDEN);
