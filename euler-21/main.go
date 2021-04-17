package main

import (
	"fmt"
	"math"
)

var memo = make(map[int]bool)

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
			factors = append(factors, i)
			factors = append(factors, n/i)
		}
	}
	return factors
}

func isAmicable(n int) (bool, int, int) {
	fac1 := getFactors(n) // n = 220
	sum1 := getSum(fac1)  // sum1 = 284
	fac2 := getFactors(sum1)
	sum2 := getSum(fac2) // sum2 = 220
	if n == sum2 && sum1 != sum2 {
		return true, n, sum1
	} else {
		return false, 0, 0
	}
}

func main() {
	limit := 10000
	amicable := make([]int, 0)
	for i := 1; i <= limit; i++ {
		if _, ok := memo[i]; !ok {
			// not present
			if ok, n1, n2 := isAmicable(i); ok {
				if n1 == n2 {
					amicable = append(amicable, n1)
				} else {
					amicable = append(amicable, n1, n2)
				}
				memo[n1] = true
				memo[n2] = true
			}
		}
	}
	fmt.Printf("Amicable Numbers under %v: %v\n", limit, amicable)
	fmt.Printf("Sum of Amicable numbers under %v : %v\n", limit, getSum(amicable))
	//isAmicable(220)
}
