package parser

type Function struct {
	ident          string
	returnType     Type
	parameterTypes []Param
	statlist       interface{}
}

type Param struct {
	ident     string
	paramType Type
}
