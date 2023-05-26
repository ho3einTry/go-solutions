package main

import (
	"fmt"

	"github.com/ho3eintry/go-solutions/section-01/filedirs"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		fmt.Println(err)
	}
}
