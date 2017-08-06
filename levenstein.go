package main

func LevensteinDistance(a, b []byte) int {
	var (
		f_curr int
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
	if i1 == l1 {
		goto s2
	}
	goto s1
s2:
	ca := a[i2]
	f_prev = f[0]
	f[0]++
s3:
	cb := b[i3]
	i3++
	if f[i3] <= f[i3-1] {
		f_curr = f[i3] + 1 // delete
	} else {
		f_curr = f[i3-1] + 1 // insert
	}
	if cb != ca {
		if f_curr > f_prev+1 { // change
			f_curr = f_prev + 1
		}
	} else {
		if f_curr > f_prev { // matched
			f_curr = f_prev
		}
	}
	f_prev, f[i3] = f[i3], f_curr
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
