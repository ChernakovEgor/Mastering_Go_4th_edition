package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	f2 string
	f3 float64
	f4 []int
}

func main() {
	t := T{3, "yep", 3.14, []int{1, 2, 3, 4}}
	tp := reflect.ValueOf(&t)
	tv := tp.Elem()
	for i := 0; i < tv.NumField(); i++ {
		f := tv.Field(i)
		if f.Type().Kind() == reflect.Int {
			f.SetInt(10)
		}
	}
	fmt.Println(t)
}
