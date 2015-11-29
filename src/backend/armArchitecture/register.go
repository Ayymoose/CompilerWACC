package armArchitecture

import "fmt"

// Register of ARM instruction set
type Register struct {
	value  int
	offset int
	inUse  bool
}

// CheckInUse returns a boolean indicating if a given register is in use
func (r Register) CheckInUse() bool {
	return r.inUse
}

// SetInUse sets a current register in use
func (r Register) SetInUse() {
	r.inUse = true
}

func (r Register) String() string {
	offsetString := ""
	if r.offset != 0 {
		offsetString = ", #" + fmt.Sprint(r.offset)
	}

	switch r.value {
	case 13:
		return "sp" + offsetString
	case 14:
		return "lr" + offsetString
	case 15:
		return "pc" + offsetString
	default:
		return "r" + fmt.Sprint(r.value) + offsetString
	}
}
