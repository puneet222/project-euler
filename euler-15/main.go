package main

import (
	"fmt"
	"math"
	"strconv"
)

var m = make(map[string]int)

func GetAllPaths(r int, c int, n int, paths int) int {
	if r == n || c == n {
		return 0
	}
	if r == n-1 && c == n-1 {
		return 1
	}
	// check if we already visited this path
	currentPath := strconv.Itoa(r) + "-" + strconv.Itoa(c)
	if val, ok := m[currentPath]; ok {
		return val
	}
	// move right
	rPaths := GetAllPaths(r, c+1, n, paths)
	// move down
	dPaths := GetAllPaths(r+1, c, n, paths)
	p := rPaths + dPaths
	// save total paths from current path
	m[currentPath] = p
	return rPaths + dPaths
}

func GetAllPathsByCombinatorics(n int) int {
	var result float64 = 1
	for i := 1; i <= n; i++ {
		result *= float64(n+i)/float64(i)
	}

	return int(math.Ceil(result))
}

func main() {
	n := 20
	paths := GetAllPaths(0, 0, n+1, 0)
	cPaths := GetAllPathsByCombinatorics(n)
	fmt.Printf("total paths for %v X %v matrix is %v\n", n, n, paths)
	fmt.Printf("total paths by combinatorics for %v X %v matrix is %v\n", n, n, cPaths)
}
