grammar Query;

options { caseInsensitive = true; }

query
    : expr EOF
    ;

expr
    : predicate      # PredicateExpr
    | '(' expr ')'   # NestedExpr
    | NOT expr       # NotExpr
    | expr AND expr  # AndExpr
    | expr OR expr   # OrExpr
    ;

predicate
    : primary                                  # PrimaryPred
    | primary IS NOT? NULL                     # IsPred
    | primary NOT? op=(IN|LIKE|MATCH) primary  # OpPred
    | primary cmp=(LT|LE|GT|GE|EQ|NE) primary  # CmpPred
    ;

primary
    : field
    | apply
    | value
    | tuple
    ;

field
    : IDENT
    ;

apply
    : IDENT '(' ')'
    | IDENT '(' args+=primary (',' args+=primary)* ')'
    ;

value
    : STRING   # StringValue
    | NUMBER   # NumberValue
    | BOOLEAN  # BooleanValue
    | NULL     # NullValue
    ;

tuple
    : '(' items+=value (',' items+=value)* ')'  # ValueTuple
    | '(' items+=field (',' items+=field)* ')'  # FieldTuple
    ;

LT:    '<';
LE:    '<=';
GT:    '>';
GE:    '>=';
EQ:    '=';
NE:    '!=';
NOT:   'not';
AND:   'and';
OR:    'or';
IS:    'is';
IN:    'in';
LIKE:  'like';
MATCH: 'match';

STRING
    : DQUOTA_STRING
    | SQUOTA_STRING
    ;

NUMBER
    : '-'? INT ('.' [0-9] +)? EXP?
    ;

BOOLEAN
    : 'true'
    | 'false'
    ;

NULL
    : 'null'
    ;


IDENT
    : [a-z_][a-z_0-9]*
    ;

WS
    : [ \r\n\t]+ -> skip
    ;

fragment DQUOTA_STRING
    : '"' ( '\\'. | ~('"'| '\\') )* '"'
    ;
fragment SQUOTA_STRING
    : '\'' ('\\'. | ~('\'' | '\\'))* '\''
    ;
fragment INT
    : '0' | [1-9] [0-9]*
    ;
fragment EXP
    : 'e' [+\-]? INT
    ;
