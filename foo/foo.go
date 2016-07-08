package foo

import "fmt"

// Foo type.
// Another line.
// And one more.
//
// Example:
//  f := &Foo{
//    A: "a-field",
//    B: "b-field",
//  }
type Foo struct {
	// A field doc
	A string

	// B field doc
	B string
}

// Bar documentation.
// Example:
//   foo := &Foo{}
//   foo.Bar()
func (f *Foo) Bar() {
	fmt.Println("Bar method")
}

// File type resource manages files and directories.
//
// Example usage:
//   foo = file.new("/tmp/foo")
//   foo.state = "present"
//   foo.mode = 0600
//   catalog:add(foo)
//
type File struct{}
