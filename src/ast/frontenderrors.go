package ast

import (
  "errors"
  "fmt"
  "strings"
)

func errorConditional(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "invalid conditional:", getLine(file, pos))
  return errors.New(msg)
}

//Returns line number and column number of current lexing item in .wacc file
func LineAndCol(file *string, pos int) (line int, col int) {
  str := *file
	line = 1 + strings.Count(str[:pos], "\n")
	col = pos - strings.LastIndex(str[:pos], "\n")
	return
}

//Returns line number and column number of current lexing item in .wacc file
func getLine(file *string, pos int) string {
  str := *file
  start := strings.LastIndex(str[:pos], "\n")
  end := strings.Index(str[pos:], "\n")
  return str[start:end-1]
}
