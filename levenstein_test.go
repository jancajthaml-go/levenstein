package main

import "testing"

func BenchmarkLevensteinParallel(b *testing.B) {
	left := []byte("xxxx")
	right := []byte("xxxy")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			LevensteinDistance(left, right)
		}
	})
}

func BenchmarkLevensteinSerial(b *testing.B) {
	left := []byte("xxxx")
	right := []byte("xxxy")

	for n := 0; n < b.N; n++ {
		LevensteinDistance(left, right)
	}
}
