package main

import "github.com/ho3eintry/go-solutions/section-01/temp_files"

func main() {
	if err := temp_files.WorkWithTemp(); err != nil {
		panic(err)
	}
}
