package main

import "fmt"

func add(r1 []int, r2 []int) []int {
	var sol []int
	carry := 0
	i := 0
	j := 0
	for i < len(r1) || j < len(r2) {
		sum := 0
		if i >= len(r1) {
			sum += r2[j]
		} else if j >= len(r2) {
			sum += r1[i]
		} else {
			sum += r1[i] + r2[j]
		}
		sum += carry
		carry = sum/10
		sum = sum%10
		sol = append(sol, sum)
		i++
		j++
	}
	if carry > 0 {
		sol = append(sol, carry)
	}
	return sol
}

func main() {
	size := 1000
	var fac []int
	prev1 := []int{1}
	curr := []int{1}
	index := 2
	for len(fac) < size {
		// the number restured by add function is reverse
		fac = add(prev1, curr)
		index++
		prev1 = curr
		curr = fac
	}
	fmt.Println(index)
}
