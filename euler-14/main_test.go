package main

import "testing"

func BenchmarkGetMaxChainTerms(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMaxChainTerms(999999)
	}
}

func BenchmarkGetMaxChainTermsWithCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMaxChainTermsWithCache(999999)
	}
}
