package main

import (
	"fmt"
	"math"
)

func sumOfN(n float64) float64 {
	return (n*(n+1))/2
}

func sumOfSqN(n float64) float64 {
	return (n*(n+1)*((2*n)+1))/6
}

func main() {
	var n float64 = 100
	sumN := sumOfN(n)
	sumSqN := sumOfSqN(n)
	sumNSq := math.Pow(sumN, 2)
	val := fmt.Sprintf("%f", sumNSq - sumSqN)
	fmt.Printf("Sum square difference of first %f numbers is %v", n, val)
}
