package main

import "fmt"

func Reverse(num []int) []int {
	for i, j := 0, len(num)-1; i < j; i,j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
	return num
}

func Sum(num []int) int {
	result := 0
	for _, v := range num {
		result += v
	}
	return result
}

func MultiplyBy2(num []int) []int {
	// num will be reverse
	var res []int
	carry := 0
	for _, v := range num {
		m := v*2 + carry
		carry = m/10
		res = append(res, m%10)
	}
	if carry > 0 {
		res = append(res, carry)
	}
	return res
}

func main() {
	num := []int{1}
	pow := 1000
	for i := 0; i < pow; i++ {
		num = MultiplyBy2(num)
	}
	fmt.Println("Value of 2^1000 = ", Reverse(num))
	fmt.Println("Sum of digits: ", Sum(num))
}
