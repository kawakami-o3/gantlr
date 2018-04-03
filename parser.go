package gantlr

import "io"

type Lexer struct {
	io.ByteReader

	c byte
	//input string // input string
	/*
		input []byte // input string
		p     int    // current position
		c     byte   // current character
	*/

	//const EOF = "eof"
}

func (self *Lexer) consume() (byte, error) {
	c, err := self.ReadByte()
	if err != nil {
		return 0, err
	}
	self.c = c
	return c, nil
}

/*
func (self *Lexer) match(x byte) {
	if c == x {
		consume()
	} else {

	}
}
*/

const (
	NAME   = 2
	COMMA  = 3
	LBLACK = 4
	RBLACK = 5
)

var tokenNames = []string{
	"n/a",
	"<EOF>",
	"NAME",
	"COMMA",
	"LBRACK",
	"RBRACK",
}

type Token struct {
	typeId int
	text   string
}

func (t *Token) toString() string {
	return ""
}

func isLetter(s string) bool {
	//c := s[0]
	return true
}
