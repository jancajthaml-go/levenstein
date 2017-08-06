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

func TestDistanceOnEmptyStrings(t *testing.T) {
	expectDistance := func(a string, b string, c int) {
		d := LevensteinDistance([]byte(a), []byte(b))
		if d != c {
			t.Errorf("a: `" + a + "`, b: `" + b + "` expected " + string(c+48) + " , got " + string(d+48))
		}
	}

	expectDistance("", "", 0)
	expectDistance("a", "a", 0)
	expectDistance("abc", "abc", 0)

	expectDistance("a", "", 1)
	expectDistance("", "a", 1)
	expectDistance("abc", "", 3)
	expectDistance("", "abc", 3)

	// inserts only
	expectDistance("a", "ab", 1)
	expectDistance("b", "ba", 1)
	expectDistance("ac", "abc", 1)
	expectDistance("abcdefg", "xabxcdxxefxgx", 6)

	// deletes only
	expectDistance("ab", "a", 1)
	expectDistance("ab", "b", 1)
	expectDistance("abc", "ac", 1)
	expectDistance("xabxcdxxefxgx", "abcdefg", 6)

	// swaps only
	expectDistance("a", "b", 1)
	expectDistance("ab", "ac", 1)
	expectDistance("ac", "bc", 1)
	expectDistance("abc", "axc", 1)
	expectDistance("xabxcdxxefxgx", "1ab2cd34ef5g6", 6)

	// combination of insert + delete + swap
	expectDistance("example", "samples", 3)
	expectDistance("sturgeon", "urgently", 6)
	expectDistance("levenshtein", "frankenstein", 6)
	expectDistance("distance", "difference", 5)
}
