package main

import (
	"fmt"
	"math"
)

func main() {
	getHighlyDivisibleTriangularNumber(500)
}

func getHighlyDivisibleTriangularNumber(over int) {
	var i int64 = 1
	var tn int64 = 1
	tf := 1
	for tf < over {
		i++
		tn = sumOfNInt(i)
		factors := getFactors(tn)
		tf = len(factors)
	}
	fmt.Println("Highly divisible triangular number:", tn)
}

func getFactors(num int64) []int64 {
	factors := []int64{1}
	for i := int64(2); i <= int64(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			factors = append(factors, i)
			factors = append(factors, num/i)
		}
	}
	factors = append(factors, num)
	return factors
}

func sumOfNInt(n int64) int64 {
	return (n*(n+1))/2
}
