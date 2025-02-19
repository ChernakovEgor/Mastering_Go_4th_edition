package main

import (
	"fmt"
	// "os"
)

func main() {
	// print(os.Args...)
	arr := make([]int, 10)
	arr2 := make([]string, 10)
	print(arr)
	print(arr2)
	print(1, 2, "string", 5)
}

func print(in ...interface{}) {
	for _, v := range in {
		fmt.Println(v)
	}
}
