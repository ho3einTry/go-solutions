package main

import (
	"fmt"
	"github.com/ho3eintry/go-solutions/section-02/ansiColor"
)

func main() {
	r := ansiColor.ColorText{
		TextColor: ansiColor.Red,
		Text:      "I'm red!",
	}

	fmt.Println(r.String())

	r.TextColor = ansiColor.Green
	r.Text = "Now I'm green!"

	fmt.Println(r.String())

	r.TextColor = ansiColor.ColorNone
	r.Text = "Back to normal..."

	fmt.Println(r.String())
}
