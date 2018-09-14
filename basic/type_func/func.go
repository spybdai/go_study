
package main

import (
	. "fmt"
	"errors"
)

func main() {
	var op1 operate
	x, y := 1, 2
	value, ok := Caculate(x, y, op1)
	Println(x, y, value, ok)

	value, ok = Caculate(x, y, add)
	Println(x, y, value, ok)
	
	add := genCaculator(add)
	value, ok = Caculate(x, y, add)
	Println(x, y, value, ok)

	examplePassParameter()
}

func add(x, y int) int {
	return x + y
}

// define a func type
type operate func(x, y int) int

// take a func as an parameter
// can work as a decorator similar as python
func Caculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("nil operate")
	}
	return op(x, y), nil
}

// work also as decorate, if want to add some new behaviors to a func,
// while not update the func itself
func genCaculator(op operate) operate {
	fun := func(x, y int) int {
		// do som thing, as a decorate
		// op is from func outside, and it's a closure
		Println("I am decorator")
		return op(x, y)
	}
	return fun
}

func examplePassParameter() {
	a := [3][]string {
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}
	
	Println(a)	

	b := modify(a)
	Println(a)	
	Println(b)	
}

func modify(s [3][]string) [3][]string {
	s[1][1] = "tt"
	return s
}
