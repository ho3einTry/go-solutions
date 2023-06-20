package dataconv

import "fmt"

func ShowConv() {
	var a, b = 24, 2.0

	c := float64(a) * b

	fmt.Println(c)

	precision := fmt.Sprintf("%.4f", b)
	fmt.Printf("%s - %T\n", precision, precision)

}
