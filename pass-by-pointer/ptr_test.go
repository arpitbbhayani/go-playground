package main

import (
	"testing"
)

func BenchmarkPBP(b *testing.B) {
	for n := 0; n < b.N; n++ { fooPBP(&obj) }
}

func BenchmarkPBV(b *testing.B) {
	for n := 0; n < b.N; n++ { fooPBV(obj) }
}
