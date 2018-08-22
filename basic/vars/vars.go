package main

import (
	"fmt"
	"flag"
)


func main(){
	
	// way 1
	var name string
	flag.StringVar(&name, "name", "Robert", "name of user")

	// way 2, 
	var age = * flag.String("age", "20", "age of user")

	// way 3, short variable declarations, only used in function block
	comp := name + " " + age
	
	fmt.Printf("My name is %s\n", name)	
	fmt.Printf("My age is %s\n", age)	
	fmt.Printf("My name and age is %s\n", comp)	

	// variable re-declaration, redeclaration can only appear in a multi-variable short declaration
	comp_info, comp := "test", name + " " + age
	fmt.Printf("My name and age is %s, %s, again\n", comp_info, comp)	
	
	// variable re-declarartion, doesn't work
	// var name, age = "Jeff", "30"
	
	// variable re-declarartion, doesn't work
	//  var name, new_age = "Jeff", "30"

	printName()

	fmt.Printf("My name is %s\n", name)	
}

func printName(){
	name := "inner_func_name"
	fmt.Printf("Name inner func is %s\n", name)
}
