package main

import (
	"fmt"
	"math"
)

func checkPythagoreanTriplet(a int, b int, c int) bool {
	asq := math.Pow(float64(a), 2)
	bsq := math.Pow(float64(b), 2)
	csq := math.Pow(float64(c), 2)
	return asq + bsq == csq
}

func findPythagoreanTriplet() (int, int, int) {
	for a := 1; a < 500; a++ {
		for b := 1; b < 500; b++ {
			c := 1000 - (a + b)
			if checkPythagoreanTriplet(a, b, c) {
				return a,b,c
			}
		}
	}
	return 0,0,0
}

func main() {
	a,b,c := findPythagoreanTriplet()
	fmt.Println("Pythagorean Triplet: ", a, b, c)
	fmt.Println("Pythagorean Triplet Product: ", a*b*c)
}
