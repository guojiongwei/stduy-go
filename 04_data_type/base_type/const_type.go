package basetype

import "fmt"

const (
	GUANGZHOU = iota
	SHENZHEN  = iota
)

func ConstIota() {
	fmt.Println("GUANGZHOU:", GUANGZHOU)
	fmt.Println("SHENZHEN:", SHENZHEN)
}

var Test = "test"
