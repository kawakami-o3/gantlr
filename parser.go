package gantlr

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
