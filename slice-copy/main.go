package main

import (
	"fmt"
)

var p = fmt.Println

func main() {
	src := []int{1,2,3}
	dst := []int{}
	copy(dst, src)
	p(dst)
}
