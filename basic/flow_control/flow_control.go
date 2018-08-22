
package main

import (
	"fmt"
	"runtime"	
	"time"
)


func main(){
	forExample()
	ifExample()
	switchExample()
	deferExample()
	deferStackExample()
}


func forExample(){

	// format	
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statemens are optional
	sum = 1
	for ; sum <= 10; {
		sum += sum
	}
	fmt.Println(sum)

	// work as "while" without semicolons
	sum = 1
	for sum <= 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// an infinite loop
	//for {
	//	fmt.Printf("High way to the hell...\n")
	//}

}

func ifExample(){

	// normal if
	i := 1
	if i == 1 {
		fmt.Printf("Normal if statement\n")
	} else {
		fmt.Printf("Normal if statement, with else\n")
	}

	// if with short statement before the condition
	if j := 1; j == 1 {
		fmt.Printf("if with short statement before condition\n")
	}

}


func switchExample(){
	fmt.Printf("Go runs on: ")
	
	switch os:= runtime.GOOS; os {
	case "darwin": 
		fmt.Printf("Mac\n")
	case "linux": 
		fmt.Printf("Linux\n")
	default:
		fmt.Printf("Windows\n")
	}

	switch os:=1; os{
	case 1:
		fmt.Println(os)
	}

	fmt.Printf("When's Saturday?\n")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Printf("today\n")
	case today + 1:
		fmt.Printf("tomorrow\n")
	case today + 2:
		fmt.Printf("in 2 days\n")
	default:
		fmt.Printf("far away\n")
	}

	// siwtch without condition, equal to switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Printf("Good morning\n")
	case t.Hour() < 17:
		fmt.Printf("Good afternoon\n")
	default:
		fmt.Printf("Good evening\n")
	}

}


func deferExample(){
	defer fmt.Printf("world\n")
	fmt.Printf("Hello ")
}

func deferStackExample(){
	fmt.Printf("counting:\n")

	for i:=0; i<10; i++ {
		defer fmt.Printf("index %d\n", i)
	}

	fmt.Printf("done\n")
}
