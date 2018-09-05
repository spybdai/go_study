
package main

import (
	. "fmt"
	"unicode/utf8"
)

func main() {

	Println(string(-1))

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	Println(sample)
	Printf("%T\n", sample) 
	Println(len(sample))

	for i:=0; i<len(sample); i++ {
		Printf("%x ", sample[i])
	}
	Printf("%x\n", sample)
	Printf("% x\n", sample)
	Printf("%q\n", sample)
	Printf("%+q\n", sample)

	const placeOfInterest = `⌘`
	Println(placeOfInterest) 
	Printf("%T\n", placeOfInterest) 
	Println(len(placeOfInterest))

	Printf("%s\n", placeOfInterest)
	Printf("%+q\n", placeOfInterest)
	Printf("% x\n", placeOfInterest)

	Println("\u0061")
	Println("\u00e0")
	Println("\u0061\u0300")

	s1 := '⌘'
	Println(s1)
	Printf("%T\n", s1) 
	//Println(len(s1))
	Printf("%x\n", s1)
	Printf("%b\n", s1)

	const nihongo = "你好"
	Println(nihongo)
	Printf("%T\n", nihongo)
	Println(len(nihongo))

	//a for range loop, decodes one UTF-8-encoded rune on each iteration.
    	for index, runeValue := range nihongo {
      		Printf("%#U starts at byte position %d\n", runeValue, index)
    	}
	
	//with utf8 package
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
	//a for range loop, decodes one UTF-8-encoded rune on each iteration.
    	for index, runeValue := range nihongo {
      		Printf("%x starts at byte position %d\n", runeValue, index)
    	}
	
	for i:=0; i<len(nihongo); i++ {
		Printf("%x", nihongo[i])
	}

	//a for range loop, decodes one UTF-8-encoded rune on each iteration.
	// if invalid UTF-8 code point, U+FFFD will be used to replace 
    	for index, runeValue := range sample{
      		Printf("%#U starts at byte position %d\n", runeValue, index)
    	}
}

