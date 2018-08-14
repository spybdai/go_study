
package main

import (
	"os"
	"fmt"
	"flag"
)

var name string
var age *string
var ourOwnCmd = flag.NewFlagSet("", flag.ExitOnError)

func init(){

	// drop flag.CommandLine and use our own
	// note: commands of flag are actually executed by flag.CommandLine

	ourOwnCmd.Usage = func(){
		fmt.Fprintf(os.Stderr, "Use our ourOwnCmd\n")
		ourOwnCmd.PrintDefaults()
	}

	age = ourOwnCmd.String("age", "20", "age")
	ourOwnCmd.StringVar(&name, "name", "everyone", "greeting object")
}




func main(){
	ourOwnCmd.Parse(os.Args[1:])
	fmt.Printf("Hello, %s, %s!\n", name, *age)
}
