package main

import (
	"fmt"
	"github.com/ho3eintry/go-solutions/section-03/math"
	"log"
)

func main() {
	math.Examples()

	fmt.Printf("\n****Calculate Fibonacci series***\n")
	fmt.Println("Please insert index of your fibonacci series: ")
	var index int
	_, err := fmt.Scanln(&index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(math.Fib(index))

}
