package main

import "fmt"

// 不用第三个变量 交换两个数

func swap(a int, b int) (int, int) {

	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	fmt.Println(swap(3,5))
}
