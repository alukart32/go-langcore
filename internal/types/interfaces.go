package types

import (
	"fmt"
)

// This interface directly specifies two methods and
// embeds two other interface types, one of which
// is a type name and the other is a type literal.
type ReadWriteCloser = interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	error                      // a type name
	interface{ Close() error } // a type literal
}

// This interface embeds an approximation type. Its type
// set inlcudes all types whose underlying type is []byte.
type AnyByteSlice = interface {
	~[]byte
}

// This interface embeds a type union. Its type set inlcudes
// 6 types: uint, uint8, uint16, uint32, uint64 and uintptr.
type Unsigned = interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Aboutable interface {
	About() string
}

type Article struct {
	name string
	// more other fields ...
}

func (book *Article) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "I'm a custom integer value"
}

func InterfacesBoxing() {
	// A *Book value is boxed into an
	// interface value of type Aboutable.
	article := &Article{"my_book"} // tmp := Book{...}, p := &tmp
	var about Aboutable = article
	fmt.Println(about)

	// bi is a blank interface value.
	var bi interface{} = &Article{"Java"}
	fmt.Println(bi)

	// Aboutable implements interface{}.
	bi = about
	fmt.Println(bi)
}

func BoxingValuesWithBlankInterface() {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i)

	i = map[rune]int{'1': 1}
	fmt.Println(i)

	// Clear the boxed value in interface value i.
	i = nil
	fmt.Println(i)
}
