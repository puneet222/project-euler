package main

import (
	"fmt"
	"math"
)

func getPrimeFactors(num int64) []int64 {
	var pfs []int64

	// all 2 factors
	for num % 2 == 0 {
		pfs = append(pfs, 2)
		num = num / 2
	}

	// all odd factors
	var i int64
	for i = 3; i <= int64(math.Sqrt(float64(num))); i = i + 2 {
		for num % i == 0 {
			pfs = append(pfs, i)
			num = num / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if num > 2 {
		pfs = append(pfs, num)
	}

	return pfs
}

func main() {
	fmt.Println("Factors: ", getPrimeFactors(600851475143))
}
