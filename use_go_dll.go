package main

import (
	"C"
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	dll := syscall.NewLazyDLL("demo_go_dll.dll")
	add := dll.NewProc("Add")
	prt := dll.NewProc("Print")
	r, err, msg := add.Call(9, 18)
	fmt.Printf("%v\n", r)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", msg)

	name := C.CString("misitebao")
	prt.Call(uintptr(unsafe.Pointer(name)))
}
