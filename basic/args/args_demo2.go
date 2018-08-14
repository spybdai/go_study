
package main

import (
	"os"
	"fmt"
	"flag"
)

var name string
var age = flag.String("age", "20", "age")

func init(){

	flag.StringVar(&name, "name", "everyone", "greeting object")

	flag.Usage = func(){ // customerize help function
		fmt.Fprintf(os.Stderr, "Customerize help message:\n")
		flag.PrintDefaults()
	}
}

func main(){
	flag.Parse()
	fmt.Printf("Hello, %s, %s!\n", name, *age)
}
