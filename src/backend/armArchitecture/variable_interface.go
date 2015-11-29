package armArchitecture

// Variable interface is a superclass of the following:
// * register
// * constant
// * address
// * empty_variable
type Variable interface {
	String() string
}
