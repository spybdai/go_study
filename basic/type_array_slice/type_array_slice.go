// show examples of slice and array

package main

import (
	. "fmt"
)

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}	
	s4 := s3[3:6]
	
	//一个切片的容量可以被看作是透过这个窗口最多可以看到的底层数组中元素的个数
	Printf("len: %v, cap: %v\n", len(s4), cap(s4))	
	Println(s4)

	s5 := s4[:cap(s4)]
	Printf("len: %v, cap: %v\n", len(s5), cap(s5))	
	Println(s5)


	s6 := make([]int, 1024)
	info(s6)	

	//在无需扩容时，append函数返回的是指向原底层数组的新切片,而在需要扩容时，append函数返回的是指向新底层数组的新切片
	s6 = append(s6, 1)
	info(s6)	

	s7 := s3[4:6]
	info(s3)
	info(s4)
	info(s7)
	s4[1] = 10
	info(s3)
	info(s4)
	info(s7)
}


func info(s []int) {
	Printf("len: %v, cap: %v\n", len(s), cap(s))	
	Println(s)
}
