package main

import "fmt"

func main() {

	/** 定义变量 */
	// var name string
	// var age int
	// var sex bool

	/** 批量定义变量 */
	// var (
	// 	name string
	// 	age  int
	// 	sex  bool
	// )

	// fmt.Println("name:", name)
	// fmt.Println("age:", age)
	// fmt.Println("sex:", sex)

	/** 定义变量初始化,	自动类型推断 */
	// var name = "小明"
	// var age = 20
	// var sex = true

	// fmt.Println("name:", name)
	// fmt.Println("age:", age)
	// fmt.Println("sex:", sex)

	/** 定义变量初始化,使用: */
	// name := "小明"
	// age := 20
	// sex := true

	// fmt.Println("name:", name)
	// fmt.Println("age:", age)
	// fmt.Println("sex:", sex)

	/** 匿名变量 */
	name, age, sex := getNameAndAge()
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("sex:", sex)
}

func getNameAndAge() (string, int, bool) {
	return "小明", 20, true
}
