package main

import (
	"fmt"
	"reflect"
)

type I interface{}

type A struct {
	Name   string
	Age    *int
	Scores []int32
}

func main() {
	var (
		a A
		i I
	)
	var (
		c = A{}
	)
	ta := reflect.ValueOf(a)
	tc := reflect.ValueOf(c)
	fmt.Println(ta, tc)
	fmt.Println(ta.IsValid())
	fmt.Println(reflect.ValueOf(i))

	ExampleSetable()
}
