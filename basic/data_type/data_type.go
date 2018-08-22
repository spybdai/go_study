// example of pointer, struct, slice and maps

package main

import (
	. "fmt"
	"strings"
	"math"
)

func main(){
	examplePointer()
	exampleStruct()
	exampleArray()
	exampleSlice()
	exampleMaps()
	exampleFunctionValues()
}

func examplePointer(){
	i, j := 1, 2

	Println(i, j)
	Println(&i, &j)
	Println(*&i, *&j)
}

func exampleStruct(){

	type Vertex struct {
		X int
		Y int
	}
	
	ver := struct { X, Y int} {3, 4}

	Println(Vertex{1, 2})
	Println(ver)
	println(ver.X, ver.Y)

	p := &ver
	p.X = 5 //(*p).X is cumbersome, so it's permitted to write just p.X
	Println(ver)
	(*p).X = 6
	Println(ver)

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		q = &Vertex{3, 4}
	)

	Println(v1, v2, v3, q)

}

func exampleArray(){
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	Println(a[0], a[1])
	Println(a)
}

func exampleSlice(){
	primes := [6]int {1, 2, 3, 4, 5, 6}
	s := primes[2:4]
	var s1 = primes[3:5]

	Println(s)
	Println(s1)
	
	s[0] = 1000

	Println(s)
	Println(primes)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{7, true},
		{11, false},
	}
	
	Println(s2)
	Println(s2[:])

	println(len(s[:]), cap(s[:]))
	println(len(s[:1]), cap(s[:1]))
	println(len(s[1:2]), cap(s[1:2]))

	s2 = s2[:4]
	Println(len(s2), cap(s2), s2)
	s2 = s2[:0]
	Println(len(s2), cap(s2), s2)
	s2 = s2[:2]
	Println(len(s2), cap(s2), s2)
	s2 = s2[2:]
	Println(len(s2), cap(s2), s2)

	// nil slice

	var s3 []int
	Println(len(s3), cap(s3), s3)
	if s3 == nil {
		Printf("nil\n")
	}

	// create slice with built-in make
	s4 := make([]int, 5)
	PrintS(s4)
	s5 := make([]int, 0, 6)
	PrintS(s5)
	s5 = s5[:cap(s5)]
	PrintS(s5)
	s5 = s5[1:]
	PrintS(s5)

	// slice of slices

	board := [][]string{
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i:=0; i<len(board); i++ {
		Printf("%s\n", strings.Join(board[i], " "))
	}

	// appening to a slice

	var s6 []int
	PrintS2(s6)	

	//too small to fit all the given values	
	s6 = append(s6, 0)
	PrintS2(s6)	

	//too small to fit all the given values	
	s6 = append(s6, 1, 2, 3)
	PrintS2(s6)	

	s7 := make([]int, 6, 6)
	PrintS2(s7)	
	
	// how the slice are exptended? 
	// looks like the cap will be extended to 2 time of the original
	// cap(new_slice) = 2 * cap(old_slice)
	s7 = append(s7, 5)
	PrintS2(s7)	

	s7 = append(s7, 1, 2, 3)
	PrintS2(s7)	


	s8 := make([]int, 0, 6)
	PrintS2(s8)	
	
	// how the slice are exptended?
	s8 = append(s8, 5)
	PrintS2(s8)	

	s8 = append(s8, 1, 2, 3)
	PrintS2(s8)	

	s8 = append(s8, 1, 2, 3, 5, 6)
	PrintS2(s8)	

	// range & for of slice

	for i, v := range s8 {
		Println(i, v)
	}
	
	// only use index of range
	for i := range s8 {
		Println(i)
	}

	// only use value of range
	for _, v := range s8 {
		Println(v)
	}
}

func PrintS(s []int){
	Println(len(s), cap(s), s)
}

func PrintS2(s []int){
	Println(len(s), cap(s), s)
	Printf("%p\n", s)
}

func exampleMaps(){
	type Vertex struct {
		Lat, Long float64
	}	

	// nil map
	var m map[string]Vertex
	Println(m)
	Printf("%p\n", m)
	
	m1 := make(map[string]Vertex)
	Println(m1)
	Printf("%p\n", m1)

	// this will lead to error, assign value to a nil map
	//m["bell labs"] = Vertex{
	//	11.00, 22.00,
	//}
	//Println(m)

	// make function will make a map ready to use
	m1["bell labs"] = Vertex{
		11.00, 22.00,
	}
	Println(m1)
	Printf("%p\n", m1)

	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	Println(m2)
	Printf("%p\n", m2)

	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m3 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	Println(m3)
	Printf("%p\n", m3)

	// mutating maps

	Println(m3["Google"])
	
	elem, ok := m3["Google"]
	Println(elem, ok)

	delete(m3, "Google")

	elem, ok = m3["Google"]
	Println(elem, ok)

	m3["Google"] = Vertex { 1.0, 2.0}
	elem, ok = m3["Google"]
	Println(elem, ok)
}

func exampleFunctionValues(){
	
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}	
	
	Println(hypot(5, 12))	

	Println(compute(hypot))
	Println(compute(math.Pow))

	// closure
	
	pos, neg := adder(), adder()	

	for i:=0; i<10; i++ {
		Println(pos(i), neg(-2*i))
	}

	// fibonacci with closure
	f := fibonacci()
	for i:=0; i<10; i++ {
		Println(f())
	}
	
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// adder is a closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func fibonacci() func() int {
	x, y, result:= 0, 1, 0
	return func() int {
		result = x
		x = y
		y = result + y
		return result
	}
}
