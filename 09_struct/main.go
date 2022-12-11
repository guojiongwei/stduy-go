package main

import "fmt"

type IInfo struct {
	a string
	b int
}

/** 定义类型 */
type IPerson struct {
	name string
	age  int
	sex  string
	info IInfo
}

func changeValue(person IPerson) {
	person.age = 27
	fmt.Println("值传递person=", person)
}

func changeValue1(person *IPerson) {
	person.age = 27
	fmt.Println("引用传值person=", person)
}

// struct定义的值，默认是值传递

func main() {
	var person1 IPerson
	person1.name = "郭炯韦"
	person1.age = 26
	person1.sex = "男"
	person1.info.a = "a"

	fmt.Println("person1=", person1)
	fmt.Println("person1.info.a=", person1.info.a)
	if person1.info.b == 0 {
		fmt.Println("person1.info.b=空")
	}
	changeValue(person1)
	fmt.Println("值传递改变后的person1=", person1)

	changeValue1(&person1)
	fmt.Println("引用传递改变后的person1=", person1)
}
