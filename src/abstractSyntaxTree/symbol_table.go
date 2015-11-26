package parser

import "grammar"

// SymbolTable constructor
type SymbolTable struct {
	parent      *SymbolTable
	semanticMap map[string][]grammar.Type
}

// Constructor creates new instance of a symbolTable with pointer to its parent
func (symbolTable *SymbolTable) New() *SymbolTable {
	newSymbolTable := &SymbolTable{}
	newSymbolTable.parent = symbolTable
	newSymbolTable.semanticMap = make(map[string][]grammar.Type)
	return newSymbolTable
}

// Inserts a given key and value into the symbol table
func (SymbolTable *SymbolTable) insert(key string, value []grammar.Type) {
	SymbolTable.semanticMap[key] = value
}

// Checks if the key is already declared in the symbol table
func (SymbolTable *SymbolTable) isDefined(key string) bool {
	curr := SymbolTable
	for curr != nil {
		if curr.contains(key) {
			return true
		}
		curr = SymbolTable.parent
	}
	return false
}

// Helper function which return a boolean depending on
// whether or not the given key is in the symbol table
func (SymbolTable *SymbolTable) contains(key string) bool {
	_, ok := SymbolTable.semanticMap[key]
	return ok
}

// Given an Ident key, returns the slice of Type tokens
func (SymbolTable *SymbolTable) getTypeOfIdent(key string) []grammar.Type {
	curr := SymbolTable
	for curr != nil {
		if curr.contains(key) {
			k := curr.semanticMap[key]
			return k
		}
		curr = SymbolTable.parent
	}
	return nil
}
