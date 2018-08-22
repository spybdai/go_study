// explain how the scope of var and how GO search var

//首先，代码引用变量的时候总会最优先查找当前代码块中的那个变量。注意，这里的“当前代码块”仅仅是引用变量的代码所在的那个代码块，并不包含任何子代码块。
//其次，如果当前代码块中没有声明以此为名的变量，那么程序会沿着代码块的嵌套关系，从直接包含当前代码块的那个代码块开始，一层一层地查找。
//一般情况下，程序会一直查到当前代码包代表的那层代码块。如果仍然找不到，那么 Go 语言的编译器就会报错了。

package main

import . "fmt" // can use fmt.Printf as Printf 

var block = "package"
var container = []string {"1", "2", "3"}

func main(){

	block := "function"
	container := map[int]string {1: "1", 2: "2", 3: "3"}

	{
		block := "inner"
		Printf("%s\n", block)
		Printf("%q\n", container[1])
	}

	Printf("%s\n", block)	
	Printf("%q\n", container[1])
}

// this will lead to re-declare, since there a Printf from fmt
//func Printf(){}
