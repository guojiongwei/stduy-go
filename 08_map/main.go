package main

import (
	"fmt"
)

/** 遍历map */
func rangeMap(map5 map[int]string) {
	map5[2] = "2"
	for index, value := range map5 {
		fmt.Println("index=", index)
		fmt.Println("value=", value)
	}
}

func main() {

	/** 声明map1， key是string，value是string */
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1是空对象")
	}

	/** 使用make声明，初始化是空对象 */
	map2 := make(map[string]string)
	map2["a"] = "a"
	map2["b"] = "b"
	fmt.Println("map2=", map2, map2["a"])

	/** 使用make声明，带初始值 */

	var map3 = map[int]string{
		1: "1",
		2: "2",
	}

	fmt.Println("map3=", map3)
	/** 使用make声明，带初始值 */

	var map4 = make(map[int]string, 3)
	map4[1] = "a"

	fmt.Println("map4=", map4)
	/** 删除map值 */
	delete(map4, 1)
	fmt.Println("map4=", map4)
	map4[3] = "3"
	/** map类型是引用传递 */
	rangeMap(map4)
	fmt.Println("map4=", map4)
}
