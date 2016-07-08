package foo

import "fmt"

// Foo type.
// Another line.
// And one more.
//
// Example:
//  f := &File{
//    A: "a-field",
//    B: "b-field",
//  }
type Foo struct {
	// A field doc
	A string

	// B field doc
	B string
}

func (f *Foo) Bar() {
	fmt.Println("Bar method")
}
