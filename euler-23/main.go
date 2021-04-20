package main

import (
	"fmt"
	"math"
)

func getSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func getFactors(n int)[]int {
	limit := int(math.Sqrt(float64(n)))
	factors := []int{1}
	for i := 2; i <= limit; i++ {
		if n % i == 0 {
			if i*i ==  n {
				// sqrt
				factors = append(factors, i)
			} else {
				factors = append(factors, i)
				factors = append(factors, n/i)
			}
		}
	}
	return factors
}

func isAbundant(n int) bool {
	fac := getFactors(n)
	sf := getSum(fac)
	return sf > n
}

var sum2Abundant = make(map[int]bool)

func main() {
	limit := 28123
	abundant := make([]int, 0)
	// get all abundant numbers till limit
	for i := 1; i <= limit; i++ {
		if isAbundant(i) {
			abundant = append(abundant, i)
		}
	}
	// get all possible sums of abundant numbers and
	// put all the values in map for faster lookup
	for i := 0; i < len(abundant); i++ {
		for j := i; j < len(abundant); j++ {
			s2a := abundant[i] + abundant[j]
			if _, ok := sum2Abundant[s2a]; !ok {
				sum2Abundant[s2a] = true
			}
		}
	}
	res := 0
	// sum of all numbers which are not present in
	// sums of abundant numbers map
	for i := 1; i <= limit; i++ {
		if _, ok := sum2Abundant[i]; !ok {
			// cannot be written as sum of 2 abundant integers
			res += i
		}
	}

	fmt.Println("Sum of all the positive integers which cannot be written as the sum of two abundant numbers:", res)
}
