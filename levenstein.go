package main

/**
 * Levenshtein distance algorithm
 *
 * @see https://en.wikipedia.org/wiki/Levenshtein_distance
 *
 * @author jan.cajthaml
 */
func LevensteinDistance(a, b []byte) int {
	if len(a) == 0 {
		return len(b)
	} else if len(b) == 0 {
		return len(a)
	} else if len(a) > len(b) {
		return distance(a, b)
	} else {
		return distance(b, a)
	}
}

func distance(a, b []byte) int {
	var (
		prev int
		i1   int   = 0
		l1   int   = len(b) + 1
		i2   int   = 0
		l2   int   = len(a)
		i3   int   = 0
		l3   int   = len(b)
		f    []int = make([]int, len(b)+1)
	)

s1:
	f[i1] = i1
	i1++
	if i1 < l1 {
		goto s1
	}
outer:
	prev = f[0]
	f[0]++
inner:
	if f[i3+1] <= f[i3] && b[i3] != a[i2] && f[i3+1] > prev {
		prev, f[i3+1] = f[i3+1], prev+1
	} else if f[i3+1] <= f[i3] && b[i3] != a[i2] {
		prev, f[i3+1] = f[i3+1], f[i3+1]+1
	} else if f[i3+1] <= f[i3] && f[i3+1]+1 > prev {
		prev, f[i3+1] = f[i3+1], prev
	} else if f[i3+1] <= f[i3] {
		prev, f[i3+1] = f[i3+1], f[i3+1]+1
	} else if b[i3] != a[i2] && f[i3] > prev {
		prev, f[i3+1] = f[i3+1], prev+1
	} else if b[i3] != a[i2] {
		prev, f[i3+1] = f[i3+1], f[i3]+1
	} else if f[i3]+1 > prev {
		prev, f[i3+1] = f[i3+1], prev
	} else {
		prev, f[i3+1] = f[i3+1], f[i3]+1
	}

	i3++
	if i3 != l3 {
		goto inner
	}
	i3 = 0
	i2++
	if i2 != l2 {
		goto outer
	}

	return f[len(f)-1]
}
