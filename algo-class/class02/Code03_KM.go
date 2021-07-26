package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

func comparators() int {
	return 0
}

const maxNum = 100
const maxTimes = 100
const numCountMax = 20

func genKM() (int, int) {

	// k >1 && k < 98
	k := number_utils.randomInt(maxTimes-2) + 1

	var m int

	for {
		m = randomInt(maxTimes)
		if m > k {
			return k, m
		}
	}
}

func genArray() []int {
	k, m := genKM()
	kNum := randomInt(maxTimes)

}

func main() {
	onlyKTimes([]int{3, 1, 1, 4, 4}, 1, 2)
}
