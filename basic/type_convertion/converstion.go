
package main

import (
	. "fmt"
)

func main() {
	exampleEncode()
	exampleTypeAssertion()
	exampleTypeConvertion()
}

func exampleEncode() {
	str := "百度一下，你就知道"
	Println("string len:", len([]rune(str)))
	Println("Byte len", len(str))
}

var container = []string{"zero", "one", "two"}

func exampleTypeAssertion() {
	container := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
	}	

	// type assertion
	value, ok := interface{}(container).([]string)
	Println(value, ok)
	Printf("%T\n", value)

	value2, ok := interface{}(container).(map[int]string)
	Println(value2, ok)
	Printf("%T\n", value2)
	
	elem, err := getElemSafe(container)
	if err != nil {
		Printf("get error %v\n", err)
		return
	}

	Println(elem)

}

func getElemSafe(containerI interface{})(elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = Errorf("Unsupported container type: %T\n", containerI)
		return
	}
	return 
}

func exampleTypeConvertion() {
	
	// int16 => int8
	src := int16(-255)
	dest := int8(src)
	Println(dest)

	// int => string, value should be a valid code point of string
	Println(string(0))
	Println(string(-1))
}
