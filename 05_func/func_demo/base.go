package func_demo

import "fmt"

/** 基本函数 */
func Base(a int, b int) int {
	fmt.Println("a+b=", a+b)
	return a + b
}

/** 多函数返回值 */
func ManyReturn(a int, b int) (int, string) {
	return a + b, "多返回值"
}

/** 命名返回值 */
func NameReturn() (r int) {
	r = 111
	return
}
