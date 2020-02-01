package algos

func ParseInteger(s string) int {
	N := len(s)

	var v int

	for _, c := range []byte(s) {
		var d int
		if c == '0' {
			d = 0
		}
		if c == '1' {
			d = 1
		}
		N -= 1
		// left shift
		// 100101 = [(1)*2^5]+[(0)*2^4]+[(0)*2^3]+[(1)*2^2]+[(0)*2^1]+[(1)*2^0]
		v += (d << N)
	}

	return v
}

func Parity(s string) int {
	// test the evenness and oddness of a number
	v := ParseInteger(s)
	n := v
	var steps int

	for i := 0; i < n; i++ {
		steps += 1

		if v&1 == 0 {
			v >>= 1 // Divide by 2
		} else {
			v -= 1
		}
		if v == 0 {
			break
		}
	}

	return steps
}

func IdenticalIndices(a []int) int {
	/*
	   finding identical pair of indices
	   Parameters
	   ----------
	   a: list
	       a is a list of N numbers of integers

	   Returns
	   -------
	   out: int
	       out is the count of the number of identical pairs

	   Notes
	   -----
	   A pair of indices (P, Q) is identical if 0<=P<Q<N and
	   A[P] == A[Q]; A being the array of integers
	*/

	var current int
	next := 1
	var count int

	l := len(a)

	for {
		if a[current] == a[next] {
			count += 1
			next += 1
		} else {
			next += 1
		}
		if next == l {
			current += 1
			next = current + 1
		}
		if current > l-2 {
			break
		}
	}
	return count
}
