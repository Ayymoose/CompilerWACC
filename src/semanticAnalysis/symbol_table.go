// Syboml Table for compiler
// Stores information about variable names, function names, etc

// Carries out the following functions
// Stores names of all entities in one place
// Verifies if a variable has ben declared
// Implements type checking by verifying assignments and expressions in source code are semantically correct
// Scope resolution

// Hash table data structure
// Maintains entries in the following format
// <symbolName, type>
// EG: int foo  -->  <foo, int>

// Each symbol can have different attributes
// EG function have a number of variable with types
//    but an int just has variable name and type

// Provide the following operations:

// Posibility of adding a scope attribute???

// insert()
// Insert an symbol into the table
// int a  -->  insert(a, int)

// lookUp()
// Searches a name in the symbol table to determine :
// If the symbol exists in the table
// If it is declared before it is used
// if the name is used in the scope
// if the symbol is initialised
// if the symbol is declared multiple times

// a --> lookUp(a)

// This method returns 0 if the symbol does not exist in the symbol table
// If the symbol exists in the symbol table it returns its attributes stored in the table

// Set up a separate symbol table for each scope

// Algorithm for whenever a name needs to be searched in the symbol table
// First a symbol will be searched in the current scope  --> i.e. the current symbol table
// If a name is found then the search is complete
// else it will be searched in the parent symcol table until
// Either the name is found or global symbol table has been searched for name

// maps are not safe for concurrent read/write

package semanticAnalysis

import "grammar"

// SymbolTable constructor
type SymbolTable struct {
	parent      *SymbolTable
	semanticMap map[string]grammar.Type // A map of strings to a list of tokens
}

// Creates a new instance of a symbolTable with a parent to its pointer
func (symbolTable *SymbolTable) New() *SymbolTable {
	newSymbolTable := &SymbolTable{}
	newSymbolTable.parent = symbolTable
	newSymbolTable.semanticMap = make(map[string][]grammar.Token)
	return newSymbolTable
}

func (SymbolTable *SymbolTable) insert(key string, value []grammar.Token) {
	// convert value string into its grammar token here
	SymbolTable[key] = value
}

//Checks if ident is already declared
func (SymbolTable *SymbolTable) isDefined(key string) bool {
	curr := SymbolTable
	for curr != nil {
		if curr.contains(key) {
			return true
		}
		curr := SymbolTable.parent
	}
	return false
}

func (SymbolTable *SymbolTable) contains(key string) bool {
	_, ok := SymbolTable.semanticMap[key]
	return ok
}

func (SymbolTable *SymbolTable) getTypeOfIdent(key string) grammar.Token {
	curr := SymbolTable
	for curr != nil {
		if curr.contains(key) {
			k := curr.semanticMap[key]
			return k
		}
		curr := SymbolTable.parent
	}
	return nil
}

/*
func (s *SymbolTable) Insert(token string, list reflect.TypeOf(list).Elem()) {

}
*/
/*
func addNode(parent *Node, val int) *Node {
  return &Node{nil, nil, parent, val}
}

func (b *BinaryTree) Insert(val int) (n *Node) {
  if b.root == nil {
    n = addNode(nil, val)
    b.root = n
  } else {
    n = b.insert(b.root, nil, val)
  }
  return
}

func (b *BinaryTree) insert(root *Node, parent *Node, val int) *Node {
  switch {
  case root == nil:
    return addNode(parent, val)
  case val <= root.value:
    root.left = b.insert(root.left, root, val)
  case val > root.value:
    root.right = b.insert(root.right, root, val)
  }
  return root
}



func main() {
  newSt = new(st)



  st := new(nil)
  st.insert(whate)

  st.Insert(whatever, )
  b.Insert(1)
  b.Insert(-1)
  fmt.Println(b.root)
  b.Delete(1)
  fmt.Println(b.root)
}
*/
