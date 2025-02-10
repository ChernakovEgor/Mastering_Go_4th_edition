package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3}
	a2 := [...]int{4, 5, 6}
	fmt.Println(arraysToSlice(a1, a2))
}

func arraysToSlice[T any](a1 [3]T, a2 [3]T) []T {
	res := make([]T, len(a1)+len(a2))
	for i := range a1 {
		res[i] = a1[i]
		res[i+len(a1)] = a2[i]
	}
	return res
}
