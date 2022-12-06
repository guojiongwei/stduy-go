package main

import (
	basetype "04_data_type/base_type"
	fmt "fmt"
)

const (
	BEIJIN   = iota
	SHANGHAI = iota
)

func main() {
	fmt.Println("hello world!")

	basetype.BoolType()

	basetype.IntType()

	basetype.StrigType()

	basetype.ConstIota()
	fmt.Println("BEIJIN:", BEIJIN)
	fmt.Println("SHANGHAI:", SHANGHAI)

	fmt.Println("Test", basetype.Test)
}
