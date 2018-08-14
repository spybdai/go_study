
package main

import (
	"os"
	"fmt"
	"flag"
)

var name string
var age *string

func init(){

	// replace flag.CommandLine
	// flag.CommandLine is like a "container", which execute flag cmds 
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func(){
		fmt.Fprintf(os.Stderr, "Customerize flag.CommandLine\n")
		flag.PrintDefaults()
	}

	age = flag.String("age", "20", "age")
	flag.StringVar(&name, "name", "everyone", "greeting object")
}




func main(){
	flag.Parse()
	fmt.Printf("Hello, %s, %s!\n", name, *age)
}
