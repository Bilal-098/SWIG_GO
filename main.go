package main

import (
	"fmt"
	"github.com/Bilal-098/SWIG_GO/tree/main/sum"
)

func main() {

	a := 5
	b := 10
	ss := sum.sum_of_two_num(a, b)
	fmt.Printf("Go says: Sum is %v\n", ss)
}
