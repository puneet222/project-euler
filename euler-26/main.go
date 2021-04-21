package main

import "fmt"

func getDecimalLength(d int) int {
	num := 1
	res := make([]int, 0)
	numerators := map[int]bool{
		1: true,
		0: true,
	}
	for {
		if num < d {
			num = num*10
		}
		q := num/d
		res = append(res, q)
		num = num%d
		if _, ok := numerators[num]; ok {
			break
		} else {
			numerators[num] = true
		}
	}
	return len(res)
}

func main() {
	maxD := 2
	maxLen := 1
	for i := 2; i < 1000; i++{
		dl := getDecimalLength(i)
		if dl > maxLen {
			maxLen = dl
			maxD = i
		}
	}
	fmt.Printf("Maximum d is %v with decimal length: %v\n", maxD, maxLen)
}
