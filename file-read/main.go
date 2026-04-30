package main

import (
	"fmt"
	"os"
	"io"
)
var p = fmt.Println

func main () {
	p(os.Args, "<<")
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		
		file.Close()
	}
}
