package armArchitecture

// Operand of ARM instruction set
type Operand int

// Operands in alphabetical order

// CHECK WHICH OPERANDS WE DON@T USE AND REMOVE THEM
const (
	ADDS Operand = iota
	B
	BEQ
	BL
	BLCS
	BLEQ
	BLLT
	BLNE
	BLVS
	BNE
	CMP
	DIV
	EOR
	LDR
	LDRCS
	LDREQ
	LDRLT
	LDRNE
	LDRSB
	MOV
	MOVEQ
	MOVGE
	MOVGT
	MOVLE
	MOVLT
	MOVNE
	MUL
	POP
	POPEQ
	PUSH
	RSBS
	SMULL
	STR
	STRB
	SUB
	SUBS
)
