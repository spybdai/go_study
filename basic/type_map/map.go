
package main

import (
	. "fmt"
	"time"
	"sync"
)

type SyncMap = sync.Map

func main() {
	exampleNilMap()
	//exampleConcurrent()
	exampleConcurrent2()
}

func exampleNilMap() {
	
	// a nil map
	var m map[string]int

	key := "two"
	
	// get key on a nil map, work
	k, v := m[key]
	Println(m, k, v)

	// delete key on a nil map, work
	delete(m, key)

	// add a key on a nil map, 
	// can complie
	// panic run
	//m[key] = 2
}

func exampleConcurrent() {

	m := map[int]string {
		1: "start",
	}

	go read(m)
	time.Sleep(time.Second)
	go write(m)
	time.Sleep(30 * time.Second)
	Println(m)	
}

func read(m map[int]string) {
	for {
		_ = m[1]
		time.Sleep(1)
	}
}

func write(m map[int]string) {
	m[1] = "test"
	time.Sleep(1)
}

func exampleConcurrent2() {

	var m SyncMap
	
	m.Store(1, "start")
	m.Store("start", 2)

	go read2(m)
	time.Sleep(time.Second)
	go write2(m)
	time.Sleep(10 * time.Second)
	Println(m)	
}

func read2(m SyncMap) {
	for {
		_, _ = m.Load(1)
		time.Sleep(1)
	}
}

func write2(m SyncMap) {
	m.LoadOrStore(1, "test")
	time.Sleep(1)
}
