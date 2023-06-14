package main

import (
	"flag"
	"fmt"
	"github.com/ho3eintry/go-solutions/section-02/flags"
)

func main() {
	c := flags.Config{}
	c.Setup()

	flag.Parse()

	fmt.Println(c.GetMessage())
}
