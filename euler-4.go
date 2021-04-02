package main

import "fmt"

func reverse(num uint32) uint32 {
	r := uint32(0)
	for num > 0 {
		r = r * 10 + num % 10
		num = num / 10
	}
	return r
}

func isPalindrome(num uint32, rev uint32) bool {
	return num == rev
}

func main() {
	p := uint32(0)
	for i := 100; i < 1000; i++ {
		for j := i+1; j < 1000; j++ {
			num := uint32(i*j)
			rev := reverse(num)
			if isPalindrome(num, rev) {
				if num > p {
					p = num
				}
			}
		}
	}
	fmt.Println("Max Palindrome: ", p)
}
