package dataconv

import (
	"fmt"
	"strconv"
)

func Strconv() error {
	res, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(res)

	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}
	fmt.Println(res)

	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}
	fmt.Println(val)
	return nil
}
