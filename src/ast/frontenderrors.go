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

func errorExit(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "exit value must be type int:", getLine(file, pos))
  return errors.New(msg)
}

func errorPair(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Cannot free non pair type:", getLine(file, pos))
  return errors.New(msg)
}

func errorRead(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Cannot read non Char or Int type:", getLine(file, pos))
  return errors.New(msg)
}

func errorArray(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "LHS is not of type Array:", getLine(file, pos))
  return errors.New(msg)
}

func errorTypeMatch(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Types do not match:", getLine(file, pos))
  return errors.New(msg)
}

func errorDeclared(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Variable already declared:", getLine(file, pos))
  return errors.New(msg)
}

func errorReturn(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Return type does not match:", getLine(file, pos))
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
  end := pos + strings.Index(str[pos:], "\n")
  result :=  str[start+1:end]
  return result
}
