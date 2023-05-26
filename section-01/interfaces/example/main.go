package main

import (
	"bytes"
	"fmt"

	"github.com/ho3eintry/go-solutions/section-01/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("this is an example *"))
	out := &bytes.Buffer{}

	fmt.Print("stdout on Copy = ")

	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout on PipeExample =\n")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
