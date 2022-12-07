package main

import (
	func_demo "05_func/func_demo"
	"fmt"
)

func main() {
	func_demo.Base(1, 2)
	func_demo.PointerValue()
	func_demo.PointerQuote()
	var v1, v2 = func_demo.ManyReturn(1, 2)
	fmt.Println("v1, v2:", v1, v2)

	var nameReturn = func_demo.NameReturn()
	fmt.Println("nameReturn:", nameReturn)
}
