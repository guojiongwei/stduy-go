package basetype

import "fmt"

func BoolType() {
	var b1 bool = true
	var b2 bool = false

	var age = 18

	if age == 18 {
		fmt.Println("b1==true")
	} else {
		fmt.Println("b1==false")
	}

	var b3 = true
	var b4 = false

	b5 := true
	b6 := false

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("b4:", b4)
	fmt.Println("b5:", b5)
	fmt.Println("b6:", b6)
}
