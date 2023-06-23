package main

import (
	"fmt"
	"github.com/ho3eintry/go-solutions/section-03/currency"
	"log"
)

func main() {
	pennies, err := currency.ConvertStringDollarsToPennies("23.23")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pennies)

	pennies += 80
	dollar := currency.ConvertPenniesToDollarString(pennies)

	fmt.Println(dollar)
}
