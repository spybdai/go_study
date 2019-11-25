package go_encode

import "fmt"

func UnderStandGoString() {
	const nihongo = "\xbd\xbd日本語\xbd"

	fmt.Printf("% x\n", nihongo)
	fmt.Printf("% q\n", nihongo)
	fmt.Printf("% +q\n", nihongo)

	for i := 0; i < len(nihongo); i++ {
		fmt.Printf("%#U, ", nihongo[i])
	}

	fmt.Print("\n")
	for index, runeValue := range nihongo {
		fmt.Printf("%#U, %d\n", runeValue, index)
	}
	for index, runeValue := range nihongo {
		fmt.Printf("%x, %d\n", runeValue, index)
	}

	const t = '日'
	fmt.Printf("%x\n", t)

	const t1 = 189
	fmt.Printf("%x\n", t1)

	const t2 = '\xbd'
	fmt.Printf("%x\n", t2)

}
