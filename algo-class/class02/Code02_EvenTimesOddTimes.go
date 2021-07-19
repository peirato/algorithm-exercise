package main

// 数组中 只有一种数 出现奇数次
func findOddTimesNum1(arr []int)int {
	eor := 0
	for i := range arr {
		eor ^= arr[i]
	}
	return eor
}

// 数组中 只有两种数 出现奇数次
func findOddTimesNum2(arr []int) (int, int) {
	// 设数为 a,b
	eor := 0
	for i := range arr {
		eor ^= arr[i]
	}

	// eor = a ^ b
	// 取最右侧为1的一位
	right :=findRight1(eor)

	// 只和某一位为1的计算，结果一定是 a 或者 b
	eor2:=0
	for i := range arr {
		if right & arr[i] !=0{
			eor2 ^=arr[i]
		}
	}

	a := eor2
	b := eor ^ eor2

	return a,b
}
// 找到最右的为1的一位
func findRight1(i int)int{
	return i & (1 + ^i)
}


func main() {
	// 11  10
	findRight1(4)
	println(findOddTimesNum1([]int{3,1,1,4,4}))
	println(findOddTimesNum2([]int{3,1,1,5,4,4}))
}

