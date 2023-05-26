package main

import (
	"fmt"

	"github.com/ho3eintry/go-solutions/section-01/csvformats"
)

func main() {
	if err := csvformats.AddMoviesFromText(); err != nil {
		panic(err)
	}
	fmt.Println("**************")
	if err := csvformats.WriteCSVOutput(); err != nil {
		panic(err)
	}
	fmt.Println("**************")
	buffer, err := csvformats.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}
	fmt.Println("**************")
	fmt.Println("Buffer = ", buffer.String())
}
