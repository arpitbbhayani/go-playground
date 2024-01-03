package main

import (
	"testing"
)

type BigStruct16 struct{ buf [1 << 16]byte }

var obj BigStruct16 = BigStruct16{}

func fooPBV(obj BigStruct16) { obj.buf[0] = 0 }

func fooPBP(obj *BigStruct16) { obj.buf[0] = 0 }

func BenchmarkPBP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fooPBP(&obj)
	}
}

func BenchmarkPBV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fooPBV(obj)
	}
}
