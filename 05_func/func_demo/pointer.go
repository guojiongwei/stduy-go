package func_demo

import "fmt"

/** 值传递 */
func PointerValue() {
	var a = 100
	var b = 200
	wrapValue(a, b)

	fmt.Println("值传递后a, b:", a, b)
}

func wrapValue(c int, d int) {
	fmt.Println("值传递前a, b:", c, d)
	var temp = d
	d = c
	c = temp
}

/** 引用传递 */
func PointerQuote() {
	var a = 100
	var b = 200

	fmt.Println("引用传递前a, b:", a, b)
	wrapQuote(&a, &b)

	fmt.Println("引用传递前a, b:", a, b)
}

func wrapQuote(c *int, d *int) {
	var temp = *d
	*d = *c
	*c = temp
}
