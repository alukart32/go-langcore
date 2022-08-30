package types

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Strings() {
	s := `
type _string struct {
	elements *byte // underlying bytes
	len      int   // number of bytes
}`

	fmt.Println("Internal structure::\n", s)
}

func StringOper() {
	var s1 string = "abc"
	s2 := "def"
	fmt.Printf("s1: %s + s2: %s = %s\n", s1, s2, s1+s2)

	fmt.Printf("s1 > s2: %v\n", s1 > s2)
	fmt.Printf("s1 < s2: %v\n", s1 < s2)
	fmt.Printf("s1 => s2: %v\n", s1 >= s2)
	fmt.Printf("s1 <= s2: %v\n", s1 <= s2)
	fmt.Printf("s1 == s2: %v\n", s1 == s2)
}

func StringLen() {
	const s = "Go101.org" // len(s) == 9

	// len(s) is a constant expression,
	// whereas len(s[:]) is not.
	var a byte = 1 << len(s) / 128
	var b byte = 1 << len(s[:]) / 128

	fmt.Println(a, b) // 4 0
}

func runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func StringConversionByteSliceToRune() {
	s := "Color Infection is a fun game."
	bs := []byte(s)      // string -> []byte
	s = string(bs)       // []byte -> string
	rs := []rune(s)      // string -> []rune
	s = string(rs)       // []rune -> string
	rs = bytes.Runes(bs) // []byte -> []rune
	bs = runes2Bytes(rs) // []rune -> []byte
}
