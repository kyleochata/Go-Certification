package main

import (
	"fmt"
	"io"
	"os"
)

/*
type Writer interface {
	Write(p []byte) (n int, err error)
	}
	*File == Writer interface
	can pass that pointer to the file straight to other methods that requires a Writer interface for params.
	func Fprintln( w Writer, a interface{})(n int, err error)

	func NewFile (f File) *File --> *File has the method Write() attached
	that means that the pointer to the File can be used anywhere it requires a Writer interface as an argument

*/

func main() {
	// Fprintln( w io.Writer, a ...any)(n int, err error)
	fmt.Fprintln(os.Stdout, "Hello world")

	//demands (w Writer, s string)(n int, err error)
	io.WriteString(os.Stdout, "from writeString")
}
