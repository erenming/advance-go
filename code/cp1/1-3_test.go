package cp1

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

func TestForOnString(t *testing.T) {
	s := "hello, 世界"
	forOnString(s, func(i int, r rune) {
		fmt.Printf("i=%#v, r=%#v\n", i, r)
	})
}

func TestSameString(t *testing.T) {
	a := "hello"
	b := "hello"
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&a)).Data)
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&b)).Data)

}
