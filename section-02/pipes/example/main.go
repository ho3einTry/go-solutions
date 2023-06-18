package main

import (
	"fmt"
	"github.com/ho3eintry/go-solutions/section-02/pipes"
	"os"
)

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range pipes.WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
