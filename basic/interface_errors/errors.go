// example of error interface

//type error interface {
//	Error() string
//}

package main

import (
	. "fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// implement method Error
// both of the 2 following forms work

//func (e *MyError) Error() string {
//	return Sprintf("at %v, %s", e.When, e.What)
//}

func (e MyError) Error() string {
	return Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return MyError{
		time.Now(),
		"no work",
	}
}

type ErrNegativeSqrt float64

// not the type converstion from ErrNegativeSqrt to float64
func (e ErrNegativeSqrt) Error() string {
	return Sprintf("cannot Sqrt negtive number: %v\n", float64(e))
}

// return (value, error)
// an exercise
func Sqrt(x float64) (float64, error) {
	var root float64
	var err error

	if x < 0 {
		err = ErrNegativeSqrt(x)
		return root, err
	}
	
	root = 1.0
	for count:=1; count<=10; count+=1 {
		root -= (root * root - x) / (2 * root)
	}
	return root, err
}

func main() {
	if err:=run(); err != nil {
		Println(err)
	}
	
	Println(Sqrt(2))
	Println(Sqrt(-2))
	Println(Sqrt(20))
	Println(Sqrt(-20))
}
