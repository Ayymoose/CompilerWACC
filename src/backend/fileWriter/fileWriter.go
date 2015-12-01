package fileWriter

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToFile(input string) {
	f, err := os.Create("/fileWriter/original.s")
	check(err)
	defer f.Close()
	n3, err := f.WriteString(input + "\n")
	f.Sync()
}
