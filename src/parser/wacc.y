%{
package parser
%}

%union {
line int
  null string
  str  string
  prog *Program
  funct *Function
  functs []*Functions
  stat *Statement
  stats []*Statement
  param *Param
  params []*Param
  assignLHS *AssignLHS
  assignRHS *AssignRHS
  args []*Arg
  arg *Arg
  pairElem *PairElem
  typeDef *Type
  base_type *BaseType
  array_type *ArrayType
pair_type *PairType
pair_elem_type *PairElemType
expr *Expr
exprs []*Expr
aexprs []*Expr
unary_op *UnaryOp
binary_op *BinaryOp
ident *Ident
array_elem *ArrayElem
int_liter *IntLiter
int_sign *IntSign
bool_liter *BoolLiter
str_liter *StrLiter
pair_liter *PairLiter
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

// Nonterminal return types of actions

%type <prog>  program
%type <functs> functions
%type <funct>  function
%type <stats> statements
%type <stat>  statement
%type <param> param
%type <param> params
%type <param> paramList
%type <assignLHS> assign_lhs
%type <assignRHS> assign_rhs
%type <arg> arg
%type <args> arg_list
%type <pairElem> pair_elem
%type <typDefe> typeDef
%type <base_type> base_type
%type <array_type> array_type
%type <pair_type> pair_type
%type <expr> expr
%type <exprs> exprs
%type <unary_op> unary_op
%type <binary_op> binary_op
%type <ident> ident
%type <array_elem> array_elem
%type <int_liter> int_liter
%type <int_sign> int_sign
%type <bool_liter> bool_liter
%type <str_liter> str_liter
%type <exprs> aexprs
%type <pair_liter> pair_liter
%type <null> NULL


/* Associativity and Precedence operators specification start -----------*/

%left NOT NEG LEN ORD CHR
%left '*' '/' '%' '+' '-' '&&' '||'
%nonassoc '>' '>=' '<' '<=' '==' '!='

/* Associativity and Precedence operators end --------------------------------*/

/* Reserved Tokens End -------------------------------------------------------*/


%start program   // Start symbol

%%

program : BEGIN functions statements END
          {
          yylex.(*lexer).prog = &Program{Functions : $2, Statements : $3}
            return 0
          }
        ;

// CORRESPONDING ACTIONS HAVE TO BE IN GOOOO

functions : $$ = NULL
          | function
            { $$ = &Functions{$1} }
          | functions function
            { $$ = append($1, $2) }
          ;

function : typeDef ident '(' params ')' IS statement END
           { $$ = &Function{Type : $1, Ident : $2, Params : $4, Stat : $7} }
         ;



optional_param_list
    : param_list { $$.Params = $1.Params }
    |
    ;

param_list
    : param ',' param_list { $$.Params = append([]Param{$1.Param}, $3.Params...) }
    | param { $$.Params = []Param{$1.Param} }
    ;

param
    : type identifier { $$.Param = Param{$1.Position, $1.Type, $2.Expr.(*IdentExpr), $2.Position} }
    ;














params : param
         { $$ = []*Param{$1} }
       | param ',' paramList
         { $$ = append($1, $2) }
	   |
       ;

param : typeDef ident
        { $$ = &Param{Type : $1, Ident : $2} }
      ;

paramList : ',' param
            { $$ = []*Param{$2} }
          |
		  ;

statements : statement    // WHAT IF PROGRAM IS EMPTY ??????
             { $$ = []*Statements{$1} }
           | statements ';' statements   // double check for error here
             { $$ = append($1, $2) }
           ;

statement : SKIP    // NOT sure what to do with this
            { $$ = &Statement{Ident : $1} }
          | typeDef ident '=' assign_rhs
            { $$ = &Statement{Type : $1, Ident : $2, AssignRHS : $4} }
          | assign_lhs '=' assign_rhs
            { $$ = &Statement{AssignLHS : $1, AssignRHS : $3} }
          | READ assign_lhs
            { $$ = &Statement{AssignLHS : $2} }
          | FREE expr
            { $$ = &Statement{Expr : $2} }
          | RETURN expr
            { $$ = &Statement{Expr : $2} }
          | EXIT expr
            { $$ = &Statement{Expr : $2} }
          | PRINT expr
            { $$ = &Statement{Expr : $2} }
          | PRINTLN expr
            { $$ = &Statement{Expr : $2} }
          | IF expr THEN statement ELSE statement FI
            { $$ = &Statement{Left : $2, Right : $4, Else : $6} }
          | WHILE expr DO statement DONE
            { $$ = &Statement{Left : $2, Right : $4} }
          | BEGIN statement END
            { $$ = &Statement{NewScope : $2} }
          | statements
            { $$ = &Statement{Stat : $1} }  // NOT SURE IF ENTIRELY CORRECT WAY OF GOIN ABOUT THIs
          ;

assign_lhs : ident
             { $$ = &AssignLHS{Ident : $1} }
           | array_elem
             { $$ = &AssignLHS{ArrayElem : $1} }
           | pair_elem
             { $$ = &AssignLHS{PairElem : $1} }
           ;

assign_rhs : expr
             { $$ = &AssignRHS{Expr : $1} }
           | array_liter
             { $$ = &AssignRHS{ArrayLiter : $1} }
           | NEW_PAIR '(' expr ',' expr ')'
             { $$ = &AssignRHS{NewPairFst : $3, NewPairSnd : $5} }
           | pair_elem
             { $$ = &AssignRHS{PairElem : $1} }
           | CALL ident '(' arg_list ')
             { $$ = &AssignRHS{Ident : $2, ArgList : $4} }
           ;

arg_list : arg
           { $$ = []*Arg{$1} }
         | args ',' arg     // CHECK FOR ERRORS HERE
           { $$ = append($1, $2) }
         |
		 ;

arg : expr
      { $$ = &Arg{Expr : $1} }
      ;

pair_elem : FST expr
            { $$ = &PairElem{Fst : $1} }
          | SND expr
            { $$ = &PairElem{Snd : $1} }
          ;

type : base_type
       { $$ = &Type{Type : $1} }
     | array_type
       { $$ = &Type{Type : $1} }
     | pair_type
       { $$ = &Type{Type : $1} }
     ;

base_type : INT
            { $$ = &BaseType{$1} }
          | BOOL
            { $$ = &BaseType{$1} }
          | CHAR
            { $$ = &BaseType{$1} }
          | STRING
            { $$ = &BaseType{$1} }
          ;

array_type : typeDef '[' ']'
             { $$ = &ArrayType{Type : $1} }
           ;

pair_type : PAIR '(' pair_elem_type ',' pair_elem_type ')'
            { $$ = &PairType{Left : $3, Right : $5} }
          ;

pair_elem_type : base_type
                 { $$ = &PairElemType{PairType : $1} }
               | array_type
                 { $$ = &PairElemType{PairType : $1} }
               | PAIR
                 { $$ = &PairElemType{PairType : $1} }
               ;

expr : int_liter { $$ = &Expr{Expr : $1} }
     | bool_liter { $$ = &Expr{Expr : $1} }
  //   | char_liter { $$ = &Expr{Expr : $1} }
     | str_liter { $$ = &Expr{Expr : $1} }
     | pair_liter { $$ = &Expr{Expr : $1} }
     | ident { $$ = &Expr{Expr : $1} }
     | array_elem { $$ = &Expr{Expr : $1} }
     | unary_op expr { $$ = &Expr{UnaryOp : $1, Expr : $2} }
     | expr binary_op expr { $$ = &Expr{LExpr : $1, BinaryOp : $2, RExpr : $3} }
     | '(' expr ')' { $$ = &Expr{Expr : $2} }
     ;

unary_op : NOT { $$ = &UnaryOp{Op : $1} }
         | NEG { $$ = &UnaryOp{Op : $1} }
         | LEN { $$ = &UnaryOp{Op : $1} }
         | ORD { $$ = &UnaryOp{Op : $1} }
         | CHAR{ $$ = &UnaryOp{Op : $1} }
         ;

binary_op : MULT { $$ = &BinaryOp{Op : $1}}
          | DIV { $$ = &BinaryOp{Op : $1}}
          | MOD { $$ = &BinaryOp{Op : $1}}
          | ADD { $$ = &BinaryOp{Op : $1}}
          | SUB { $$ = &BinaryOp{Op : $1}}
          | AND { $$ = &BinaryOp{Op : $1}}
          | OR { $$ = &BinaryOp{Op : $1}}
          | GT { $$ = &BinaryOp{Op : $1}}
          | GTE { $$ = &BinaryOp{Op : $1}}
          | LT { $$ = &BinaryOp{Op : $1}}
          | LTE{ $$ = &BinaryOp{Op : $1}}
          | EQ { $$ = &BinaryOp{Op : $1}}
          | NEQ{ $$ = &BinaryOp{Op : $1}}
          ;

ident : IDENTIFIER { $$ = &Ident{Ident : $1} } ;

array_elem : ident '[' expr ']' { $$ = &ArrayElem{Ident : $1, Expr : $3} }
           | ident '[' exprs expr ']' { $$ = append($1, []*Expr{Ident : $1, Expr : $4}) }   // SOURCE OF ERRORS HERE
           ;

exprs : expr { $$ = []*Expr{Expr : $1}}
      | exprs expr { $$ = append($1, $2) }
      ;

int_liter : int_sign INTEGER { $$ = &IntLiter{Sign : $1, Num : $2} } ;

int_sign :  POSITIVE { $$ = &IntSign{Sign : $1} }
          | NEGATIVE { $$ = &IntSign{Sign : $1} }
          |
		  ;

bool_liter : TRUE { $$ = &BoolLiter{Bool : $1}}
           | FALSE { $$ = &BoolLiter{Bool : $1}}
           ;

str_liter : CHARACTER { &StrLiter{Char : $1} } ;

array_liter : '[' ']'
            | '[' aexprs expr ']'
            ;

aexprs : expr { $$ = []*Expr{Expr : $1}}
       | exprs expr { $$ = append($1, $2) }
       |
	   ;

pair_liter : NULL { $$ = &PairLiter{Val : $1} }
