package armArchitecture

import "fmt"

// Address of ARM instruction set:
// * [value]
// * [value, #offset]
type Address struct {
	value  interface{}
	offset int
}

// Returns string representation of Address
func (a Address) String() string {
	if a.offset == 0 {
		return fmt.Sprint("[", a.value, "]")
	}
	return fmt.Sprint("[", a.value, ", ", a.offset, "]")
}
