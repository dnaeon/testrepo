package main

import "fmt"

// File type.
// Another line.
// And one more.
//
// Example:
//  f := &File{
//    A: "a-field",
//    B: "b-field",
//  }
type File struct {
	// A field doc
	A string

	// B field doc
	B string
}

func main() {
	fmt.Println("hello")
}
