package armArchitecture

// Directive of ARM instuction set
type Directive struct {
	directive string
}

// Returns the string representation of a Directive
func (d Directive) String() string {
	return "\t." + d.directive
}
