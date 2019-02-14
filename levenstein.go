//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Levenshtein_distance
//
package levenstein

// Distance returns levenstein distance between two strings
func Distance(left string, right string) int {

	if len(left) == 0 {
		return len(right)
	}

	if len(right) == 0 {
		return len(left)
	}

	var a []byte
	var b []byte
	if len(left) > len(right) {
		a = []byte(left)
		b = []byte(right)
	} else {
		a = []byte(right)
		b = []byte(left)
	}

	var (
		prev int
		i1   int
		l1   int = len(b) + 1
		i2   int
		l2   int = len(a)
		i3   int
		l3   int   = len(b)
		f    []int = make([]int, len(b)+1)
	)

pre:
	f[i1] = i1
	i1++
	if i1 < l1 {
		goto pre
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
