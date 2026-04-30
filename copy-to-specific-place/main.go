package main

import (
	"fmt"
)

var p = fmt.Println

func main() {
	data := make([]int, 5)
	additional := []int{1,2}
	
	copy(data[3:5], additional)
	
	p(data)
}
