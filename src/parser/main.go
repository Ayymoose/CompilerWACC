package main

const SYNTAX_ERROR = 100
const SEMANTIC_ERROR = 200

func func main() {
  file := os.Args[1] // index 1 is file path
  bytearray, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  s := string(bytearray)
program, errs := Parse(file, s)

}
