package main

import (
	"unicode"
	"strings"
)

type Lexer struct {
	content string
}

func New(content string) *Lexer {
	return &Lexer {
		content : content,
	}
}

func (self *Lexer) TrimLeft() {
	for len(self.content) > 0 && unicode.IsSpace(rune(self.content[0])) {
		self.content = string([]byte(self.content)[1:])
	}
}

func (self *Lexer) chop(n int) string {
	var result = self.content[0:n]
	self.content = self.content[n:]
	return strings.ToUpper(result)
}

func (self *Lexer) chopWhile(condition func(rune)bool) string {
	var n = 0
	for n < len(self.content) && condition(rune(self.content[n])) {
		n += 1
	}
	return self.chop(n)
}

func (self *Lexer) Next() (string, bool) {
	self.TrimLeft()
	if len(self.content) == 0 {
		return "", false
	}

	if unicode.IsDigit(rune(self.content[0])) {
		result := self.chopWhile(unicode.IsDigit)
		return result, true
	}

	if unicode.IsLetter(rune(self.content[0])) {
		result := self.chopWhile(func(x rune) bool { 
			return unicode.IsLetter(x) || unicode.IsDigit(x)
		})
		return result, true
	}
	return self.chop(1), true
}