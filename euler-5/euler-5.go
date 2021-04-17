package main

import (
	"fmt"
	"sort"
)

func getIntList(n int) []int {
	var xi []int
	for i := 1; i <= n; i++ {
		xi = append(xi, i)
	}
	return xi
}

func getLCM(nums []int) int64 {
	var lcm int64 = 1
	sort.Ints(nums) // sort all
	for i, n := range nums {
		if n > 1 {
			lcm = lcm * int64(n)
			nums[i] = 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%n == 0 {
					nums[j] = nums[j] / n
				}
			}
		}
	}
	return lcm
}

func main() {
	n := 20
	list := getIntList(n)
	fmt.Printf("LCM of first : %d numbers is %v", n, getLCM(list))
}
