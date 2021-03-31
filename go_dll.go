// go_dll.go
package main

import "C"

//export Add
func Add(a C.int, b C.int) C.int {
	return a + b
}

//export Print
func Print(s *C.char) {

	// 函数参数可以用 string, 但是用*C.char更通用一些。
	// 由于string的数据结构，是可以被其它go程序调用的，
	// 但其它语言（如 nodejs）就不行了

	print("Hello ", C.GoString(s)) //这里不能用fmt包，会报错，调了很久...
}

func main() {
}
