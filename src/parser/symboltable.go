package parser

// SymbolTable constructor
type SymbolTable struct {
	Parent *SymbolTable
	Table  map[string]Type
}

// New Constructor creates new instance of a symbolTable with pointer to its parent
func (symbolTable *SymbolTable) New() *SymbolTable {
	return &SymbolTable{Parent: symbolTable, Table: make(map[string]Type)}
}

// Inserts a given key and value into the symbol table
func (symbolTable *SymbolTable) insert(key string, value Type) {
	symbolTable.Table[key] = value
}

// Checks if the key is already declared in the symbol table
func (symbolTable *SymbolTable) isDefined(key string) bool {
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
func (symbolTable *SymbolTable) contains(key string) bool {
	_, ok := symbolTable.Table[key]
	return ok
}

// Given an Ident key, returns the slice of Type tokens
func (symbolTable *SymbolTable) getTypeOfIdent(key string) Type {
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
