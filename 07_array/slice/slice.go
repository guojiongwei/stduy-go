package slice

import "fmt"

func SliceDemo() {

	/** 声明slice1是一个切片，长度为4 */
	slice1 := []int{1, 2, 3, 4}
	fmt.Println("slice1=", slice1)

	/** 声明slice1是一个切片，长度为4 */
	var slice2 = make([]int, 3)
	slice2[0] = 1
	fmt.Println("slice2=", slice2, len(slice2))

	/** 声明slice1是一个切片，长度为4 */
	slice3 := make([]int, 3)
	slice3[0] = 1
	fmt.Println("slice3=", slice3, len(slice3))

	/** 切片容量 */
	slice4 := make([]int, 2, 4)
	// slice4[len(slice4)-1]
	slice4[1] = 1
	slice4 = append(slice4, 2)
	slice4 = append(slice4, 3)
	fmt.Println("slice4=", slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, 4)
	slice4 = append(slice4, 5)
	fmt.Println("slice4=", slice4, len(slice4), cap(slice4))

	/** 截取slice */
	slice5 := slice4[0:2]
	fmt.Println("slice5=", slice5)
	slice6 := slice4[2:4]
	fmt.Println("slice6=", slice6)
	slice7 := slice4[2:]
	fmt.Println("slice7=", slice7)
}
