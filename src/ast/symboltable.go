package ast

// SymbolTable constructor
type SymbolTable struct {
	Parent     *SymbolTable
	Table      map[Ident]Type
	OffsetVals map[string]int
}

// New Constructor creates new instance of a symbolTable with pointer to its parent
func (symbolTable *SymbolTable) New() *SymbolTable {
	return &SymbolTable{Parent: symbolTable, Table: make(map[Ident]Type), OffsetVals: make(map[string]int)}
}

// Inserts a given key and value into the symbol table
func (symbolTable *SymbolTable) insert(key Ident, value Type) {
	symbolTable.Table[key] = value
}

// Inserts the offset of a given key into symbol table
func (symbolTable *SymbolTable) insertOffset(key string, offset int) {
	symbolTable.OffsetVals[key] = offset
}

// Returns the offset of a given key
func (symbolTable *SymbolTable) getOffset(key string) int {
	return symbolTable.OffsetVals[key]
}

// Checks if the key is already declared in the symbol table
func (symbolTable *SymbolTable) isDefined(key Ident) bool {
	curr := symbolTable
	for curr != nil {
		if curr.contains(key) {
			return true
		}
		curr = symbolTable.Parent
	}
	return false
}

// Helper function which return a boolean depending on
// whether or not the given key is in the symbol table
func (symbolTable *SymbolTable) contains(key Ident) bool {
	_, ok := symbolTable.Table[key]
	return ok
}

// Given an Ident key, returns the slice of Type tokens
func (symbolTable *SymbolTable) getTypeOfIdent(key Ident) Type {
	curr := symbolTable
	for curr != nil {
		if curr.contains(key) {
			k := curr.Table[key]
			return k
		}
		curr = symbolTable.Parent
	}
	return -1
}
