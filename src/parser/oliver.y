%{
package parser
%}

%union{
str string
number int
functions []*Function
function *Function
stmt  interface{}
stmts []interface{}
assignrhs interface{}
assignlhs interface{}
expr  interface{}
exprs []interface{}
params  []Param
param Param
arglist []interface{}
bracketed []interface{}
pairliter interface{}
arrayliter ArrayLiter
pairelem  PairElem
arrayelem ArrayElem
typedefinition Type
pairelemtype Type
}

%start program

%token BEGIN END                                     // Program delimiters
%token IS
%token <number> SKIP
%token READ FREE RETURN EXIT PRINT PRINTLN
%token IF THEN ELSE FI                               // If statement
%token WHILE DO DONE                                 // While statement
%token NEWPAIR
%token CALL
%token FST SND
%token <number> INT BOOL CHAR STRING PAIR
%token <number> NOT NEG LEN ORD CHR                                              // Unary ops   DO WE NEED THESEEEE
%token <number> MUL DIV MOD PLUS SUB AND OR GT GTE LT LTE EQ NEQ                 // Binary ops
%token POSITIVE NEGATIVE
%token <number> TRUE FALSE                                    // Booleans
%token NULL
%token OPENSQUARE OPENROUND CLOSESQUARE CLOSEROUND
%token ASSIGNMENT
%token COMMA SEMICOLON
%token ERROR

%token <str> STRINGCONST
%token <str> IDENTIFIER
%token <number> INTEGER
%token <str> CHARACTER

%type <prog> program
%type <functions> functions
%type <function> function
%type <stmt> statement
%type <stmts> statements
%type <assignrhs> assignrhs
%type <assignlhs> assignlhs
%type	<expr>		expr
%type <exprs> exprlist
%type <params> paramlist
%type <param> param
%type <pairelem> pairelem
%type <arrayliter> arrayliter
%type <arrayelem>  arrayelem
%type <bracketed> bracketed
%type <pairliter> pairliter
%type <typedefinition>  basetype typeDef arraytype pairtype
%type <pairelemtype> pairelemtype

%left OR AND EQ NEQ LT GT LTE GTE PLUS SUB MUL DIV MOD
%right NOT NEG LEN ORD CHR

%%

program : BEGIN functions statements END {
                                          parserlex.(*Lexer).prog = &Program{functionlist : $2 , statList : $3 , symbolTable : &SymbolTable{Table: make(map[string]Type)}}
                                         }

functions : functions function  { $$ = append($1, $2)}
          |                     { $$ = []*Function{} }

function : typeDef IDENTIFIER OPENROUND CLOSEROUND IS statements END
           { if !checkStats($6) {
          	parserlex.Error("Missing return statement")
           }
             $$ = &Function{ident : $2, returnType : $1, statlist : $6}
           }
         | typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND IS statements END
           { if !checkStats($7) {
            	parserlex.Error("Missing return statement")
            }
             $$ = &Function{ident : $2, returnType : $1, statlist : $7, parameterTypes : $4}
           }

paramlist : paramlist COMMA param { $$ = append($1, $3)}
          | param                 { $$ = []Param{ $1 } }

param : typeDef IDENTIFIER { $$ = Param{paramType : $1, ident : $2} }

assignlhs : IDENTIFIER    {$$ = $1}
          | arrayelem     {$$ = $1}
          | pairelem      {$$ = $1}

assignrhs : expr                                           {$$ = $1}
          | arrayliter                                     {$$ = $1}
          | pairelem                                       {$$ = $1}
          | NEWPAIR OPENROUND expr COMMA expr CLOSEROUND   { $$ = NewPair{fstExpr : $3, sndExpr : $5} }
          | CALL IDENTIFIER OPENROUND exprlist CLOSEROUND  { $$ = Call{ident : $2, exprlist : $4} }

statements : statements SEMICOLON statement           { $$ = append($1,$3) }
           | statement                                { $$ = []interface{}{$1} }


statement : SKIP                                      { $$ = $1 }
          | typeDef IDENTIFIER ASSIGNMENT assignrhs   { $$ = Declare{Type : $1, lhs : $2, rhs : $4} }
          | assignlhs ASSIGNMENT assignrhs            { $$ = Assignment{lhs : $1, rhs : $3} }
          | READ assignlhs                            { $$ = Read{$2} }
          | FREE expr                                 { $$ = Free{$2} }
          | RETURN expr                               { $$ = Return{$2} }
          | EXIT expr                                 { $$ = Exit{$2} }
          | PRINT expr                                { $$ = Print{$2} }
          | PRINTLN expr                              { $$ = Println{$2} }
          | IF expr THEN statements ELSE statements FI  { $$ = If{conditional : $2, thenStat : $4, elseStat : $6} }
          | WHILE expr DO statements DONE              { $$ = While{conditional : $2, doStat : $4} }
          | BEGIN statements END                       { $$ = Scope{statlist : $2, symbolTable : &SymbolTable{Table: make(map[string]Type)} } }

expr : INTEGER       { $$ =  $1 }
     | TRUE          { $$ =  $1 }
     | FALSE         { $$ =  $1 }
     | CHARACTER     { $$ =  $1 }
     | STRINGCONST   { $$ =  $1 }
     | pairliter     { $$ =  $1 }
     | IDENTIFIER    { $$ =  $1 }
     | arrayelem     { $$ =  $1 }
     | NOT expr     { $$ = Unop{unary : $1, expr : $2} }
     | LEN expr     { $$ = Unop{unary : $1, expr : $2} }
     | ORD expr     { $$ = Unop{unary : $1, expr : $2} }
     | CHR expr     { $$ = Unop{unary : $1, expr : $2} }
     | SUB expr     { $$ = Unop{unary : $1, expr : $2} }
     | PLUS expr    { $$ = Unop{unary : $1, expr : $2} }

     | expr PLUS expr { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr SUB expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr MUL expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr MOD expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr DIV expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr LT expr   { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr GT expr   { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr LTE expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr GTE expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr EQ expr   { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr NEQ expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr AND expr  { $$ = Binop{left : $1, binary : $2, right : $3} }
     | expr OR expr   { $$ = Binop{left : $1, binary : $2, right : $3} }
     | OPENROUND expr CLOSEROUND  { $$ = $2 }

arrayliter : OPENSQUARE exprlist CLOSESQUARE { $$ = ArrayLiter{$2}}

exprlist : exprlist COMMA expr {$$ = append($1, $3)}
         | expr                {$$ = []interface{}{$1}}
         |                     {$$ = []interface{}{}}

arrayelem : IDENTIFIER bracketed {$$ = ArrayElem{ident: $1, exprs : $2}}

bracketed : bracketed OPENSQUARE expr CLOSESQUARE {$$ = append($1, $3)}
          | OPENSQUARE expr CLOSESQUARE {$$ = []interface{}{$2}}

pairliter : NULL    { $$ =  NULL }

pairelem : FST expr { $$ = PairElem{fsnd: FST, expr : $2} }
         | SND expr { $$ = PairElem{fsnd: SND, expr : $2} }

pairtype : PAIR OPENROUND pairelemtype COMMA pairelemtype CLOSEROUND  { $$ = PairType{fstType : $3, sndType : $5} }

pairelemtype : basetype  { $$ = $1 }
             | arraytype { $$ = $1 }
             | PAIR      { $$ = Pair }

typeDef : basetype  { $$ =  $1 }
        | arraytype { $$ =  $1 }
        | pairtype  { $$ =  $1 }

basetype : INT      { $$ =  Int }
         | BOOL     { $$ =  Bool }
         | CHAR     { $$ =  Char }
         | STRING   { $$ =  String }

arraytype : typeDef OPENSQUARE CLOSESQUARE { $$ = ArrayType{Type : $1} }

%%
