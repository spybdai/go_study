// examples of interface

package main

import (
	. "fmt"
	"math"
)

// a set of method signatures
// A value of interface type can hold any value that implements those methods.
type Abser interface {
	Abs() float64
}

type MyFloat2 float64

// MyFloat2 implements Abs in Abser
func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct {
	X, Y float64
}

// *Vertex2 implements Abs in Abser
func (v *Vertex2) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main(){
	var a Abser

	f := MyFloat2(math.Sqrt2)
	v := Vertex2{3.0, 4.0}

	a = f //success, since MyFloat2 implements Abs	
	Println(a.Abs())

	a = &v // success, since *Vertex2 implements Abs
	Println(a.Abs())

	// this will lead to error
	//cannot use v (type Vertex2) as type Abser in assignment:
        //Vertex2 does not implement Abser (Abs method has pointer receiver)
	//a = v 

	//interface values
	a = f
	Printf("%v, %T\n", a, a)

	//interface values
	a = &v
	Printf("%v, %T\n", a, a)

	//Interface values with nil underlying values
	// interface value itself is non-nil
	var v2 Vertex2
	a = &v2
	Printf("%v, %T\n", a, a)
	Println(a.Abs())

	// nil interface value
	//A nil interface value holds neither value nor concrete type.
	//Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	var b Abser
	Printf("%v, %T\n", b, b)
	// will lead run-time error
	//b.Abs()


	// empty interface
	var i interface{}
	Printf("%v, %T\n", i, i)
	
	i = 32	
	Printf("%v, %T\n", i, i)

	i = "hello world"
	Printf("%v, %T\n", i, i)
	

	// type assertion
	
	var i2 interface{} = "hello"

	s := i2.(string)
	Println(s)

	s, ok := i2.(string)
	Println(s, ok)

	f2, ok := i2.(float64)
	Println(f2, ok)

	//f3 := i2.(float64) //panic
	//Println(f3)

	// type switch	
	do(21)
	do("hello")
	do(21.00)
	
}

func do(i interface{}) {
	switch v:=i.(type){
	case int:
		Printf("twice %v, %v\n", v, v*2)
	case string:
		Printf("%q leng is %v\n", v, len(v))
	default:
		Printf("Dont know, %v, %T\n", v, v)
	}
}
