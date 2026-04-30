package main

import (
	"fmt"
)
var p = fmt.Println
func main() {
	data := []int{1,2,3}
	for i := range data {
		data = append(data, 10)
		p(i)
	}
	// vs
	for i := 0; i < len(data); i++ {
		data = append(data, 10)
		p(i)
	}
}