package main

import "fmt"

func deferDemo() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

func main() {
	deferDemo()
}
