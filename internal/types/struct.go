package types

import (
	"fmt"
)

type Book struct {
	author, title string
	pages         int
}

func BookStructExample() {
	book := Book{author: "me", title: "mybook", pages: 1}
	fmt.Println(book)

	_ = Book{
		author: "Tapir",
		title:  "Go 101",
		pages:  256, // here, the "," must be present
	}
	// None of the fields are specified. The title and
	// author fields are both "", pages field is 0.
	book = Book{}

	// Only specify the author field. The title field
	// is "" and the pages field is 0.
	book = Book{author: "Tapir"}

	var book2 Book
	book2.author = "not_me"
	book2.title = "not_my_book"
	fmt.Println(book2)
}
