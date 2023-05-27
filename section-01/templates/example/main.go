package main

import (
	"fmt"

	"github.com/ho3eintry/go-solutions/section-01/templates"
)

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}
	fmt.Println("**************")
	if err := templates.InitTemplate(); err != nil {
		panic(err)
	}
	fmt.Println("**************")
	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}
