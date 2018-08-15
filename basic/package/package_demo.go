
package main

import (
	"fmt"
	"github.com/spybdai/go_study/basic/package/lib" // foler name is lib, while package name is lib1
	"github.com/spybdai/go_study/basic/package/lib2"
	"github.com/spybdai/go_study/basic/package/lib/lib"
	
	// this is problematic, since 'lib2' is deplicated with the previous lib2
	//"github.com/spybdai/go_study/basic/package/lib2/lib2"
	// set up an alias for 2nd lib2, as lib3
	lib3 "github.com/spybdai/go_study/basic/package/lib2/lib2"

	// internal package only can be refered by:
	// DIRECT parent package and 
	// child packages
	"github.com/spybdai/go_study/basic/package/lib4"

	// import an internal package which obey the rule will lead to build failure
	//in "github.com/spybdai/go_study/basic/package/lib4/internal"
)

func main(){
	fmt.Printf("Hello from main.main\n")	
	// show how to use code entity from same package
	sayHello()

	// note what's lib1?
	lib1.SayHello()
	lib2.SayHello()
	
	// use "github.com/spybdai/go_study/basic/package/lib/lib"
	// which has SAME fole name as 
	// "github.com/spybdai/go_study/basic/package/lib" 
	// but it's OK, since the package is not duplicate
	lib.SayHello()

	lib3.SayHello()
	
	lib4.SayHello()
}

