package main

import "github.com/ho3eintry/go-solutions/section-03/dataconv"

func main() {
	dataconv.ShowConv()

	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
