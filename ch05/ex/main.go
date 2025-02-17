package main

import "fmt"

type s1 struct {
	name string
}

type s2 struct {
	count int
}

func main() {
	a := s1{"yep"}
	b := s2{34}
	c := "some var"
	print(a)
	print(b)
	print(c)
}

func print(a interface{}) {
	switch T := a.(type) {
	case s1:
		fmt.Printf("s1 name is %s\n", T.name)
	case s2:
		fmt.Printf("s2 count is %d\n", T.count)
	default:
		fmt.Printf("not a known type\n")
	}
}
