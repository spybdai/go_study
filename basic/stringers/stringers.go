//A Stringer is a type that can describe itself as a string. 
//The fmt package (and many others) look for this interface to print values

//type Stringer interface {
//    String() string
//}

package main

import (
	"fmt"
)

type Persion struct {
	Name string
	Age int
}

func (p Persion) String() string {
	return fmt.Sprintf("%v-%v", p.Name, p.Age)
}

func main(){
	a := Persion{"A", 20}
	b := Persion{"B", 30}

	fmt.Println(a, b)	
	
}
