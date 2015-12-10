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
pos int
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

%left OR
%left AND
%left EQ NEQ
%left PLUS SUB
%left MUL DIV MOD
%left LT GT LTE GTE
%right NOT NEG LEN ORD CHR

%%

program : BEGIN functions statements END {
                                          parserlex.(*Lexer).prog = &Program{FunctionList : $2 , StatList : $3 , SymbolTable : NewInstance(), FileText :&parserlex.(*Lexer).input}
                                         }

functions : functions function  { $$ = append($1, $2)}
          |                     { $$ = []*Function{} }

function : typeDef IDENTIFIER OPENROUND CLOSEROUND IS statements END
           { if !checkStats($6) {
          	parserlex.Error("Missing return statement")
           }
             $$ = &Function{Ident : $2, ReturnType : $1, StatList : $6, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
           }
         | typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND IS statements END
           { if !checkStats($7) {
            	parserlex.Error("Missing return statement")
            }
             $$ = &Function{Ident : $2, ReturnType : $1, StatList : $7, ParameterTypes : $4, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
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
          | NEWPAIR OPENROUND expr COMMA expr CLOSEROUND   { $$ = NewPair{FstExpr : $3, SndExpr : $5, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input } }
          | CALL IDENTIFIER OPENROUND exprlist CLOSEROUND  { $$ = Call{Ident : $2, ParamList : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }

statements : statements SEMICOLON statement           { $$ = append($1,$3) }
           | statement                                { $$ = []Statement{$1} }


statement : SKIP                                        { $$ = Skip{Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input } }
          | typeDef IDENTIFIER ASSIGNMENT assignrhs     { $$ = Declare{DecType : $1, Lhs : $2, Rhs : $4, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input } }
          | assignlhs ASSIGNMENT assignrhs              { $$ = Assignment{Lhs : $1, Rhs : $3, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
          | READ assignlhs                              { $$ = Read{ &parserlex.(*Lexer).input, $<pos>1 , $2, } }
          | FREE expr                                   { $$ = Free{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | RETURN expr                                 { $$ = Return{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | EXIT expr                                   { $$ = Exit{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | PRINT expr                                  { $$ = Print{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | PRINTLN expr                                { $$ = Println{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | IF expr THEN statements ELSE statements FI  { $$ = If{Conditional : $2, ThenStat : $4, ElseStat : $6, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input } }
          | WHILE expr DO statements DONE               { $$ = While{Conditional : $2, DoStat : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input} }
          | BEGIN statements END                        { $$ = Scope{StatList : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input } }
          | error SEMICOLON                             {
                                                          parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }
          | error END                                   { parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }
          | error FI                                     {
                                                          parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }
          | error DONE                                   {
                                                          parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }

expr : INTEGER       { $$ =  $1 }

     | TRUE          { $$ =  $1 }
     | FALSE         { $$ =  $1 }

     | CHARACTER     { $$ =  $1 }
     | STRINGCONST   { $$ =  $1 }
     | pairliter     { $$ =  $1 }
     | IDENTIFIER    { $$ =  $1 }
     | arrayelem     { $$ =  $1 }

     | NOT expr     { $$ = Unop{Unary : NOT, Expr : $2, Pos : $<pos>1,FileText :&parserlex.(*Lexer).input   } }
     | LEN expr     { $$ = Unop{Unary : LEN, Expr : $2, Pos : $<pos>1,FileText :&parserlex.(*Lexer).input  } }
     | ORD expr     { $$ = Unop{Unary : ORD, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | CHR expr     { $$ = Unop{Unary : CHR, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | SUB expr     { $$ = Unop{Unary : SUB, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | PLUS expr    { $$ = $2 }


     | expr PLUS expr { $$ = Binop{Left : $1, Binary : PLUS, Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr SUB expr  { $$ = Binop{Left : $1, Binary : SUB, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr MUL expr  { $$ = Binop{Left : $1, Binary : MUL, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr MOD expr  { $$ = Binop{Left : $1, Binary : MOD, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr DIV expr  { $$ = Binop{Left : $1, Binary : DIV, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr LT expr   { $$ = Binop{Left : $1, Binary : LT, Right : $3,   Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr GT expr   { $$ = Binop{Left : $1, Binary : GT, Right : $3,   Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr LTE expr  { $$ = Binop{Left : $1, Binary : LTE, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr GTE expr  { $$ = Binop{Left : $1, Binary : GTE, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr EQ expr   { $$ = Binop{Left : $1, Binary : EQ, Right : $3,   Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr NEQ expr  { $$ = Binop{Left : $1, Binary : NEQ, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr AND expr  { $$ = Binop{Left : $1, Binary : AND, Right : $3,  Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr OR expr   { $$ = Binop{Left : $1, Binary : OR, Right : $3,   Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | OPENROUND expr CLOSEROUND  { $$ = $2 }


arrayliter : OPENSQUARE exprlist CLOSESQUARE { $$ = ArrayLiter{&parserlex.(*Lexer).input, $<pos>1, $2 } }

exprlist : exprlist COMMA expr {$$ = append($1, $3)}
         | expr                {$$ = []Evaluation{$1}}
         |                     {$$ = []Evaluation{}}

arrayelem : IDENTIFIER bracketed {$$ = ArrayElem{Ident: $1, Exprs : $2, Pos : $<pos>1,FileText :&parserlex.(*Lexer).input  } }

bracketed : bracketed OPENSQUARE expr CLOSESQUARE {$$ = append($1, $3)}
          | OPENSQUARE expr CLOSESQUARE {$$ = []Evaluation{$2}}


pairliter : NULL    { $$ =  PairLiter{} }


pairelem : FST expr { $$ = PairElem{Fsnd: Fst, Expr : $2, Pos : $<pos>1  } }
         | SND expr { $$ = PairElem{Fsnd: Snd, Expr : $2, Pos : $<pos>1  } }

pairtype : PAIR OPENROUND pairelemtype COMMA pairelemtype CLOSEROUND  { $$ = PairType{FstType : $3, SndType : $5} }

pairelemtype : basetype  { $$ = $1 }
             | arraytype { $$ = $1 }
             | PAIR      { $$ = Pair}

typeDef : basetype  { $$ =  $1 }
        | arraytype { $$ =  $1 }
        | pairtype  { $$ =  $1 }

basetype : INT      { $$ =  Int }
         | BOOL     { $$ =  Bool }
         | CHAR     { $$ =  Char }
         | STRING   { $$ =  String }

arraytype : typeDef OPENSQUARE CLOSESQUARE { $$ = ArrayType{Type : $1} }

%%