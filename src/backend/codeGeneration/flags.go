package codeGeneration

// Flag indicate the state reached
// while traversing the AST nodes
type Flag int

// Flags emun
const (
	AND Flag = iota
	ARRAYBOUNDS
	DIVIDEDBYZERO
	FREE
	OR
	OVERFLOW
	READINT
	READCHAR
	RUNTIME
	PRINT
	PRINTLN
	PRINTNULLPOINTER
	PRINTSTRING
	PRINTINT
	PRINTBOOL
	PRINTREFERENCE
)

// FlagMap Flags are initially set to false
var FlagMap = map[Flag]bool{
	AND:              false,
	ARRAYBOUNDS:      false,
	DIVIDEDBYZERO:    false,
	FREE:             false,
	OR:               false,
	OVERFLOW:         false,
	READINT:          false,
	READCHAR:         false,
	RUNTIME:          false,
	PRINT:            false,
	PRINTLN:          false,
	PRINTNULLPOINTER: false,
	PRINTSTRING:      false,
	PRINTINT:         false,
	PRINTBOOL:        false,
	PRINTREFERENCE:   false,
}

// setFlag sets the passed flag to true in FlagMap
func (f Flag) setFlag(flag Flag) {
	FlagMap[flag] = true
}

// clearFlag sets the passed flag to false in FlagMap
func (f Flag) clearFlag(flag Flag) {
	FlagMap[flag] = false
}

// getFlag returns the state of the passed flag
func (f Flag) getFlag(flag Flag) bool {
	return FlagMap[flag]
}
