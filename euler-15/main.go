package main

import (
	"fmt"
	"strconv"
)

var m = make(map[string]int)

func getAllPaths(r int, c int, n int, paths int) int {
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
	rPaths := getAllPaths(r, c+1, n, paths)
	// move down
	dPaths := getAllPaths(r+1, c, n, paths)
	p := rPaths + dPaths
	// save total paths from current path
	m[currentPath] = p
	return rPaths + dPaths
}

func main() {
	n := 20
	paths := getAllPaths(0, 0, n+1, 0);
	fmt.Printf("total paths for %v X %v matrix is %v\n", n, n, paths)
}
