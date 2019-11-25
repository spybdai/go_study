package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	for i := 0; i < 2; i++ {
		go work(i)
	}
	time.Sleep(time.Duration(10 * math.Pow(1000, 3)))
}

func work(index int) {
	go func() {
		count := 0
		for {
			time.Sleep(time.Duration(math.Pow(1000, 3)))
			count += 1
			fmt.Println(index, ":", count)
		}
	}()

	fmt.Println(index, "I am quit")
}
