package main

import (
	"testing"
)

type BigStruct struct{ buf [1 << 16]byte }

var obj BigStruct = BigStruct{}

func fooPBV(obj BigStruct) { obj.buf[0] = 0 }

func fooPBP(obj *BigStruct) { obj.buf[0] = 0 }

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
