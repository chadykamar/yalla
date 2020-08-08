package common

import (
	"fmt"
	"io"
	"log"
	"unicode"

	"github.com/chadykamar/bufrr"
)

// Scanner parses Yalla source code
type Scanner struct {
	line    int
	start   int
	current int
	text    []rune
	*bufrr.Reader
}

// NewScanner initializes a scanner for the given reader
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{1, 0, 0, []rune{}, bufrr.NewReader(reader)}

}

func (sc *Scanner) isAtEnd() bool {
	r, _, err := sc.PeekRune()
	if err != nil {
		log.Fatal(err)
	}
	return r == bufrr.EOF

}

func (sc *Scanner) skipWhitespace() {
	if sc.isAtEnd() {
		return
	}
	for {
		r := sc.peek()

		switch r {
		case ' ', '\r', '\t':
			{
				sc.advance()
				continue
			}
		case '\n':
			{
				sc.line++
				sc.advance()
				continue
			}
		case '/':
			{
				// Check if it is followed by another '/'
				if sc.peekNext() == '/' {

					// If so advance until newline or end of file
					for ; sc.peek() != '\n' && !sc.isAtEnd(); sc.advance() {
						sc.advance()
					}
				} else {
					return
				}

			}
		default:
			return
		}
	}
}

func (sc *Scanner) peek() rune {
	if sc.isAtEnd() {
		return bufrr.EOF
	}

	r, _, err := sc.PeekRune()
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func (sc *Scanner) peekNext() rune {
	if sc.isAtEnd() {
		return bufrr.EOF
	}

	_, _, err := sc.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}

	r, _, err := sc.PeekRune()
	if err != nil {
		log.Fatalln(err)
	}

	err = sc.UnreadRune()
	if err != nil {
		log.Fatalln(err)
	}

	return r

}

func (sc *Scanner) advance() rune {
	if sc.isAtEnd() {
		return bufrr.EOF
	}

	r, _, err := sc.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}
	sc.current++
	sc.text = append(sc.text, r)
	return r
}

func (sc *Scanner) scanToken() Token {
	sc.skipWhitespace()
	sc.start = sc.current
	sc.text = []rune{}

	r := sc.advance()

	if r == bufrr.EOF {
		return sc.NewToken(TokenEOF)
	}

	if unicode.IsDigit(r) {
		return sc.numberToken()
	}
	if unicode.IsLetter(r) {
		return sc.identifierToken()
	}

	switch r {

	// Single rune tokens
	case '(':
		return sc.NewToken(TokenLeftParen)
	case ')':
		return sc.NewToken(TokenRightParen)
	case '{':
		return sc.NewToken(TokenLeftBrace)
	case '}':
		return sc.NewToken(TokenRightBrace)
	case '[':
		return sc.NewToken(TokenLeftBracket)
	case ']':
		return sc.NewToken(TokenRightBracket)
	case ';':
		return sc.NewToken(TokenSemicolon)
	case ':':
		return sc.NewToken(TokenColon)
	case ',':
		return sc.NewToken(TokenComma)
	case '.':
		return sc.NewToken(TokenDot)
	case '-':
		return sc.NewToken(TokenMinus)
	case '+':
		return sc.NewToken(TokenPlus)
	case '/':
		return sc.NewToken(TokenSlash)
	case '*':
		return sc.NewToken(TokenAsterisk)

	// Two-rune tokens
	case '!':
		{
			if sc.match('=') {
				return sc.NewToken(TokenBangEqual)
			}
			return sc.NewToken(TokenBang)

		}
	case '=':
		{
			if sc.match('=') {
				return sc.NewToken(TokenEqualEqual)
			}
			return sc.NewToken(TokenEqual)

		}
	case '<':
		{
			if sc.match('=') {
				return sc.NewToken(TokenLessEqual)
			}
			return sc.NewToken(TokenLess)

		}
	case '>':
		{
			if sc.match('=') {
				return sc.NewToken(TokenGreaterEqual)
			}
			return sc.NewToken(TokenGreater)

		}

	// Literal tokens
	case '"':
		return sc.stringToken()

	}

	msg := fmt.Sprintf("Unexpected character %U", r)
	return sc.ErrorToken(msg)
}

func (sc *Scanner) numberToken() Token {

	for unicode.IsDigit(sc.peek()) {
		sc.advance()
	}

	if sc.peek() == '.' && unicode.IsDigit(sc.peekNext()) {
		sc.advance()

		for unicode.IsDigit(sc.peek()) {
			sc.advance()
		}
	}
	return sc.NewToken(TokenNumber)

}

func (sc *Scanner) identifierToken() Token {

	for unicode.IsDigit(sc.peek()) || unicode.IsLetter(sc.peek()) {
		sc.advance()
	}

	return sc.NewToken(sc.identifierType())

}

func (sc *Scanner) identifierType() TokenType {

	switch string(sc.text) {
	case "and":
		return TokenAnd
	case "else":
		return TokenElse
	case "if":
		return TokenIf
	case "nil":
		return TokenNil
	case "or":
		return TokenOr
	case "return":
		return TokenReturn
	case "false":
		return TokenFalse
	case "for":
		return TokenFor
	case "func":
		return TokenFunc
	case "true":
		return TokenTrue
	}

	return TokenIdentifier
}

// func (sc *Scanner) checkKeyword(tokenType TokenType) TokenType {

// 	return TokenIdentifier
// }

func (sc *Scanner) stringToken() Token {
	for sc.peek() != '"' && !sc.isAtEnd() {
		r := sc.advance()
		if r == '\n' {
			sc.line++
		}

	}

	if sc.isAtEnd() {
		return sc.ErrorToken("Unterminated string")
	}

	// Consume closing quote
	sc.advance()
	return sc.NewToken(TokenString)

}

// resolveTwoRuneToken a token with the correct type by peeking one rune ahead
func (sc *Scanner) match(expectedRune rune) bool {
	if sc.isAtEnd() {
		return false
	}
	r := sc.peek()

	if r == expectedRune {
		_ = sc.advance()
		return true
	}

	return false

}

// NewToken instantiates the Token that the scanner is currently scanning
func (sc *Scanner) NewToken(tokenType TokenType) Token {
	var token Token
	token.tokenType = tokenType
	token.str = string(sc.text)

	return token
}

// ErrorToken instantiates error tokens
func (sc *Scanner) ErrorToken(message string) Token {
	var token Token
	token.tokenType = TokenError
	token.str = message
	token.line = sc.line

	return token
}

// TokenType identifies the different types of tokens in the Yalla language
type TokenType int

const (
	// TokenLeftParen identifies the left parenthesis token '('
	TokenLeftParen TokenType = iota
	// TokenRightParen identifies the right parenthesis token ')'
	TokenRightParen
	// TokenLeftBrace identifies the left brace token '{'
	TokenLeftBrace
	// TokenRightBrace identifies the right brace token '}'
	TokenRightBrace
	// TokenLeftBracket identifies the left bracket token '['
	TokenLeftBracket
	// TokenRightBracket identifies the right bracket token ']'
	TokenRightBracket
	// TokenComma identifies the comma token ','
	TokenComma
	// TokenDot identifies the dot or period token '.'
	TokenDot
	// TokenMinus identifies the minus token '-'
	TokenMinus
	// TokenPlus identifies the plus token '+'
	TokenPlus
	// TokenSemicolon identifies the semicolon token ';'
	TokenSemicolon
	// TokenColon identifies the colon token ':'
	TokenColon
	// TokenSlash identifies the slash token '/'
	TokenSlash
	// TokenAsterisk identifies the asterisk token '*'
	TokenAsterisk

	// TokenBang identifies the bang token '!'
	TokenBang
	// TokenBangEqual identifies the bang equal token '!='
	TokenBangEqual
	// TokenEqual identifies the equal token '='
	TokenEqual
	// TokenEqualEqual identifies the equal equal token '=='
	TokenEqualEqual
	// TokenGreater is the comparison greater than operator token '>'
	TokenGreater
	// TokenGreaterEqual is the comparison greater than or equal operator token '>='
	TokenGreaterEqual
	// TokenLess is the comparison less than operator token '<'
	TokenLess
	// TokenLessEqual is the comparison less than or equal operator token '<='
	TokenLessEqual

	// TokenIdentifier denotes identifier tokens such as function, class or variable names
	TokenIdentifier
	// TokenString identifies string literals
	TokenString
	// TokenNumber identifies number literals
	TokenNumber

	// TokenAnd is the boolean 'and' keyword token
	TokenAnd
	// TokenElse is the token for the keyword 'else'
	TokenElse
	// TokenFalse is the token for the keyword 'false'
	TokenFalse
	// TokenFor is the token for keyword 'for'
	TokenFor
	// TokenFunc is the token for the keyword 'fun'
	TokenFunc
	// TokenIf is the token for the keyword 'if'
	TokenIf
	// TokenNil is the token for the keyword 'nil'
	TokenNil
	// TokenOr is the token for the keyword 'or'
	TokenOr
	// TokenReturn is the token for the keyword 'return'
	TokenReturn
	// TokenTrue is the token for the keyword 'true'
	TokenTrue

	// TokenError is the token type returned when no token can be parsed
	TokenError
	// TokenEOF is simply the end of file token returned when the scanner reaches the end of the file
	TokenEOF
)

// Token contains the information about a single token as returned by the scanner
type Token struct {
	tokenType TokenType
	str       string
	line      int
}
