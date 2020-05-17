package main

import (
	"fmt"
	"reflect"
)

func ExampleSetable() {

	var a int = 32

	av := reflect.ValueOf(a)
	at := reflect.TypeOf(a)

	fmt.Printf("Type: %s, Setable: %t\n", at, av.CanSet())

	apv := reflect.ValueOf(&a)
	apt := reflect.TypeOf(&a)

	fmt.Printf("Type: %s, Setable: %t\n", apt, apv.CanSet())

	obj := apv.Elem()
	fmt.Printf("Type: %s, Setable: %t\n", obj.Type(), obj.CanSet())

	fmt.Printf("Origin value: %d\n", a)
	obj.SetInt(int64(33))
	fmt.Printf("new value: %d\n", a)
}
