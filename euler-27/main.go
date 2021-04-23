package main

import (
	"fmt"
	"math/big"
)

func solveQuadratic(a int, b int, n int) int {
	res := n * n + a * n + b
	return res
}

func main() {
	maxA := 0
	maxB := 0
	maxN := 0
	for a := -999; a < 1000; a++ {
		for b := 2; b <= 1000; b++ {
			n := 0
			for {
				// check consecutive primes
				ans := solveQuadratic(a, b, n)
				num := big.NewInt(int64(ans))
				if ans < 0 || !num.ProbablyPrime(1) {
					break
				}
				n++
			}
			if n > maxN {
				maxN = n
				maxA = a
				maxB = b
			}
		}
	}
	fmt.Printf("coefficients are %v and %v with prime for consective values 0 <= n <= %v\n", maxA, maxB, maxN)
	fmt.Printf("Product of coefficients %v\n", maxA*maxB)
}
