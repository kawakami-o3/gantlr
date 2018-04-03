package gantlr

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	hoge := []rune(string("hogehoge"))
	fmt.Println('a' == hoge[0])
	fmt.Println('h' == hoge[0])
	fmt.Println("hello")
}
