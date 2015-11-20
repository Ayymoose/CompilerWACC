%{
package parser
%}

%union{
str string
functions  []*Function
function     *Function
statements []*Statement
statement    *Statement
params     []*Parameter
param        *Parameter
exprs       []*Expr
expr         *Expr
chars       []*Char
char        *Char
}

/* Grammar declaration -------------------------------------------------------*/

// Reserved Terminal Tokens recognized by Lexer

%token <str> BEGIN END                                     // Program delimiters
%token <str> IS
%token <str> SKIP
%token <str> READ FREE RETURN EXIT PRINT PRINTLN
%token <str> IF THEN ELSE FI                               // If statement
%token <str> WHILE DO DONE                                 // While statement
%token <str> NEW_PAIR
%token <str> CALL
%token <str> FST SND
%token <str> INT BOOL CHAR STRING                          // Base types
%token <str> PAIR
%token <str> IDENTIFIER
%token <str> NOT NEG LEN ORD CHR                           // Unary ops
%token <str> MULT DIV MOD ADD SUB AND OR                   // Binary ops
%token <str> GT GTE LT LTE EQ NEQ
%token <str> INTEGER
%token <str> POSITIVE NEGATIVE
%token <str> TRUE FALSE                                    // Booleans
%token <str> CHARACTER
%token <str> NULL

%type <functions> functions
%type <function>  function
//%type <statements> statements
%type <statement> statement
%type <typeDef>  typeDef
%type <typeDef>  base_type
%type <typeDef>  array_type
%type <typeDef>  pair_type
%type <typeDef>  pair_elem_type
%type <str> ident
%type <params> optional_param_list
%type <params> param_list
%type <param> param
%type <assign_rhs> assign_rhs
%type <assign_lhs> assign_lhs
%type <expr> expr
%type <expr> optional_expr_list
%type <expr> expr_list
%type <array_elem> array_elem
%type <pair_elem> pair_elem
%type <array_liter> array_liter
%type <arg> optional_arg_list
%type <arg> arg_list
%type <int> int_liter
%type <bool> bool_liter
%type <rune> char_liter
%type <str> str_liter
%type <pair> pair_liter
%type <uop> unaryOp
%type <bop> binaryOp
%type <sign> optional_sign
%type <chr> character
%type <chr> characters
%type <gte> '>='
%type <lte> '<='
%type <gt>  '>'
%type <lt>  '<' 
%type <not> '!='
%type <eq>  '=='


/* Associativity and Precedence operators specification start -----------*/

%left     NOT NEG LEN ORD CHR
%left     '*' '/' '%' '+' '-' '&&' '||'
%nonassoc '>' '>='
%nonassoc  '<' '<=' '==' '!='

/* Associativity and Precedence operators end --------------------------------*/

%%

program : BEGIN functions statement END
          { yylex.(*lexer).prog = &Program{Functions : $2, Statements : $3}
            return 0
          }
        ;

functions : function functions  { $$ = append([]Function{$1}, $2...) }
          | function            { $$ = []Function{$1} }
          | { $$ } // IS THIS EVEN LEGAL????
          ;

function : typeDef ident '(' optional_param_list ')' IS statement END
           { $$ = &Function{Type : $1, Ident : $2, Params : $4, Stat : $7} }
         ; 

optional_param_list : param_list { $$ = $1 }
                    | { $$ } // IS THIS EVEN LEGAL????
                    ;

param_list : param ',' param_list { $$ = append([]Parameter{$1}, $3...) }
           | param                { $$ = []Parameter{$1} } 
           ;

param : typeDef ident { $$ = Parameter{Type : $1, Ident : $2} }
      ;

          // SKIP ???
statement : typeDef ident '=' assign_rhs { $$ = Statement{Type : $1, Ident : $2, AssignRHS : $4} }
          | assign_lhs '=' assign_rhs    { $$ = Statement{AssignLHS : $1, AssignRHS : $3} }
          | READ assign_lhs              { $$ = Statement{AssignLHS : $2} }
          | FREE expr                    { $$ = Statement{Expr : $2} }
          | RETURN expr                  { $$ = Statement{Expr : $2} }
          | EXIT expr                    { $$ = Statement{Expr : $2} }
          | PRINT expr                   { $$ = Statement{Expr : $2} } 
          | PRINTLN expr                 { $$ = Statement{Expr : $2} } 
          | IF expr THEN statement ELSE statement FI  { $$ = Statement{Expr : $2, StatL : $4, StatR : $6} }
          | WHILE expr DO statement DONE { $$ = Statement{Expr : $2, Stat : $4} } 
          | BEGIN statement END          { $$ = Statement{Stat : $2} }
          | statement ';' statement      { $$ = append([]Statement{$1}, $3...) }  
          ; 

assign_lhs : ident      { $$ = AssignLHS{Ident : $1} }
           | array_elem { $$ = AssignLHS{ArrayElem : $1} }
           | pair_elem  { $$ = AssignLHS{PairElem :$1} }
           ;

 assign_rhs : expr                                  { $$ = AssignRHS{Expr : $1} }
            | array_liter                           { $$ = AssignRHS{ArrayLiter : $1} }
            | NEW_PAIR '(' expr ',' expr ')'        { $$ = AssignRHS{ExprL : $3, ExprR : $5} } 
            | pair_elem                             { $$ = AssignRHS{PairElem : $1} }
            | CALL ident '(' optional_arg_list ')'  { $$ = AssignRHS{Ident : $2, Args : $4} }
            ; 

optional_arg_list : arg_list { $$ = $1 }
                  | { $$ } // IS THIS EVAN LEGAL????
                  ;

arg_list : expr optional_expr_list { $$ = append([]Expr{$1}, $2...) }
         ;

optional_expr_list : ',' expr { $$ = Expr{Expr : $2} }
                   ; 

pair_elem : FST expr { $$ = PairElem{Expr : $2} }
          | SND expr { $$ = PairElem{Expr : $2} }
          ;

typeDef : base_type  { $$ = TypeDef{Type : $1} }
        | array_type { $$ = TypeDef{Type : $1} }
        | pair_type  { $$ = TypeDef{Type : $1} }
        ;
    
base_type : INT    { $$ = BaseType{Type : $1} } 
          | BOOL   { $$ = BaseType{Type : $1} } 
          | CHAR   { $$ = BaseType{Type : $1} } 
          | STRING { $$ = BaseType{Type : $1} } 
          ;

array_type : typeDef '[' ']' { $$ = ArrayType{Type : $1} }
           ;

pair_type : PAIR '(' pair_elem_type ',' pair_elem_type ')'  { $$ = PairType{Lpair : $3, Rpair : $5} }
          ;

pair_elem_type : base_type  { $$ = PairElemType{Type : $1} } 
               | array_type { $$ = PairElemType{Type : $1} } 
               | PAIR       { $$ = PairElemType{Type : $1} }   
               ;

expr : int_liter    { $$ = Expr{Expr : $1} } 
     | bool_liter   { $$ = Expr{Expr : $1} } 
     | char_liter   { $$ = Expr{Expr : $1} } 
     | str_liter    { $$ = Expr{Expr : $1} } 
     | pair_liter   { $$ = Expr{Expr : $1} } 
     | ident        { $$ = Expr{Expr : $1} } 
     | array_elem   { $$ = Expr{Expr : $1} } 
   //  | '!' expr     { $$ = Expr{UnaryOp : $1, Expr : $2} }
   //  | '-' expr     { $$ = Expr{UnaryOp : $1, Expr : $2} }
   //  | '  

//%left     NOT NEG LEN ORD CHR
//%left     '*' '/' '%' '+' '-' '&&' '||'
//%nonassoc '>' '>='
//%nonassoc  '<' '<=' '==' '!='


     | unaryOp expr { $$ = Expr{UnaryOp : $1, Expr : $2} }
     | expr '!=' expr { $$ = Expr{Lexpr : $1, BinaryOp : $2, Rexpr : $3} }
     | expr '==' expr { $$ = Expr{Lexpr : $1, BinaryOp : $2, Rexpr : $3} }
     | expr '>=' expr { $$ = Expr{Lexpr : $1, BinaryOp : $2, Rexpr : $3} }
     | expr '<=' expr { $$ = Expr{Lexpr : $1, BinaryOp : $2, Rexpr : $3} }
     | expr '>' expr { $$ = Expr{Lexpr : $1, BinaryOp : $2, Rexpr : $3} }
     | expr '<' expr { $$ = Expr{Lexpr : $1, BinaryOp : $2, Rexpr : $3} }
 
     | expr binaryOp expr { $$ = Expr{Lexpr : $1, BinaryOp : $2, Rexpr : $3} }
     | '(' expr ')'  { $$ = Expr{Expr : $2} } 
     ;

unaryOp : NOT { $$ = UnaryOp{Op : $1} }
        | NEG { $$ = UnaryOp{Op : $1} }
        | LEN { $$ = UnaryOp{Op : $1} }
        | ORD { $$ = UnaryOp{Op : $1} }
        | CHR { $$ = UnaryOp{Op : $1} }
        ;

binaryOp : MULT { $$ = BinaryOp{Op : $1} }
         | DIV  { $$ = BinaryOp{Op : $1} }
         | MOD  { $$ = BinaryOp{Op : $1} }
         | ADD  { $$ = BinaryOp{Op : $1} }
         | SUB  { $$ = BinaryOp{Op : $1} }
         | AND  { $$ = BinaryOp{Op : $1} }
         | OR   { $$ = BinaryOp{Op : $1} }
         ;

ident : IDENTIFIER ;

array_elem : ident expr_list { $$ = ArrayElem{Ident : $1, Exprs : $2} }
           ;

// NOT SURE ABOUT THIS IMPLEMENTATION>>
expr_list : '[' expr expr_list ']' { $$ = append([]Expr{$2}, $3...) }
          | '[' expr ']'       { $$ = []Expr{$2} } 
          ;

int_liter : optional_sign INTEGER { $$ = IntLiter{Sign : $1, Int : $2} }
          ;

optional_sign : POSITIVE { $$ = Sign{Sign : $1} }
              | NEGATIVE { $$ = Sign{Sign : $1} }
              | { $$ }  // IS THIS EVAN LEGALL
              ;

bool_liter : TRUE  { $$ = Bool{Bool : $1} }
           | FALSE { $$ = Bool{Bool : $1} }
           ; 

char_liter : '’' character  '’' { $$ = Char{Char : $2} }
           ;
 
str_liter : '"' characters '"' { $$ = Str{Chars : $2} }
          ;

characters : character characters { $$ = append([]Char{$1}, $2...) } 
           | character            { $$ = []Char{$1} } 
           ;

character : CHARACTER { $$ = Char{Char : $1} } ;

array_liter : '[' expr optional_expr_list ']'  { $$ = append([]Expr{$2}, $3...)}
            ;

pair_liter : NULL { $$ = PairLiter{PairLit : $1} }
