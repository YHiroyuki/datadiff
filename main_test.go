package datadiff

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println("test")
	answer := Sum(1, 3)
	if answer != 4 {
		t.Fatal("error")
	}
}
