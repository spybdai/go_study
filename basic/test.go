package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(time.Now().Unix())
	a := uint64(time.Now().UnixNano())
	fmt.Println(a)
	b := a / 1000000
	fmt.Println(b)
	c := a % 1000000000
	fmt.Println(c)

	fmt.Println(b << 32)
	fmt.Println(uint64(math.Pow(2, 64) - 1))
	fmt.Println(uint32(math.Pow(2, 32) - 1))
	fmt.Println(rand.Int31n(10000))

}
