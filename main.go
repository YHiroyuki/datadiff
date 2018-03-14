package datadiff

import (
	"fmt"
)

func manin() {
	a := 1
	b := 3
	fmt.Println(Sum(a, b))
}

func Sum(a, b int) int {
	return a + b
}
