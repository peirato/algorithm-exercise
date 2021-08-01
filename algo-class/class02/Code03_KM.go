package main

import (
	"fmt"
	"utils"
)

// 一个数组，一种数出现了k次，其他都出现了m次，m>1，m>k,找到k

func onlyKTimes(arr []int, k int, m int) int {

	var t [32]int

	// 初始化数组
	for i := range t {
		t[i] = 0
	}

	for i := range arr {
		for j := range t {
			// 判断 数组中的数 下标为j的位是不是1
			// 让数右移 j-1 次 & 1
			if arr[i]>>(j)&1 == 1 {
				t[j]++
			}
			if arr[1] == 0 {
				break
			}
		}
	}

	x := 0

	for i := range arr {

		if t[i]%m != 0 {
			// % m 后不为零 表示这一位为1
			x |= 1 << i
		}
	}

	return x
}

const maxNum = 100
const maxTimes = 100
const numCountMax = 20

func genKM() (int, int) {

	// k >1 && k < 98
	k := utils.RandomInt(maxTimes-2) + 1

	var m int

	for {
		m = utils.RandomInt(maxTimes)
		if m > k {
			return k, m
		}
	}
}

func genArray(k int, m int) []int {

	// 一共有多少个数
	numCount := utils.RandomIntBetween(2, 20)

	// 生成这些数
	numMap := make(map[int]int)
	for i := 0; i < numCount; i++ {
		for {
			num := utils.RandomInt(maxNum)
			// 是否存在
			_, ok := numMap[num]
			if !ok {
				numMap[num] = 1
				break
			}
		}
	}

	var li []int

	for key := range numMap {

		var times int
		if li == nil {
			times = k
		} else {
			times = m
		}

		// 依次向数组里添加数
		for i := 0; i < times; i++ {
			li = append(li, key)
		}
	}

	utils.Shuffle(li)

	return li

}

func comparators(l []int) int {
	numMap := make(map[int]int)
	for i := range l {
		num := l[i]
		_, ok := numMap[num]
		if ok {
			numMap[num] = numMap[num] + 1
		} else {
			numMap[num] = 1
		}
	}

	minTime := 0
	for key := range numMap {
		value := numMap[key]
		if minTime > value {
			minTime = value
		}
	}
	return minTime
}

func compare() {
	k, m := genKM()
	l := genArray(k, m)
	fmt.Println(l)
	result := onlyKTimes(l, k, m)
	fmt.Println(result)
	result2 := comparators(l)
	fmt.Println(result2)

}

func main() {
	for i := 0; i < 10; i++ {
		compare()
	}

}
