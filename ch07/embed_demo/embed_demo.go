package main

import (
	_ "embed"
	"fmt"
)

//go:embed embed_demo.go
var sourceCode string

func main() {
	fmt.Println(sourceCode)
}
