package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var p = fmt.Println

func main() {
	if len(os.Args) == 1 {
		panic("you didn't pass parameters")
	}
	fileName := os.Args[1]
	cwd, _ := os.Getwd()
	fullPath := filepath.Join(cwd, fileName)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}
	stringData := string(data)

	codeStart := strings.Index(stringData, "```go\n")
	if codeStart == -1 {
		panic("there is no ```go section")
	}
	afterOpen := codeStart + len("```go\n")
	codeEnd := strings.Index(stringData[afterOpen:], "```")

	if codeEnd == -1 {
		panic("there is no ``` section")
	}
	content := stringData[afterOpen:afterOpen+codeEnd]
	p(content)

	err = os.WriteFile(filepath.Join(filepath.Dir(fullPath), "main.go"), []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
