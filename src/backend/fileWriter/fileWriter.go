package fileWriter

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// File struct which contains an array of ARM instuction set strings
type File struct {
	code []string
}

// Appends a string to the code slice in the File struct
func (f File) appendStringToFile(input string) {
	f.code = append(f.code, input)
}

// Writes all the strings contained in the File struct to the file
func (f File) writeToFile(input string) {
	file, err := os.Create("/fileWriter/original.s")
	check(err)
	defer file.Close()
	for _, str := range f.code {
		n3, err := file.WriteString(str)
	}
	file.Sync()
}
