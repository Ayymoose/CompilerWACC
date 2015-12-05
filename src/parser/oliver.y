%{
package parser

import (
. "ast"
)

%}

%union{
str string
stringconst Str
number int
integer Integer
ident Ident
character Character
boolean Boolean

functions []*Function
function *Function
stmt  Statement
stmts []Statement
assignrhs Evaluation
assignlhs Evaluation
expr  Evaluation
exprs []Evaluation
params  []Param
param Param
bracketed []Evaluation
pairliter Evaluation
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
%token  INT BOOL CHAR STRING PAIR
%token  NOT NEG LEN ORD CHR                                              // Unary ops   DO WE NEED THESEEEE
%token  MUL DIV MOD PLUS SUB AND OR GT GTE LT LTE EQ NEQ                 // Binary ops
%token POSITIVE NEGATIVE
%token <boolean> TRUE FALSE                                    // Booleans
%token NULL
%token OPENSQUARE OPENROUND CLOSESQUARE CLOSEROUND
%token ASSIGNMENT
%token COMMA SEMICOLON
%token ERROR


%token <stringconst> STRINGCONST
%token <ident> IDENTIFIER
%token <integer> INTEGER
%token <character> CHARACTER


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
%type <exprs> bracketed
%type <pairliter> pairliter
%type <typedefinition>  basetype typeDef arraytype pairtype
%type <pairelemtype> pairelemtype

%left OR AND EQ NEQ LT GT LTE GTE PLUS SUB MUL DIV MOD
%right NOT NEG LEN ORD CHR

%%

program : BEGIN functions statements END {
                                          parserlex.(*Lexer).prog = &Program{FunctionList : $2 , StatList : $3 , SymbolTable : &SymbolTable{Table: make(map[Ident]Type)}}
                                         }

functions : functions function  { $$ = append($1, $2)}
          |                     { $$ = []*Function{} }

function : typeDef IDENTIFIER OPENROUND CLOSEROUND IS statements END
           { if !checkStats($6) {
          	parserlex.Error("Missing return statement")
           }
             $$ = &Function{Ident : $2, ReturnType : $1, StatList : $6}
           }
         | typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND IS statements END
           { if !checkStats($7) {
            	parserlex.Error("Missing return statement")
            }
             $$ = &Function{Ident : $2, ReturnType : $1, StatList : $7, ParameterTypes : $4}
           }

paramlist : paramlist COMMA param { $$ = append($1, $3)}
          | param                 { $$ = []Param{ $1 } }

param : typeDef IDENTIFIER { $$ = Param{ParamType : $1, Ident : $2} }

assignlhs : IDENTIFIER    {$$ = $1}
          | arrayelem     {$$ = $1}
          | pairelem      {$$ = $1}

assignrhs : expr                                           {$$ = $1}
          | arrayliter                                     {$$ = $1}
          | pairelem                                       {$$ = $1}
          | NEWPAIR OPENROUND expr COMMA expr CLOSEROUND   { $$ = NewPair{FstExpr : $3, SndExpr : $5} }
          | CALL IDENTIFIER OPENROUND exprlist CLOSEROUND  { $$ = Call{Ident : $2, ParamList : $4} }

statements : statements SEMICOLON statement           { $$ = append($1,$3) }
           | statement                                { $$ = []Statement{$1} }


statement : SKIP                                        { $$ = Skip{} }
          | typeDef IDENTIFIER ASSIGNMENT assignrhs     { $$ = Declare{DecType : $1, Lhs : $2, Rhs : $4} }
          | assignlhs ASSIGNMENT assignrhs              { $$ = Assignment{Lhs : $1, Rhs : $3} }
          | READ assignlhs                              { $$ = Read{$2} }
          | FREE expr                                   { $$ = Free{$2} }
          | RETURN expr                                 { $$ = Return{$2} }
          | EXIT expr                                   { $$ = Exit{$2} }
          | PRINT expr                                  { $$ = Print{$2} }
          | PRINTLN expr                                { $$ = Println{$2} }
          | IF expr THEN statements ELSE statements FI  { $$ = If{Conditional : $2, ThenStat : $4, ElseStat : $6} }
          | WHILE expr DO statements DONE               { $$ = While{Conditional : $2, DoStat : $4} }
          | BEGIN statements END                        { $$ = Scope{StatList : $2, SymbolTable : &SymbolTable{Table: make(map[Ident]Type)} } }

expr : INTEGER       { $$ =  $1 }

     | TRUE          { $$ =  $1 }
     | FALSE         { $$ =  $1 }

     | CHARACTER     { $$ =  $1 }
     | STRINGCONST   { $$ =  $1 }
     | pairliter     { $$ =  PairLiter{} }
     | IDENTIFIER    { $$ =  $1 }
     | arrayelem     { $$ =  $1 }

     | NOT expr     { $$ = Unop{Unary : NOT, Expr : $2} }
     | LEN expr     { $$ = Unop{Unary : LEN, Expr : $2} }
     | ORD expr     { $$ = Unop{Unary : ORD, Expr : $2} }
     | CHR expr     { $$ = Unop{Unary : CHR, Expr : $2} }
     | SUB expr     { $$ = Unop{Unary : SUB, Expr : $2} }
     | PLUS expr    { $$ = $2 }


     | expr PLUS expr { $$ = Binop{Left : $1, Binary : PLUS, Right : $3} }
     | expr SUB expr  { $$ = Binop{Left : $1, Binary : SUB, Right : $3} }
     | expr MUL expr  { $$ = Binop{Left : $1, Binary : MUL, Right : $3} }
     | expr MOD expr  { $$ = Binop{Left : $1, Binary : MOD, Right : $3} }
     | expr DIV expr  { $$ = Binop{Left : $1, Binary : DIV, Right : $3} }
     | expr LT expr   { $$ = Binop{Left : $1, Binary : LT, Right : $3} }
     | expr GT expr   { $$ = Binop{Left : $1, Binary : GT, Right : $3} }
     | expr LTE expr  { $$ = Binop{Left : $1, Binary : LTE, Right : $3} }
     | expr GTE expr  { $$ = Binop{Left : $1, Binary : GTE, Right : $3} }
     | expr EQ expr   { $$ = Binop{Left : $1, Binary : EQ, Right : $3} }
     | expr NEQ expr  { $$ = Binop{Left : $1, Binary : NEQ, Right : $3} }
     | expr AND expr  { $$ = Binop{Left : $1, Binary : AND, Right : $3} }
     | expr OR expr   { $$ = Binop{Left : $1, Binary : OR, Right : $3} }
     | OPENROUND expr CLOSEROUND  { $$ = $2 }

arrayliter : OPENSQUARE exprlist CLOSESQUARE { $$ = ArrayLiter{$2}}

exprlist : exprlist COMMA expr {$$ = append($1, $3)}
         | expr                {$$ = []Evaluation{$1}}
         |                     {$$ = []Evaluation{}}

arrayelem : IDENTIFIER bracketed {$$ = ArrayElem{Ident: $1, Exprs : $2}}

bracketed : bracketed OPENSQUARE expr CLOSESQUARE {$$ = append($1, $3)}
          | OPENSQUARE expr CLOSESQUARE {$$ = []Evaluation{$2}}


pairliter : NULL    { $$ =  PairLiter{} }


pairelem : FST expr { $$ = PairElem{Fsnd: FST, Expr : $2} }
         | SND expr { $$ = PairElem{Fsnd: SND, Expr : $2} }

pairtype : PAIR OPENROUND pairelemtype COMMA pairelemtype CLOSEROUND  { $$ = PairType{FstType : $3, SndType : $5} }

pairelemtype : basetype  { $$ = $1 }
             | arraytype { $$ = $1 }
             | PAIR      { $$ = PairType{} }

typeDef : basetype  { $$ =  $1 }
        | arraytype { $$ =  $1 }
        | pairtype  { $$ =  $1 }

basetype : INT      { $$ =  Int }
         | BOOL     { $$ =  Bool }
         | CHAR     { $$ =  Char }
         | STRING   { $$ =  String }

arraytype : typeDef OPENSQUARE CLOSESQUARE { $$ = ArrayType{Type : $1} }

%%
