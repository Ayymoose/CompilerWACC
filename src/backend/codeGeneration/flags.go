package codeGeneration

// Flag indicate the state reached
// while traversing the AST nodes
type Flag int

// Flags emun
const (
	flag_AND Flag = iota
	flag_ARRAYBOUNDS
	flag_DIVIDEDBYZERO
	flag_FREE
	flag_OR
	flag_OVERFLOW
	flag_READINT
	flag_READCHAR
	flag_RUNTIME
	flag_PRINT
	flag_PRINTLN
	flag_PRINTNULLPOINTER
	flag_PRINTSTRING
	flag_PRINTINT
	flag_PRINTBOOL
	flag_PRINTREFERENCE
)

// FlagMap Flags are initially set to false
var FlagMap = map[Flag]bool{
	flag_AND:              false,
	flag_ARRAYBOUNDS:      false,
	flag_DIVIDEDBYZERO:    false,
	flag_FREE:             false,
	flag_OR:               false,
	flag_OVERFLOW:         false,
	flag_READINT:          false,
	flag_READCHAR:         false,
	flag_RUNTIME:          false,
	flag_PRINT:            false,
	flag_PRINTLN:          false,
	flag_PRINTNULLPOINTER: false,
	flag_PRINTSTRING:      false,
	flag_PRINTINT:         false,
	flag_PRINTBOOL:        false,
	flag_PRINTREFERENCE:   false,
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
