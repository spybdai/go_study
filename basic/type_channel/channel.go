
package main

import (
	. "fmt"
	"time"
)


func main() {
	exampleNonBufferChannel()
	exampleNilChannel()
	exampleClosedChannel()
	ch := exampleSingleChannel()
	exampleRangeOnChannel(ch)
	exampleSelectOnChannel()
}

func exampleNonBufferChannel() {
	ch := make(chan int, 0)

	// the print log will help with the schedule log
	go receive(ch)
	go send(ch)

	
	time.Sleep(time.Second)
}


func send(ch chan int) {
	for i:= 0; i<10; i++{
		Printf("sending %v\n", i)
		ch <- i
	}
}

func receive(ch chan int) {
	var v int
	for i:= 0; i<10; i++{
		Printf("receiving %v\n", i)
		v = <- ch
		Println(v)
	}
}

func exampleNilChannel() {
	
	// a nil channel
	// var ch chan int

	// send to/receive from nill channel will block
	
	// panic
	//ch <- 0
	//_ = <- ch

}

func exampleClosedChannel() {

	ch := make(chan int, 1)
	
	ch <- 5

	close(ch)
	
	// will work, even channel is cloesed
	v := <- ch
	Println(v)

	// will work
	// 2nd valures will be used to detect if the channel is closed
	v, ok := <- ch
	Println(v, ok)

	// panic
	//close(ch)	

	// panic
	//ch <- 1

	// no panic, but will get 0 on a closed channel
	v = <- ch
	Println(v)
	v = <- ch
	Println(v)
	v = <- ch
	Println(v)
}

func exampleSingleChannel() <-chan int {

	ch := getSingleChannel()
	Println(ch)
	return ch
}

func getSingleChannel() <-chan int {
	
	ch := make(chan int, 3)
	for i:=0; i<3; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func exampleRangeOnChannel(ch <-chan int) {
	
	// if the channle is not closed, the for...range will block, 
	// if there is no more elements, wich lead to panic
	
	// close not close a recive-only channel
	// close(ch)
	for v := range ch {
		Println(v)
	}
}

func exampleSelectOnChannel() {
	
	Println("exampleSelectOnChannel")
	chs := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	
	for i:=0; i<len(chs); i++ {
		chs[i] <- (i + 10)
	}
	
	
	// close all the channels after 1 microsecond
	time.AfterFunc(time.Microsecond, func() {
		for i:=0; i<len(chs); i++ {
			close(chs[i])
		}
	})

	// will first evaluate all the cases first, 
	// then, select 1 hit to execute

	var flag int

	//for i:=0; i<2; i++ {
	for {
		select {
		case v, ok := <- chs[0]:
			if !ok {
				flag = 1
				Println(0)
				break
			}
			Println(v)
		case v, ok := <- chs[1]:
			if !ok {
				flag = 1
				Println(1)
				break
			}
			Println(v)
		case v, ok := <- chs[2]:
			if !ok {
				flag = 1
				Println(2)
				break
			}
			Println(v)
		default:
			Println("default")
		}

		Println(flag)
		if flag == 1 {
			break
		}
	}
}
