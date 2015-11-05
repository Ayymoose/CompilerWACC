package lexer

import "unicode"

func lexBegin(l *lexer) stateFn {
	l.width = len(begin)
	l.pos += l.width
	if unicode.IsLetter(l.next()) {
		l.backup()
		return lexIdentifier
	}
	l.emit(itemBegin)
	return lexInsideProgram
}

func lexEnd(l *lexer) stateFn {
	l.width = len(end)
	l.pos += l.width
	if unicode.IsLetter(l.next()) {
		l.backup()
		return lexIdentifier
	}
	l.emit(itemEnd)
	return nil
}

func lexIdentifier(l *lexer) stateFn {
	if !unicode.IsLetter(l.next()) {
		return l.errorf("bad identifier syntax: %q", l.input[l.start:l.pos])
	}
	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) {
		l.next()
	}
	l.backup()
	l.emit(itemIdentifier)
	return lexInsideProgram
}

func lexChar(l *lexer) stateFn {
}

func lexString(l *lexer) stateFn {

}

func lexNumber(l *lexer) stateFn {
	l.accept("+-")
	digits := "0123456789"
	l.acceptRun(digits)
	if unicode.IsLetter(l.peek()) {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	l.emit(itemInt)
	return lexInsideProgram
}
