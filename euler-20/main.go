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

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	limit := 100
	fac := []int{1}
	for i := 2; i <= limit; i++ {
		num := createSlice(i)
		fac = multiplyInts(fac, num)
	}
	fmt.Println(reverse(fac))
	fmt.Printf("Sum of digits of %v! = %v\n", limit, sum(fac))
}
