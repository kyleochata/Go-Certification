package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
	type Writer interface{
		Write(p []byte)(n int, err error)
	}

Writer is the interface that wraps the basic Write method
*/
type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	//Create(name string)(*File, error)
	//"*" = pointer ==> available to all func(f *File) methods
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s\n", err)
	// }
	// defer f.Close() //open/create a file => defer closing file

	// s := []byte("Hello world!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("my friends!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("This is my line!")
	fmt.Println(b.String())
	b.Reset()
	b.Write([]byte("convert str to []byte"))
	fmt.Println(b.String())

}
