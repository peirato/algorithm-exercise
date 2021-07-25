package main

import (
	"crypto/rand" //真随机
	"fmt"
	"math/big"
)

// 寻找局部最小值
// 一个无序数组，任意相邻两个元素不相等，找到一个局部最小的值
// 用二分法

func getLessIndex(intList []int) int {

	if intList == nil || len(intList) == 0 {
		return -1
	}

	if len(intList) == 1 {
		return 0
	}

	left := 0
	right := len(intList) - 1

	for right > left {

		// 判断 最左是不是最小
		if intList[left] < intList[left+1] {
			return left
		}

		// 判断最右是不是最小
		if intList[right] < intList[right-1] {
			return right
		}

		mid := (left + right) / 2
		// mid 比那一侧小 就再哪一测继续二分
		if intList[mid-1] < intList[mid] {
			right = mid - 1
			continue
		}
		if intList[mid+1] < intList[mid] {
			left = mid + 1
			continue
		}

		return mid

	}

	return -1
}

func randomList() []int {

	var l []int
	max := 100
	listSize := randomInt(98) + 2

	for i := 0; i < listSize; i++ {
		a := randomInt(max)
		for i > 0 && a == l[i-1] {
			a = randomInt(max)
		}

		l = append(l, a)

	}

	return l

}

func randomInt(max int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(a.Int64())
}

/**
对数器
*/
func comparator(l []int) []int {
	if len(l) == 1 {
		return []int{0}
	}
	var rightList []int
	for i := range l {
		// 最左
		if i == 0 {
			if l[i] < l[i+1] {
				rightList = append(rightList, i)
			}
			continue
		}

		// 最右
		if i == len(l)-1 {
			if l[i-1] > l[i] {
				rightList = append(rightList, i)
			}
			continue
		}

		if l[i-1] > l[i] && l[i+1] > l[i] {
			rightList = append(rightList, i)
		}

	}
	return rightList
}

func compare() {
	intList := randomList()

	index := getLessIndex(intList)

	rightIndexList := comparator(intList)

	for i := range rightIndexList {
		if index == rightIndexList[i] {
			//fmt.Println("right !")
			//fmt.Println(intList)
			//fmt.Println(index)
			//fmt.Println(rightIndexList)
			return
		}
	}

	fmt.Println("error !")
	fmt.Println(intList)
	fmt.Println(index)
	fmt.Println(rightIndexList)
}

func main() {
	for i := 0; i < 100000; i++ {
		compare()
	}

}
