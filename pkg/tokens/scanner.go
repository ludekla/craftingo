package tokens

import (
	"fmt"
	"log"
	"strconv"
)

// helpers
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch == '_'
}

func isAlphaNum(ch byte) bool {
	return isDigit(ch) || isAlpha(ch)
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(src string) Scanner {
	return Scanner{source: src, line: 1}
}

func (s *Scanner) addToken(tokenType int) {
	text := s.source[s.start:s.current]
	fmt.Println(s.start, s.current, text)
	s.tokens = append(s.tokens, NewToken(tokenType, text, 0.0, s.line))
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) ScanTokens() {
	for !s.EOF() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, NewToken(EOF, "", 0.0, s.line))
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
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
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.EOF() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\t', '\r':
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			log.Fatalf("unexpected character in line %d: %v", s.line, c)
		}
	}
}

func (s *Scanner) EOF() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) Tokens() []Token {
	return s.tokens
}

func (s *Scanner) match(expected byte) bool {
	switch {
	case s.EOF():
		return false
	case s.source[s.current] != expected:
		return false
	default:
		s.current++
		return true
	}
}

func (s *Scanner) peek() byte {
	if s.EOF() {
		return '\000'
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 == len(s.source) {
		return '\000'
	}
	return s.source[s.current+1]
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// first get over the dot
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	value, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.tokens = append(s.tokens, NewToken(NUMBER, "", value, s.line))
}

func (s *Scanner) identifier() {
	for isAlphaNum(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	if tokenTyp, ok := keywords[text]; ok {
		s.addToken(tokenTyp)
	} else {
		s.addToken(IDENTIFIER)
	}
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.EOF() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.EOF() {
		log.Fatalf("unterminated string in line %d", s.line)
	}
	s.advance()
	text := s.source[s.start+1 : s.current-1]
	s.tokens = append(s.tokens, NewToken(STRING, text, 0.0, s.line))
}
