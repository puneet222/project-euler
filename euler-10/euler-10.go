package main

import "fmt"

func primeSummation(limit int) int{
	sum := 0
	sieve := make([]bool, limit, limit)
	for i := 2; i < limit; i++ {
		if !sieve[i] {
			sum += i
			for j := i; j < limit; j = j + i {
				sieve[j] = true
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Prime Summation below 10: ", primeSummation(10))
	l := 2000000
	fmt.Printf("Prime Summation below %v: %v", l, primeSummation(l))
}
