package main

import "fmt"

// 寻找局部最小值
// 一个无序数组，任意相邻两个元素不相等，找到一个局部最小的值

func getLessIndex(intList []int) int {

	if intList == nil || len(intList) == 0 {
		return -1
	}

	// 判断 最左是不是最小
	if intList[0] < intList[1] {
		return 0
	}

	l := len(intList)

	// 判断最右是不是最小
	if intList[l-1] < intList[l-2] {
		return l - 1
	}

	start := 0
	mid := 0 + (len(intList)-1)/2

	for start != mid {
		if intList[mid] > intList[mid+1] {

		}
	}

	// TODO

	return -1

}

func main() {
	intList := []int{2, 1, 3}
	fmt.Println(getLessIndex(intList))
}
