package lexer

type TokenType string

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	NEWLINE TokenType = "NEWLINE"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"  // add, foobar, x, y, ...
	INT    TokenType = "INT"    // 1343456
	REAL   TokenType = "REAL"   // 3.14
	STRING TokenType = "STRING" // "foo"
	BOOL   TokenType = "BOOL"   // true, false
	NIL    TokenType = "NIL"    // nil

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	PERCENT  TokenType = "%"

	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
	LT     TokenType = "<"
	GT     TokenType = ">"
	LTE    TokenType = "<="
	GTE    TokenType = ">="

	AND TokenType = "and"
	OR  TokenType = "or"
	NOT TokenType = "not"

	// Delimiters
	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	COMMA    TokenType = ","
	COLON    TokenType = ":"
	ARROW    TokenType = "->"
	PIPE     TokenType = "|"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	DIM      TokenType = "DIM"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELIF     TokenType = "ELIF"
	ELSE     TokenType = "ELSE"
	THEN     TokenType = "THEN"
	END      TokenType = "END"
	RETURN   TokenType = "RETURN"
	LOOP     TokenType = "LOOP"
	REPEAT   TokenType = "REPEAT"
	BREAK    TokenType = "BREAK"
	CONTINUE TokenType = "CONTINUE"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

var keywords = map[string]TokenType{
	"func":     FUNCTION,
	"dim":      DIM,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"elif":     ELIF,
	"else":     ELSE,
	"then":     THEN,
	"end":      END,
	"return":   RETURN,
	"loop":     LOOP,
	"repeat":   REPEAT,
	"break":    BREAK,
	"continue": CONTINUE,
	"and":      AND,
	"or":       OR,
	"not":      NOT,
	"nil":      NIL,
}

var matchingTokens = map[string]TokenType{
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  MINUS,
	"*":  ASTERISK,
	"/":  SLASH,
	"%":  PERCENT,
	"<":  LT,
	">":  GT,
	"(":  LPAREN,
	")":  RPAREN,
	"[":  LBRACKET,
	"]":  RBRACKET,
	"{":  LBRACE,
	"}":  RBRACE,
	",":  COMMA,
	":":  COLON,
	"|":  PIPE,
	"\n": NEWLINE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
