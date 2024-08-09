package main

import "strconv"

var keywoards = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
	"break":  BREAK,
}

type Scanner struct {
	sourceArr []rune
	source    string
	tokens    []Token
	start     int
	current   int
	line      int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, sourceArr: []rune(source), line: 1}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: s.line})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	r := s.advance()
	switch r {
	// Symbols
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		s.addTokenConditional('=', BANG_EQUAL, BANG)
	case '=':
		s.addTokenConditional('=', EQUAL_EQUAL, EQUAL)
	case '<':
		s.addTokenConditional('=', LESS_EQUAL, LESS)
	case '>':
		s.addTokenConditional('=', GREATER_EQUAL, GREATER)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}

	// Ignore whitespace
	case ' ', '\r', '\t':
		break
	case '\n':
		s.line++

	// Literals
	case '"':
		s.string()
	default:
		if s.isDigit(r) {
			s.number()
		} else if s.isAlpha(r) {
			s.identifier()
		} else {
			err(s.line, "Unexpected character")
		}
	}
}

func (s *Scanner) advance() rune {
	s.current++
	return s.sourceArr[s.current-1]
}

func (s *Scanner) addToken(t TokenType) {
	s.addTokenLiteral(t, nil)
}

func (s *Scanner) addTokenConditional(expected rune, matchType, elseType TokenType) {
	if s.match(expected) {
		s.addToken(matchType)
	} else {
		s.addToken(elseType)
	}
}

func (s *Scanner) addTokenLiteral(t TokenType, literal interface{}) {
	text := s.sourceArr[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: t, Lexeme: string(text), Literal: literal, Line: s.line})
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.sourceArr[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	return s.sourceArr[s.current]
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		err(s.line, "Unterminated string")
		return
	}
	s.advance()
	text := s.sourceArr[s.start+1 : s.current-1]
	s.addTokenLiteral(STRING, string(text))
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	text := s.sourceArr[s.start:s.current]
	f, _ := strconv.ParseFloat(string(text), 64)
	s.addTokenLiteral(NUMBER, f)
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := string(s.sourceArr[s.start:s.current])
	t, ok := keywoards[text]
	if !ok {
		t = IDENTIFIER
	}
	s.addToken(t)
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\000'
	}
	return s.sourceArr[s.current+1]
}

func (s *Scanner) isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func (s *Scanner) isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func (s *Scanner) isAlphaNumeric(r rune) bool {
	return s.isAlpha(r) || s.isDigit(r)
}
