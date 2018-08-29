
package main

import (
	. "fmt"
	"time"
)


func main() {
	//exampleChannel()
	//exampleFibonacci()
	//exampleSelect()
	//exampleSelectDefault()
	Resolve()
	
	
}

func treeSolution() {
	
}

func exampleSelectDefault() {
	tick := time.Tick(2 * time.Second)
	boom := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			Println("tick .")	
		case <-boom:
			Println("boom!")
			return
		default:
			Println(".")
			time.Sleep(time.Second)
			Println("sleeped")
		}
	}
}

func exampleSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i:=0; i<10; i++ {
			Println(<-c)
		}
		quit <- 0
	}()

	go fibonacci2(c, quit)

	time.Sleep(time.Second)
}

// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. 
// It chooses one at random if multiple are ready.
func fibonacci2(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			Printf("quit")
			return
		}
	}
}

func exampleFibonacci() {
	c := make(chan int, 10)
	
	go fibonacci(cap(c), c)	

	// he loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		Println(i)
	}
}

// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// Channels aren't like files; you don't usually need to close them. 
// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for count:=0; count<n; count++ {
		c <- x	
		x, y = y, x+y
	}
	close(c)
}

func exampleChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	Println(<-ch)
	Println(<-ch)

	ch <- 1
	ch <- 2

	//will get error
	//fatal error: all goroutines are asleep - deadlock!
	//ch <- 3

	Println(<-ch)
	Println(<-ch)

	//will get error
	//fatal error: all goroutines are asleep - deadlock!
	//Println(<-ch)
	
	ch2 := make(chan int, 5)
	go produce(ch2)
	go consume(ch2)
	
	for count:=0; count<10; count++ {
		Printf("wait %v times\n", count)	
		time.Sleep(time.Second)
	}

}

func produce(ch chan int) {
	for count:=0; ; count++ {
		Printf("Producing...\n")
		ch <- count
		Printf("Produced: %v\n", count)
		time.Sleep(1000*time.Millisecond)
	}
}

func consume(ch chan int) {
	for {
		Printf("Consume...\n")
		value := <- ch
		Printf("Consumed: %v\n", value)
		time.Sleep(1000*time.Millisecond)
	}
}
