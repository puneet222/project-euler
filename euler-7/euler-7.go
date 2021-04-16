package main

import "fmt"

func getNPrimes(MaxNum int, n int) []int {
	// creating sieve of length MaxNum
	sieve := make([]bool, MaxNum, MaxNum)
	for i, _ := range sieve {
		sieve[i] = true
	}
	// slice where we store all the primes
	primes := make([]int, 0, n)
	count := 0
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			count++
			primes = append(primes, i)
			if count == n {
				break
			}
			// filling the sieve of multiple of i with false
			// as it is not primes because divisible by i
			for j := i; j < len(sieve); j = j + i {
				sieve[j] = false
			}
		}
	}
	return primes
}

func main() {
	MaxNum := 150000
	primes := getNPrimes(MaxNum, 10001)
	fmt.Println("6th Prime: ", primes[5])
	fmt.Println("10001st Prime: ", primes[10000])
}
