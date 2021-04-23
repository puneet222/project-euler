package main

import "fmt"

func reverse(n []int) []int {
	i := 0
	j := len(n) - 1
	for i < j {
		n[i], n[j] = n[j], n[i]
		i++
		j--
	}
	return n
}

// return reverse slice
func createSlice(n int) []int {
	res := make([]int,0)
	for n > 0 {
		res = append(res, n%10)
		n = n/10
	}
	return res
}

func sumN(nums [][]int) []int {
	if len(nums) == 1 {
		return nums[0]
	}
	//res := make([]int, 0)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		// sum of prev and curr
		carry, p, c := 0,0,0
		temp := make([]int, 0)
		for p < len(prev) && c < len(curr) {
			sum := prev[p] + curr[c] + carry
			temp = append(temp, sum%10)
			carry = sum/10
			p++
			c++
		}
		// remaining
		for p < len(prev) {
			sum := prev[p] + carry
			temp = append(temp, sum%10)
			carry = sum/10
			p++
		}
		for c < len(curr) {
			sum := curr[c] + carry
			temp = append(temp, sum%10)
			carry = sum/10
			c++
		}
		// check carry
		if carry > 0 {
			temp = append(temp, carry)
		}
		prev = temp // assigning
	}
	return prev
}

func product(num []int, n int) []int {
	res := make([]int, 0)
	carry := 0
	for i := 0; i < len(num); i++ {
		pr := n*num[i] + carry
		res = append(res, pr%10)
		carry = pr/10
	}
	if carry > 0 {
		res = append(res, carry)
	}
	return res
}

func multiplyInts(n1 []int, n2 []int) []int {
	// n1 and n2 are reversed
	az := 0 // appending zeros
	var nums [][]int
	for i := 0; i < len(n2); i++ {
		pr := product(n1, n2[i])
		for z := 0; z < az; z++ {
			pr = append([]int{0}, pr...)
		}
		az++ // append extra zero for next level
		nums = append(nums, pr)
	}
	return sumN(nums)
}

func main() {
	limit := 1000
	list := make([][]int, limit, limit)
	for i := 1; i <= limit; i++ {
		pr := createSlice(i)
		for j := 1; j < i; j++ {
			pr = multiplyInts(pr, createSlice(i))
		}
		list[i-1] = pr
	}
	sum := sumN(list)
	fmt.Println(reverse(sum))
}

// 7 8 1 2 7 8 1 9 1 1 0 8 4 6 7 0 0
// 9110846700
