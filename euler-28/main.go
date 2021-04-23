package main

import "fmt"

func getDiagonalSum(size int) int {
	sum := 1
	level := 1
	for i := 2; i <= size; i++ {
		if i % 2 == 0 {
			// even
			sum += i*i + 1
			sum += i*i - level
		} else {
			// odd
			sum += i*i
			sum += i*i - (level + 1)
			level += 2
		}
	}
	return sum
}

func main() {
	size := 1001
	sum := getDiagonalSum(size)
	fmt.Printf("Number spiral diagonal sum of matrix of size %v -> %v", size, sum)
}
