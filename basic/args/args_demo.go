
package main

import (
	"fmt"
	"flag"
)

var name string
var age = flag.String("age", "20", "age")

func init(){
	flag.StringVar(&name, "name", "everyone", "greeting object")
}

func main(){
	flag.Parse()
	fmt.Printf("Hello, %s, %s\n", name, *age)
}
