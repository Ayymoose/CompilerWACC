package armArchitecture

// Immediate value of ARM instruction set
type Immediate struct {
	value string // can be int or string, if int convert to string in caller
}

// Returns the string representation of Immediate value
func (i Immediate) String() string {
	return "#" + i.value
}
