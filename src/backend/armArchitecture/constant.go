package armArchitecture

// Constant of ARM instruction set
type Constant struct {
	value string // value can be string or int. If int then convert to string in caller
}

// Returns string representation of Constant
func (c Constant) String() string {
	return "=" + c.value
}
