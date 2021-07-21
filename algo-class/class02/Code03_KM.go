package main

import "fmt"

// 一个数组，一种数出现了k次，其他都出现了m次，m>1，m>k,找到k

func onlyKTimes(arr []int, k int, m int) {

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

	fmt.Println(t)

	x := 0

	for i := range arr {
		if t[i]%m != 0 {
			// % m 后不为零 表示这一位为1
			x |= 1 << i
		}
	}

	fmt.Println(x)
}

// TODO 实现一个对数器

func main() {
	onlyKTimes([]int{3, 1, 1, 4, 4}, 1, 2)

}
