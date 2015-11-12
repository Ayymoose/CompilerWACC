package grammar

import "fmt"

type Pos int

// Declare and assign
var NoPos Pos = -1

// isValid returns true if its is a valid position

//DO WE FUCKING NEED THIS SHIT NIGGA??

func (p Pos) isValid() bool {
	return p != NoPos
}

type Position struct {
	LineNum, ColNum int
}

func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.LineNum, p.ColNum)
}
