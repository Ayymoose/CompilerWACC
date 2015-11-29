package armArchitecture

import "fmt"

// Branch of ARM instruction set
type Branch struct {
	branch int
}

// Returns string representation of Branch
func (b Branch) String() string {
	return fmt.Sprint("L", b.branch, ":")
}
