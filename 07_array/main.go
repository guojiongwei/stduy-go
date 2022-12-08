package main

import (
	slice "07_array/slice"
	"fmt"
)

func rangeArrValue(arr *[4]int) {
	/** 固定数组的传值默认是值传递 */
	arr[3] = 10
	/** range遍历数组 */
	for index, value := range arr {
		fmt.Println("arr", index, value)
	}
}

func ArrayDemo() {
	arr1 := [4]int{1, 2}
	arr1[3] = 4
	rangeArrValue(&arr1)
	/** for循环数组 */
	for i := 0; i < len(arr1); i++ {
		fmt.Println(i, arr1[i])
	}

	slice.SliceDemo()
}

func main() {
	ArrayDemo()
}
