package main

import "testing"

var n = 20

func BenchmarkGetAllPaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAllPaths(0, 0, n+1, 0);
	}
}

func BenchmarkGetAllPathsByCombinatorics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAllPathsByCombinatorics(n)
	}
}
