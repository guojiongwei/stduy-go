package main

import "fmt"

func returnFunc() int {
	fmt.Println("returnFunc")
	return 1
}

/** defer会入栈，先进后出 */
func deferDemo() int {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	return returnFunc()
}

func main() {
	deferDemo()
}
