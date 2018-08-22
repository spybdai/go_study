
//examples of methods

package main

import (
	. "fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

type MyFloat float64

func main(){
	exampleMethods()	
}


func exampleMethods(){
	v := Vertex{ 3.0, 4.0}
	Println(v.Abs())
	Println(Abs(v))

	// will get error "cannot define new methods on non-local type float64"
	//f := float64(-math.Sqrt2)
	//Println(f.Abs2())

	f := MyFloat(-math.Sqrt2)
	Println(f.Abs3())

	v = Vertex{ 3.0, 4.0}
	v.scale(10)
	Println(v.Abs())

	v = Vertex{ 3.0, 4.0}
	v.scale2(10)
	Println(v.Abs())

	v = Vertex{ 3.0, 4.0}
	p := &v
	Println(p.Abs())
	Println(v.Abs())
	Println(p.Abs2())
	Println(v.Abs2())
	
	v = Vertex{ 3.0, 4.0}
	p = &v
	v.scale(10)
	Println(v)	
	p.scale(10)
	Println(v)	

	v = Vertex{ 3.0, 4.0}
	p = &v
	v.scale2(10)
	Println(v)	
	p.scale2(10)
	Println(v)	

}


// receiver keyword between 'func' keyword and func name
// can both work with Vertex and *Vertex

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Abs2() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// same function as the above method

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// cannot do this
// will get error "cannot define new methods on non-local type float64"
//You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).	
//func (f float64) Abs2() float64 {
//	if f < 0 {
//		return float64(-f)
//	} else {
//		return float64(f)
//	}
//}

func (f MyFloat) Abs3() float64 {

	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}

}

//choosing a value or pointer receiver
//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.


// receiver with pointer
// will change the value pointed by the pointer
func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// note the difference with scale
// will NOT change the value pointed by the pointer
// will change the value copied
func (v Vertex) scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// difference with scale
func scale3(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
