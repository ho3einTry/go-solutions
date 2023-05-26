package main

import "github.com/ho3eintry/go-solutions/section-01/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	bytestrings.SearchString()
	bytestrings.ModifyStrings()
	bytestrings.StringReader()
}
