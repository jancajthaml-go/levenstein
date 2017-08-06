package main

func LevensteinDistance(a, b []byte) int {
	if len(a) > len(b) {
		return distance(a, b)
	} else {
		return distance(b, a)
	}
}

func distance(a, b []byte) int {
	var (
		f_prev int
		i1     int   = 0
		l1     int   = len(b) + 1
		i2     int   = 0
		l2     int   = len(a)
		i3     int   = 0
		l3     int   = len(b)
		f      []int = make([]int, len(b)+1)
	)

s1:
	f[i1] = i1
	i1++
	if i1 != l1 {
		goto s1
	}
s2:
	f_prev = f[0]
	f[0]++
s3:
	cb := b[i3]
	i3++

	if f[i3] <= f[i3-1] && cb != a[i2] && f[i3] > f_prev {
		f_prev, f[i3] = f[i3], f_prev+1
	} else if f[i3] <= f[i3-1] && cb != a[i2] {
		f_prev, f[i3] = f[i3], f[i3]+1
	} else if f[i3] <= f[i3-1] && f[i3]+1 > f_prev {
		f_prev, f[i3] = f[i3], f_prev
	} else if f[i3] <= f[i3-1] {
		f_prev, f[i3] = f[i3], f[i3]+1
	} else if cb != a[i2] && f[i3-1] > f_prev {
		f_prev, f[i3] = f[i3], f_prev+1
	} else if cb != a[i2] {
		f_prev, f[i3] = f[i3], f[i3-1]+1
	} else if f[i3-1]+1 > f_prev {
		f_prev, f[i3] = f[i3], f_prev
	} else {
		f_prev, f[i3] = f[i3], f[i3-1]+1
	}

	if i3 != l3 {
		goto s3
	}
	i3 = 0
	i2++
	if i2 != l2 {
		goto s2
	}

	return f[len(f)-1]
}
