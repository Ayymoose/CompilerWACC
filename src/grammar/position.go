package grammar

import "fmt"

type Pos uint

// Declare and assign
var NoPos Pos = 0

// isValid returns true if its is a valid position
func (p Pos) isValid() bool {
	return p != NoPos
}

type Position struct {
	FileName string
	Row, Col int
}

func (p Position) String() string {
	if p.FileName == "" {
		return fmt.Sprintf("%d:%d", p.Row, p.Col)
	}
	return fmt.Sprintf("%s:%d:%d", p.FileName, p.Row, p.Col)
}
