package main

import "fmt"

var m  = make(map[int]int)

func getTotalChainTermsWithCache(start int) int {
	count := 1
	if val, ok := m[start]; ok {
		return val
	}
	for start > 1 {
		if start % 2 == 0 {
			// even
			start = start/2
		} else {
			// odd
			start = (3*start + 1)/2
		}
		count++
	}
	m[start] = count
	return count
}

func getTotalChainTerms(start int) int {
	count := 1
	for start > 1 {
		if start % 2 == 0 {
			// even
			start = start/2
		} else {
			// odd
			start = (3*start + 1)/2
		}
		count++
	}
	return count
}

func GetMaxChainTerms(limit int) (int, int){
	maxCount := 1
	maxNum := 1
	for i := 1; i < limit; i++ {
		count := getTotalChainTerms(i)
		//fmt.Println("num: ", i, " count: ", count)
		if count > maxCount {
			maxCount = count
			maxNum = i
		}
	}
	return maxNum, maxCount
}

func GetMaxChainTermsWithCache(limit int) (int, int){
	maxCount := 1
	maxNum := 1
	for i := 1; i < limit; i++ {
		count := getTotalChainTermsWithCache(i)
		//fmt.Println("num: ", i, " count: ", count)
		if count > maxCount {
			maxCount = count
			maxNum = i
		}
	}
	return maxNum, maxCount
}

func main() {
	limit := 999999
	maxNum, maxCount := GetMaxChainTerms(limit)
	maxNum, maxCount = GetMaxChainTermsWithCache(limit)
	fmt.Printf("Number with the max chain terms is %v with total chains: %v\n", maxNum, maxCount)
}
