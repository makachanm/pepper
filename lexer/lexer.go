package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int
	col          int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, col: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	if l.ch == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	startLine := l.line
	startCol := l.col

	if l.ch == '`' {
		l.readChar() // skip the backtick
		tok.Type = STRING
		tok.Literal = l.readString()
		tok.Line = startLine
		tok.Column = startCol
		l.readChar() // skip the closing backtick
		return tok
	}

	if tokType, ok := matchingTokens[string(l.ch)]; ok {
		literal := string(l.ch)
		peek := l.peekChar()
		switch tokType {
		case ASSIGN:
			if peek == '=' {
				l.readChar()
				literal = "=="
				tokType = EQ
			}
		case LT:
			if peek == '=' {
				l.readChar()
				literal = "<="
				tokType = LTE
			}
		case GT:
			if peek == '=' {
				l.readChar()
				literal = ">="
				tokType = GTE
			}
		}
		tok = Token{Type: tokType, Literal: literal, Line: startLine, Column: startCol}
	} else {
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			// return here because readIdentifier advances the cursor
			return Token{Type: LookupIdent(literal), Literal: literal, Line: startLine, Column: startCol}
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			tokType := INT
			if !isInt(literal) {
				tokType = REAL
			}
			// return here because readNumber advances the cursor
			return Token{Type: tokType, Literal: literal, Line: startLine, Column: startCol}
		} else if l.ch == 0 {
			tok = Token{Type: EOF, Literal: "", Line: startLine, Column: startCol}
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch), Line: startLine, Column: startCol}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for {
		if l.ch == ' ' || l.ch == '\t' || l.ch == '\n' {
			l.readChar()
		} else if l.ch == '/' && l.peekChar() == '*' {
			l.readChar() // consume "/"
			l.readChar() // consume "*",
			for {
				if l.ch == 0 { // EOF
					return
				}
				if l.ch == '*' && l.peekChar() == '/' {
					l.readChar() // consume "*",
					l.readChar() // consume "/"
					break
				}
				l.readChar()
			}
		} else {
			break
		}
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) || l.ch == '.' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position
	for {
		l.readChar()
		if l.ch == '`' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isInt(s string) bool {
	for _, c := range s {
		if c == '.' {
			return false
		}
	}
	return true
}
