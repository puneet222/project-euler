package main

import (
	"fmt"
	"sort"
)

func covertToInt(digits []int) int {
	num := 0
	for _, d := range digits {
		num = num*10 + d
	}
	return num
}

func getAllPermutations(nums []int) [][]int {
	var perm [][]int
	if len(nums) == 1 {
		perm = append(perm, nums)
	} else {
		for i, n := range nums {
			cpy := make([]int, len(nums)) // copying slice
			copy(cpy, nums)
			p := getAllPermutations(append(cpy[0:i], cpy[i+1:]...))
			for ind := 0; ind < len(p); ind++ {
				p[ind] = append(p[ind], n)
			}
			perm = append(perm, p...)
		}
	}
	return perm
}

func getNumberPermutations(perm [][]int) []int {
	np := make([]int, 0, len(perm))
	for _, p := range perm {
		num := covertToInt(p)
		np = append(np, num)
	}
	return np
}

func main() {
	list := []int{0,1,2,3,4,5,6,7,8,9}
	perm := getAllPermutations(list)
	numPerm := getNumberPermutations(perm)
	sort.Ints(numPerm)
	fmt.Println(len(numPerm))
	fmt.Printf("Millionth lexicographic permutation of the digits %v -> %v\n", list, numPerm[999999])
}
