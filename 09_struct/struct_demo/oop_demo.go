package structdemo

import "fmt"

/** 定义类型 */
type iPersonOpp struct {
	name string
	age  int
	sex  string
}

func (p *iPersonOpp) changeAge(age int) {
	p.age = age
}

func (p *iPersonOpp) changeName(name string) {
	p.name = name
}

/** 继承基类 */
type iChild struct {
	iPersonOpp
	child bool
}

func OppDemo() {

	var person1 iPersonOpp
	person1.name = "郭炯韦"
	person1.age = 26
	person1.sex = "男"
	person1.changeName("新名称")
	person1.changeAge(20)

	fmt.Println("person1=", person1)
	fmt.Println("引用传递改变后的person1=", person1)

	var person2 = iPersonOpp{"郭炯韦", 20, "男"}
	fmt.Println("初始化赋值=", person2)

	/** 继承基类后赋值 */
	var child = iChild{iPersonOpp{"郭炯韦", 20, "男"}, true}
	fmt.Println("child=", child)
}
