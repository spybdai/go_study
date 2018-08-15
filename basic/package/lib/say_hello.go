package lib1

import "fmt"

// Upper case make this function public
// generally the package name SHOULD be same as the folder name, 
// but this example is an exception.
// make sure how to use this

func SayHello(){
	fmt.Printf("Hello from lib1.SayHello\n")
}

