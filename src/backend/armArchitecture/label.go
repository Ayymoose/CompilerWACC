package armArchitecture

// Label of ARM instruction set
type Label struct {
	label string
}

// Returns the string representation of Label
func (l Label) String() string {
	return l.label
}
