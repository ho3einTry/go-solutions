package flags

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays Named types
type CountTheWays []int //is name typing

func (c *CountTheWays) String() string {
	result := ""
	for _, v := range *c {
		if len(result) > 0 {
			result += fmt.Sprint("...")
		}
		result += fmt.Sprint(v)
	}
	return result
}

func (c *CountTheWays) Set(value string) error {
	values := strings.Split(value, ",")
	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}
	return nil
}
