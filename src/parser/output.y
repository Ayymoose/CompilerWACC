
state 0
	$accept: .program $end 

	BEGIN  shift 2
	.  error

	program  goto 1

state 1
	$accept:  program.$end 

	$end  accept
	.  error


state 2
	program:  BEGIN.functions statements END 
	functions: .    (3)

	.  reduce 3 (src line 85)

	functions  goto 3

state 3
	program:  BEGIN functions.statements END 
	functions:  functions.function 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	function  goto 5
	statement  goto 6
	statements  goto 4
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 7
	arraytype  goto 20
	pairtype  goto 21

state 4
	program:  BEGIN functions statements.END 
	statements:  statements.SEMICOLON statement 

	END  shift 32
	SEMICOLON  shift 33
	.  error


state 5
	functions:  functions function.    (2)

	.  reduce 2 (src line 84)


state 6
	statements:  statement.    (34)

	.  reduce 34 (src line 130)


state 7
	function:  typeDef.IDENTIFIER OPENROUND CLOSEROUND IS funcstats END 
	function:  typeDef.IDENTIFIER OPENROUND paramlist CLOSEROUND IS funcstats END 
	statement:  typeDef.IDENTIFIER ASSIGNMENT assignrhs 
	arraytype:  typeDef.OPENSQUARE CLOSESQUARE 

	OPENSQUARE  shift 35
	IDENTIFIER  shift 34
	.  error


state 8
	statement:  SKIP.    (35)

	.  reduce 35 (src line 132)


state 9
	statement:  assignlhs.ASSIGNMENT assignrhs 

	ASSIGNMENT  shift 36
	.  error


state 10
	statement:  READ.assignlhs 

	FST  shift 30
	SND  shift 31
	IDENTIFIER  shift 22
	.  error

	assignlhs  goto 37
	pairelem  goto 24
	arrayelem  goto 23

state 11
	statement:  FREE.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 38
	arrayelem  goto 46
	pairliter  goto 44

state 12
	statement:  RETURN.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 55
	arrayelem  goto 46
	pairliter  goto 44

state 13
	statement:  EXIT.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 56
	arrayelem  goto 46
	pairliter  goto 44

state 14
	statement:  PRINT.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 57
	arrayelem  goto 46
	pairliter  goto 44

state 15
	statement:  PRINTLN.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 58
	arrayelem  goto 46
	pairliter  goto 44

state 16
	statement:  IF.expr THEN statements ELSE statements FI 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 59
	arrayelem  goto 46
	pairliter  goto 44

state 17
	statement:  WHILE.expr DO statements DONE 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 60
	arrayelem  goto 46
	pairliter  goto 44

state 18
	statement:  BEGIN.statements END 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	statement  goto 6
	statements  goto 61
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 62
	arraytype  goto 20
	pairtype  goto 21

state 19
	typeDef:  basetype.    (89)

	.  reduce 89 (src line 197)


state 20
	typeDef:  arraytype.    (90)

	.  reduce 90 (src line 198)


state 21
	typeDef:  pairtype.    (91)

	.  reduce 91 (src line 199)


state 22
	assignlhs:  IDENTIFIER.    (25)
	arrayelem:  IDENTIFIER.bracketed 

	OPENSQUARE  shift 64
	.  reduce 25 (src line 119)

	bracketed  goto 63

state 23
	assignlhs:  arrayelem.    (26)

	.  reduce 26 (src line 120)


state 24
	assignlhs:  pairelem.    (27)

	.  reduce 27 (src line 121)


state 25
	basetype:  INT.    (92)

	.  reduce 92 (src line 201)


state 26
	basetype:  BOOL.    (93)

	.  reduce 93 (src line 202)


state 27
	basetype:  CHAR.    (94)

	.  reduce 94 (src line 203)


state 28
	basetype:  STRING.    (95)

	.  reduce 95 (src line 204)


state 29
	pairtype:  PAIR.OPENROUND pairelemtype COMMA pairelemtype CLOSEROUND 

	OPENROUND  shift 65
	.  error


state 30
	pairelem:  FST.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 66
	arrayelem  goto 46
	pairliter  goto 44

state 31
	pairelem:  SND.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 67
	arrayelem  goto 46
	pairliter  goto 44

state 32
	program:  BEGIN functions statements END.    (1)

	.  reduce 1 (src line 80)


state 33
	statements:  statements SEMICOLON.statement 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	statement  goto 68
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 62
	arraytype  goto 20
	pairtype  goto 21

state 34
	function:  typeDef IDENTIFIER.OPENROUND CLOSEROUND IS funcstats END 
	function:  typeDef IDENTIFIER.OPENROUND paramlist CLOSEROUND IS funcstats END 
	statement:  typeDef IDENTIFIER.ASSIGNMENT assignrhs 

	OPENROUND  shift 69
	ASSIGNMENT  shift 70
	.  error


state 35
	arraytype:  typeDef OPENSQUARE.CLOSESQUARE 

	CLOSESQUARE  shift 71
	.  error


state 36
	statement:  assignlhs ASSIGNMENT.assignrhs 

	NEWPAIR  shift 76
	CALL  shift 77
	FST  shift 30
	SND  shift 31
	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENSQUARE  shift 78
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	assignrhs  goto 72
	expr  goto 73
	pairelem  goto 75
	arrayliter  goto 74
	arrayelem  goto 46
	pairliter  goto 44

state 37
	statement:  READ assignlhs.    (38)

	.  reduce 38 (src line 135)


state 38
	statement:  FREE expr.    (39)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 39 (src line 136)


state 39
	expr:  INTEGER.    (47)

	.  reduce 47 (src line 145)


state 40
	expr:  TRUE.    (48)

	.  reduce 48 (src line 146)


state 41
	expr:  FALSE.    (49)

	.  reduce 49 (src line 147)


state 42
	expr:  CHARACTER.    (50)

	.  reduce 50 (src line 148)


state 43
	expr:  STRINGCONST.    (51)

	.  reduce 51 (src line 149)


state 44
	expr:  pairliter.    (52)

	.  reduce 52 (src line 150)


state 45
	expr:  IDENTIFIER.    (53)
	arrayelem:  IDENTIFIER.bracketed 

	OPENSQUARE  shift 64
	.  reduce 53 (src line 151)

	bracketed  goto 63

state 46
	expr:  arrayelem.    (54)

	.  reduce 54 (src line 152)


state 47
	expr:  NOT.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 92
	arrayelem  goto 46
	pairliter  goto 44

state 48
	expr:  LEN.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 93
	arrayelem  goto 46
	pairliter  goto 44

state 49
	expr:  ORD.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 94
	arrayelem  goto 46
	pairliter  goto 44

state 50
	expr:  CHR.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 95
	arrayelem  goto 46
	pairliter  goto 44

state 51
	expr:  SUB.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 96
	arrayelem  goto 46
	pairliter  goto 44

state 52
	expr:  PLUS.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 97
	arrayelem  goto 46
	pairliter  goto 44

state 53
	expr:  OPENROUND.expr CLOSEROUND 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 98
	arrayelem  goto 46
	pairliter  goto 44

state 54
	pairliter:  NULL.    (82)

	.  reduce 82 (src line 186)


state 55
	statement:  RETURN expr.    (40)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 40 (src line 137)


state 56
	statement:  EXIT expr.    (41)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 41 (src line 138)


state 57
	statement:  PRINT expr.    (42)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 42 (src line 139)


state 58
	statement:  PRINTLN expr.    (43)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 43 (src line 140)


state 59
	statement:  IF expr.THEN statements ELSE statements FI 
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	THEN  shift 99
	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  error


state 60
	statement:  WHILE expr.DO statements DONE 
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	DO  shift 100
	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  error


state 61
	statements:  statements.SEMICOLON statement 
	statement:  BEGIN statements.END 

	END  shift 101
	SEMICOLON  shift 33
	.  error


state 62
	statement:  typeDef.IDENTIFIER ASSIGNMENT assignrhs 
	arraytype:  typeDef.OPENSQUARE CLOSESQUARE 

	OPENSQUARE  shift 35
	IDENTIFIER  shift 102
	.  error


state 63
	arrayelem:  IDENTIFIER bracketed.    (79)
	bracketed:  bracketed.OPENSQUARE expr CLOSESQUARE 

	OPENSQUARE  shift 103
	.  reduce 79 (src line 181)


state 64
	bracketed:  OPENSQUARE.expr CLOSESQUARE 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 104
	arrayelem  goto 46
	pairliter  goto 44

state 65
	pairtype:  PAIR OPENROUND.pairelemtype COMMA pairelemtype CLOSEROUND 

	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 108
	.  error

	basetype  goto 106
	typeDef  goto 109
	arraytype  goto 107
	pairtype  goto 21
	pairelemtype  goto 105

state 66
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	pairelem:  FST expr.    (83)

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 83 (src line 188)


state 67
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	pairelem:  SND expr.    (84)

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 84 (src line 189)


state 68
	statements:  statements SEMICOLON statement.    (33)

	.  reduce 33 (src line 129)


state 69
	function:  typeDef IDENTIFIER OPENROUND.CLOSEROUND IS funcstats END 
	function:  typeDef IDENTIFIER OPENROUND.paramlist CLOSEROUND IS funcstats END 

	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	CLOSEROUND  shift 110
	.  error

	paramlist  goto 111
	param  goto 112
	basetype  goto 19
	typeDef  goto 113
	arraytype  goto 20
	pairtype  goto 21

state 70
	statement:  typeDef IDENTIFIER ASSIGNMENT.assignrhs 

	NEWPAIR  shift 76
	CALL  shift 77
	FST  shift 30
	SND  shift 31
	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENSQUARE  shift 78
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	assignrhs  goto 114
	expr  goto 73
	pairelem  goto 75
	arrayliter  goto 74
	arrayelem  goto 46
	pairliter  goto 44

state 71
	arraytype:  typeDef OPENSQUARE CLOSESQUARE.    (96)

	.  reduce 96 (src line 206)


state 72
	statement:  assignlhs ASSIGNMENT assignrhs.    (37)

	.  reduce 37 (src line 134)


state 73
	assignrhs:  expr.    (28)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 28 (src line 123)


state 74
	assignrhs:  arrayliter.    (29)

	.  reduce 29 (src line 124)


state 75
	assignrhs:  pairelem.    (30)

	.  reduce 30 (src line 125)


state 76
	assignrhs:  NEWPAIR.OPENROUND expr COMMA expr CLOSEROUND 

	OPENROUND  shift 115
	.  error


state 77
	assignrhs:  CALL.IDENTIFIER OPENROUND exprlist CLOSEROUND 

	IDENTIFIER  shift 116
	.  error


state 78
	arrayliter:  OPENSQUARE.exprlist CLOSESQUARE 
	exprlist: .    (78)

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  reduce 78 (src line 179)

	expr  goto 118
	exprlist  goto 117
	arrayelem  goto 46
	pairliter  goto 44

state 79
	expr:  expr PLUS.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 119
	arrayelem  goto 46
	pairliter  goto 44

state 80
	expr:  expr SUB.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 120
	arrayelem  goto 46
	pairliter  goto 44

state 81
	expr:  expr MUL.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 121
	arrayelem  goto 46
	pairliter  goto 44

state 82
	expr:  expr MOD.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 122
	arrayelem  goto 46
	pairliter  goto 44

state 83
	expr:  expr DIV.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 123
	arrayelem  goto 46
	pairliter  goto 44

state 84
	expr:  expr LT.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 124
	arrayelem  goto 46
	pairliter  goto 44

state 85
	expr:  expr GT.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 125
	arrayelem  goto 46
	pairliter  goto 44

state 86
	expr:  expr LTE.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 126
	arrayelem  goto 46
	pairliter  goto 44

state 87
	expr:  expr GTE.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 127
	arrayelem  goto 46
	pairliter  goto 44

state 88
	expr:  expr EQ.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 128
	arrayelem  goto 46
	pairliter  goto 44

state 89
	expr:  expr NEQ.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 129
	arrayelem  goto 46
	pairliter  goto 44

state 90
	expr:  expr AND.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 130
	arrayelem  goto 46
	pairliter  goto 44

state 91
	expr:  expr OR.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 131
	arrayelem  goto 46
	pairliter  goto 44

state 92
	expr:  NOT expr.    (55)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 55 (src line 153)


state 93
	expr:  LEN expr.    (56)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 56 (src line 154)


state 94
	expr:  ORD expr.    (57)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 57 (src line 155)


state 95
	expr:  CHR expr.    (58)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 58 (src line 156)


state 96
	expr:  SUB expr.    (59)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 59 (src line 157)


state 97
	expr:  PLUS expr.    (60)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 60 (src line 158)


state 98
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	expr:  OPENROUND expr.CLOSEROUND 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	CLOSEROUND  shift 132
	.  error


state 99
	statement:  IF expr THEN.statements ELSE statements FI 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	statement  goto 6
	statements  goto 133
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 62
	arraytype  goto 20
	pairtype  goto 21

state 100
	statement:  WHILE expr DO.statements DONE 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	statement  goto 6
	statements  goto 134
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 62
	arraytype  goto 20
	pairtype  goto 21

state 101
	statement:  BEGIN statements END.    (46)

	.  reduce 46 (src line 143)


state 102
	statement:  typeDef IDENTIFIER.ASSIGNMENT assignrhs 

	ASSIGNMENT  shift 70
	.  error


state 103
	bracketed:  bracketed OPENSQUARE.expr CLOSESQUARE 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 135
	arrayelem  goto 46
	pairliter  goto 44

state 104
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	bracketed:  OPENSQUARE expr.CLOSESQUARE 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	CLOSESQUARE  shift 136
	.  error


state 105
	pairtype:  PAIR OPENROUND pairelemtype.COMMA pairelemtype CLOSEROUND 

	COMMA  shift 137
	.  error


state 106
	pairelemtype:  basetype.    (86)
	typeDef:  basetype.    (89)

	OPENSQUARE  reduce 89 (src line 197)
	.  reduce 86 (src line 193)


state 107
	pairelemtype:  arraytype.    (87)
	typeDef:  arraytype.    (90)

	OPENSQUARE  reduce 90 (src line 198)
	.  reduce 87 (src line 194)


state 108
	pairtype:  PAIR.OPENROUND pairelemtype COMMA pairelemtype CLOSEROUND 
	pairelemtype:  PAIR.    (88)

	OPENROUND  shift 65
	.  reduce 88 (src line 195)


state 109
	arraytype:  typeDef.OPENSQUARE CLOSESQUARE 

	OPENSQUARE  shift 35
	.  error


state 110
	function:  typeDef IDENTIFIER OPENROUND CLOSEROUND.IS funcstats END 

	IS  shift 138
	.  error


state 111
	function:  typeDef IDENTIFIER OPENROUND paramlist.CLOSEROUND IS funcstats END 
	paramlist:  paramlist.COMMA param 

	CLOSEROUND  shift 139
	COMMA  shift 140
	.  error


state 112
	paramlist:  param.    (23)

	.  reduce 23 (src line 115)


state 113
	param:  typeDef.IDENTIFIER 
	arraytype:  typeDef.OPENSQUARE CLOSESQUARE 

	OPENSQUARE  shift 35
	IDENTIFIER  shift 141
	.  error


state 114
	statement:  typeDef IDENTIFIER ASSIGNMENT assignrhs.    (36)

	.  reduce 36 (src line 133)


state 115
	assignrhs:  NEWPAIR OPENROUND.expr COMMA expr CLOSEROUND 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 142
	arrayelem  goto 46
	pairliter  goto 44

state 116
	assignrhs:  CALL IDENTIFIER.OPENROUND exprlist CLOSEROUND 

	OPENROUND  shift 143
	.  error


state 117
	arrayliter:  OPENSQUARE exprlist.CLOSESQUARE 
	exprlist:  exprlist.COMMA expr 

	CLOSESQUARE  shift 144
	COMMA  shift 145
	.  error


state 118
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	exprlist:  expr.    (77)

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 77 (src line 178)


state 119
	expr:  expr.PLUS expr 
	expr:  expr PLUS expr.    (61)
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 61 (src line 160)


state 120
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr SUB expr.    (62)
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 62 (src line 161)


state 121
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr MUL expr.    (63)
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 63 (src line 162)


state 122
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr MOD expr.    (64)
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 64 (src line 163)


state 123
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr DIV expr.    (65)
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 65 (src line 164)


state 124
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr LT expr.    (66)
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 66 (src line 165)


state 125
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr GT expr.    (67)
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 67 (src line 166)


state 126
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr LTE expr.    (68)
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 68 (src line 167)


state 127
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr GTE expr.    (69)
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 69 (src line 168)


state 128
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr EQ expr.    (70)
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 70 (src line 169)


state 129
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr NEQ expr.    (71)
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	.  reduce 71 (src line 170)


state 130
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr AND expr.    (72)
	expr:  expr.OR expr 

	.  reduce 72 (src line 171)


state 131
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	expr:  expr OR expr.    (73)

	.  reduce 73 (src line 172)


state 132
	expr:  OPENROUND expr CLOSEROUND.    (74)

	.  reduce 74 (src line 173)


state 133
	statements:  statements.SEMICOLON statement 
	statement:  IF expr THEN statements.ELSE statements FI 

	ELSE  shift 146
	SEMICOLON  shift 33
	.  error


state 134
	statements:  statements.SEMICOLON statement 
	statement:  WHILE expr DO statements.DONE 

	DONE  shift 147
	SEMICOLON  shift 33
	.  error


state 135
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	bracketed:  bracketed OPENSQUARE expr.CLOSESQUARE 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	CLOSESQUARE  shift 148
	.  error


state 136
	bracketed:  OPENSQUARE expr CLOSESQUARE.    (81)

	.  reduce 81 (src line 184)


state 137
	pairtype:  PAIR OPENROUND pairelemtype COMMA.pairelemtype CLOSEROUND 

	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 108
	.  error

	basetype  goto 106
	typeDef  goto 109
	arraytype  goto 107
	pairtype  goto 21
	pairelemtype  goto 149

state 138
	function:  typeDef IDENTIFIER OPENROUND CLOSEROUND IS.funcstats END 

	BEGIN  shift 164
	SKIP  shift 155
	READ  shift 158
	FREE  shift 159
	RETURN  shift 153
	EXIT  shift 160
	PRINT  shift 161
	PRINTLN  shift 162
	IF  shift 152
	WHILE  shift 163
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	prefuncstat  goto 154
	prefuncstats  goto 151
	funcstats  goto 150
	assignlhs  goto 157
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 156
	arraytype  goto 20
	pairtype  goto 21

state 139
	function:  typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND.IS funcstats END 

	IS  shift 165
	.  error


state 140
	paramlist:  paramlist COMMA.param 

	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	.  error

	param  goto 166
	basetype  goto 19
	typeDef  goto 113
	arraytype  goto 20
	pairtype  goto 21

state 141
	param:  typeDef IDENTIFIER.    (24)

	.  reduce 24 (src line 117)


state 142
	assignrhs:  NEWPAIR OPENROUND expr.COMMA expr CLOSEROUND 
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	COMMA  shift 167
	.  error


state 143
	assignrhs:  CALL IDENTIFIER OPENROUND.exprlist CLOSEROUND 
	exprlist: .    (78)

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  reduce 78 (src line 179)

	expr  goto 118
	exprlist  goto 168
	arrayelem  goto 46
	pairliter  goto 44

state 144
	arrayliter:  OPENSQUARE exprlist CLOSESQUARE.    (75)

	.  reduce 75 (src line 175)


state 145
	exprlist:  exprlist COMMA.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 169
	arrayelem  goto 46
	pairliter  goto 44

state 146
	statement:  IF expr THEN statements ELSE.statements FI 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	statement  goto 6
	statements  goto 170
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 62
	arraytype  goto 20
	pairtype  goto 21

state 147
	statement:  WHILE expr DO statements DONE.    (45)

	.  reduce 45 (src line 142)


state 148
	bracketed:  bracketed OPENSQUARE expr CLOSESQUARE.    (80)

	.  reduce 80 (src line 183)


state 149
	pairtype:  PAIR OPENROUND pairelemtype COMMA pairelemtype.CLOSEROUND 

	CLOSEROUND  shift 171
	.  error


state 150
	function:  typeDef IDENTIFIER OPENROUND CLOSEROUND IS funcstats.END 

	END  shift 172
	.  error


state 151
	funcstats:  prefuncstats.SEMICOLON RETURN expr 
	funcstats:  prefuncstats.SEMICOLON IF expr THEN funcstats ELSE funcstats FI 
	prefuncstats:  prefuncstats.SEMICOLON prefuncstat 

	SEMICOLON  shift 173
	.  error


state 152
	funcstats:  IF.expr THEN funcstats ELSE funcstats FI 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 174
	arrayelem  goto 46
	pairliter  goto 44

state 153
	funcstats:  RETURN.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 175
	arrayelem  goto 46
	pairliter  goto 44

state 154
	prefuncstats:  prefuncstat.    (21)

	.  reduce 21 (src line 112)


state 155
	prefuncstat:  SKIP.    (10)

	.  reduce 10 (src line 100)


state 156
	prefuncstat:  typeDef.IDENTIFIER ASSIGNMENT assignrhs 
	arraytype:  typeDef.OPENSQUARE CLOSESQUARE 

	OPENSQUARE  shift 35
	IDENTIFIER  shift 176
	.  error


state 157
	prefuncstat:  assignlhs.ASSIGNMENT assignrhs 

	ASSIGNMENT  shift 177
	.  error


state 158
	prefuncstat:  READ.assignlhs 

	FST  shift 30
	SND  shift 31
	IDENTIFIER  shift 22
	.  error

	assignlhs  goto 178
	pairelem  goto 24
	arrayelem  goto 23

state 159
	prefuncstat:  FREE.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 179
	arrayelem  goto 46
	pairliter  goto 44

state 160
	prefuncstat:  EXIT.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 180
	arrayelem  goto 46
	pairliter  goto 44

state 161
	prefuncstat:  PRINT.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 181
	arrayelem  goto 46
	pairliter  goto 44

state 162
	prefuncstat:  PRINTLN.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 182
	arrayelem  goto 46
	pairliter  goto 44

state 163
	prefuncstat:  WHILE.expr DO statements DONE 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 183
	arrayelem  goto 46
	pairliter  goto 44

state 164
	prefuncstat:  BEGIN.statements END 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	statement  goto 6
	statements  goto 184
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 62
	arraytype  goto 20
	pairtype  goto 21

state 165
	function:  typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND IS.funcstats END 

	BEGIN  shift 164
	SKIP  shift 155
	READ  shift 158
	FREE  shift 159
	RETURN  shift 153
	EXIT  shift 160
	PRINT  shift 161
	PRINTLN  shift 162
	IF  shift 152
	WHILE  shift 163
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	prefuncstat  goto 154
	prefuncstats  goto 151
	funcstats  goto 185
	assignlhs  goto 157
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 156
	arraytype  goto 20
	pairtype  goto 21

state 166
	paramlist:  paramlist COMMA param.    (22)

	.  reduce 22 (src line 114)


state 167
	assignrhs:  NEWPAIR OPENROUND expr COMMA.expr CLOSEROUND 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 186
	arrayelem  goto 46
	pairliter  goto 44

state 168
	assignrhs:  CALL IDENTIFIER OPENROUND exprlist.CLOSEROUND 
	exprlist:  exprlist.COMMA expr 

	CLOSEROUND  shift 187
	COMMA  shift 145
	.  error


state 169
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 
	exprlist:  exprlist COMMA expr.    (76)

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 76 (src line 177)


state 170
	statements:  statements.SEMICOLON statement 
	statement:  IF expr THEN statements ELSE statements.FI 

	FI  shift 188
	SEMICOLON  shift 33
	.  error


state 171
	pairtype:  PAIR OPENROUND pairelemtype COMMA pairelemtype CLOSEROUND.    (85)

	.  reduce 85 (src line 191)


state 172
	function:  typeDef IDENTIFIER OPENROUND CLOSEROUND IS funcstats END.    (4)

	.  reduce 4 (src line 87)


state 173
	funcstats:  prefuncstats SEMICOLON.RETURN expr 
	funcstats:  prefuncstats SEMICOLON.IF expr THEN funcstats ELSE funcstats FI 
	prefuncstats:  prefuncstats SEMICOLON.prefuncstat 

	BEGIN  shift 164
	SKIP  shift 155
	READ  shift 158
	FREE  shift 159
	RETURN  shift 189
	EXIT  shift 160
	PRINT  shift 161
	PRINTLN  shift 162
	IF  shift 190
	WHILE  shift 163
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	prefuncstat  goto 191
	assignlhs  goto 157
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 156
	arraytype  goto 20
	pairtype  goto 21

state 174
	funcstats:  IF expr.THEN funcstats ELSE funcstats FI 
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	THEN  shift 192
	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  error


state 175
	funcstats:  RETURN expr.    (9)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 9 (src line 98)


state 176
	prefuncstat:  typeDef IDENTIFIER.ASSIGNMENT assignrhs 

	ASSIGNMENT  shift 193
	.  error


state 177
	prefuncstat:  assignlhs ASSIGNMENT.assignrhs 

	NEWPAIR  shift 76
	CALL  shift 77
	FST  shift 30
	SND  shift 31
	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENSQUARE  shift 78
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	assignrhs  goto 194
	expr  goto 73
	pairelem  goto 75
	arrayliter  goto 74
	arrayelem  goto 46
	pairliter  goto 44

state 178
	prefuncstat:  READ assignlhs.    (13)

	.  reduce 13 (src line 103)


state 179
	prefuncstat:  FREE expr.    (14)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 14 (src line 104)


state 180
	prefuncstat:  EXIT expr.    (15)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 15 (src line 105)


state 181
	prefuncstat:  PRINT expr.    (16)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 16 (src line 106)


state 182
	prefuncstat:  PRINTLN expr.    (17)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 17 (src line 107)


state 183
	prefuncstat:  WHILE expr.DO statements DONE 
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	DO  shift 195
	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  error


state 184
	prefuncstat:  BEGIN statements.END 
	statements:  statements.SEMICOLON statement 

	END  shift 196
	SEMICOLON  shift 33
	.  error


state 185
	function:  typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND IS funcstats.END 

	END  shift 197
	.  error


state 186
	assignrhs:  NEWPAIR OPENROUND expr COMMA expr.CLOSEROUND 
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	CLOSEROUND  shift 198
	.  error


state 187
	assignrhs:  CALL IDENTIFIER OPENROUND exprlist CLOSEROUND.    (32)

	.  reduce 32 (src line 127)


state 188
	statement:  IF expr THEN statements ELSE statements FI.    (44)

	.  reduce 44 (src line 141)


state 189
	funcstats:  prefuncstats SEMICOLON RETURN.expr 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 199
	arrayelem  goto 46
	pairliter  goto 44

state 190
	funcstats:  prefuncstats SEMICOLON IF.expr THEN funcstats ELSE funcstats FI 

	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	expr  goto 200
	arrayelem  goto 46
	pairliter  goto 44

state 191
	prefuncstats:  prefuncstats SEMICOLON prefuncstat.    (20)

	.  reduce 20 (src line 111)


state 192
	funcstats:  IF expr THEN.funcstats ELSE funcstats FI 

	BEGIN  shift 164
	SKIP  shift 155
	READ  shift 158
	FREE  shift 159
	RETURN  shift 153
	EXIT  shift 160
	PRINT  shift 161
	PRINTLN  shift 162
	IF  shift 152
	WHILE  shift 163
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	prefuncstat  goto 154
	prefuncstats  goto 151
	funcstats  goto 201
	assignlhs  goto 157
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 156
	arraytype  goto 20
	pairtype  goto 21

state 193
	prefuncstat:  typeDef IDENTIFIER ASSIGNMENT.assignrhs 

	NEWPAIR  shift 76
	CALL  shift 77
	FST  shift 30
	SND  shift 31
	NOT  shift 47
	LEN  shift 48
	ORD  shift 49
	CHR  shift 50
	PLUS  shift 52
	SUB  shift 51
	TRUE  shift 40
	FALSE  shift 41
	NULL  shift 54
	OPENSQUARE  shift 78
	OPENROUND  shift 53
	STRINGCONST  shift 43
	IDENTIFIER  shift 45
	INTEGER  shift 39
	CHARACTER  shift 42
	.  error

	assignrhs  goto 202
	expr  goto 73
	pairelem  goto 75
	arrayliter  goto 74
	arrayelem  goto 46
	pairliter  goto 44

state 194
	prefuncstat:  assignlhs ASSIGNMENT assignrhs.    (12)

	.  reduce 12 (src line 102)


state 195
	prefuncstat:  WHILE expr DO.statements DONE 

	BEGIN  shift 18
	SKIP  shift 8
	READ  shift 10
	FREE  shift 11
	RETURN  shift 12
	EXIT  shift 13
	PRINT  shift 14
	PRINTLN  shift 15
	IF  shift 16
	WHILE  shift 17
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	statement  goto 6
	statements  goto 203
	assignlhs  goto 9
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 62
	arraytype  goto 20
	pairtype  goto 21

state 196
	prefuncstat:  BEGIN statements END.    (19)

	.  reduce 19 (src line 109)


state 197
	function:  typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND IS funcstats END.    (5)

	.  reduce 5 (src line 90)


state 198
	assignrhs:  NEWPAIR OPENROUND expr COMMA expr CLOSEROUND.    (31)

	.  reduce 31 (src line 126)


state 199
	funcstats:  prefuncstats SEMICOLON RETURN expr.    (6)
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  reduce 6 (src line 95)


state 200
	funcstats:  prefuncstats SEMICOLON IF expr.THEN funcstats ELSE funcstats FI 
	expr:  expr.PLUS expr 
	expr:  expr.SUB expr 
	expr:  expr.MUL expr 
	expr:  expr.MOD expr 
	expr:  expr.DIV expr 
	expr:  expr.LT expr 
	expr:  expr.GT expr 
	expr:  expr.LTE expr 
	expr:  expr.GTE expr 
	expr:  expr.EQ expr 
	expr:  expr.NEQ expr 
	expr:  expr.AND expr 
	expr:  expr.OR expr 

	THEN  shift 204
	MUL  shift 81
	DIV  shift 83
	MOD  shift 82
	PLUS  shift 79
	SUB  shift 80
	AND  shift 90
	OR  shift 91
	GT  shift 85
	GTE  shift 87
	LT  shift 84
	LTE  shift 86
	EQ  shift 88
	NEQ  shift 89
	.  error


state 201
	funcstats:  IF expr THEN funcstats.ELSE funcstats FI 

	ELSE  shift 205
	.  error


state 202
	prefuncstat:  typeDef IDENTIFIER ASSIGNMENT assignrhs.    (11)

	.  reduce 11 (src line 101)


state 203
	prefuncstat:  WHILE expr DO statements.DONE 
	statements:  statements.SEMICOLON statement 

	DONE  shift 206
	SEMICOLON  shift 33
	.  error


state 204
	funcstats:  prefuncstats SEMICOLON IF expr THEN.funcstats ELSE funcstats FI 

	BEGIN  shift 164
	SKIP  shift 155
	READ  shift 158
	FREE  shift 159
	RETURN  shift 153
	EXIT  shift 160
	PRINT  shift 161
	PRINTLN  shift 162
	IF  shift 152
	WHILE  shift 163
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	prefuncstat  goto 154
	prefuncstats  goto 151
	funcstats  goto 207
	assignlhs  goto 157
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 156
	arraytype  goto 20
	pairtype  goto 21

state 205
	funcstats:  IF expr THEN funcstats ELSE.funcstats FI 

	BEGIN  shift 164
	SKIP  shift 155
	READ  shift 158
	FREE  shift 159
	RETURN  shift 153
	EXIT  shift 160
	PRINT  shift 161
	PRINTLN  shift 162
	IF  shift 152
	WHILE  shift 163
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	prefuncstat  goto 154
	prefuncstats  goto 151
	funcstats  goto 208
	assignlhs  goto 157
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 156
	arraytype  goto 20
	pairtype  goto 21

state 206
	prefuncstat:  WHILE expr DO statements DONE.    (18)

	.  reduce 18 (src line 108)


state 207
	funcstats:  prefuncstats SEMICOLON IF expr THEN funcstats.ELSE funcstats FI 

	ELSE  shift 209
	.  error


state 208
	funcstats:  IF expr THEN funcstats ELSE funcstats.FI 

	FI  shift 210
	.  error


state 209
	funcstats:  prefuncstats SEMICOLON IF expr THEN funcstats ELSE.funcstats FI 

	BEGIN  shift 164
	SKIP  shift 155
	READ  shift 158
	FREE  shift 159
	RETURN  shift 153
	EXIT  shift 160
	PRINT  shift 161
	PRINTLN  shift 162
	IF  shift 152
	WHILE  shift 163
	FST  shift 30
	SND  shift 31
	INT  shift 25
	BOOL  shift 26
	CHAR  shift 27
	STRING  shift 28
	PAIR  shift 29
	IDENTIFIER  shift 22
	.  error

	prefuncstat  goto 154
	prefuncstats  goto 151
	funcstats  goto 211
	assignlhs  goto 157
	pairelem  goto 24
	arrayelem  goto 23
	basetype  goto 19
	typeDef  goto 156
	arraytype  goto 20
	pairtype  goto 21

state 210
	funcstats:  IF expr THEN funcstats ELSE funcstats FI.    (8)

	.  reduce 8 (src line 97)


state 211
	funcstats:  prefuncstats SEMICOLON IF expr THEN funcstats ELSE funcstats.FI 

	FI  shift 212
	.  error


state 212
	funcstats:  prefuncstats SEMICOLON IF expr THEN funcstats ELSE funcstats FI.    (7)

	.  reduce 7 (src line 96)


64 terminals, 25 nonterminals
97 grammar rules, 213/2000 states
0 shift/reduce, 0 reduce/reduce conflicts reported
74 working sets used
memory: parser 434/30000
154 extra closures
1407 shift entries, 3 exceptions
134 goto entries
198 entries saved by goto default
Optimizer space used: output 614/30000
614 table entries, 167 zero
maximum spread: 64, maximum offset: 209
