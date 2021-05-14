package prime

import "math/rand"

// Iterative Function to calculate (a^n)%p in O(logy)
func power(a uint64, n uint64, p uint64) uint64 {
	res := uint64(1) // Initialize result
	a = a % p       // Update 'a' if 'a' >= p

	for n > 0 {
		// If n is odd, multiply 'a' with result
		if n&1 > 0 {
			res = (res * a) % p
		}

		// n must be even now
		n = n >> 1 // n = n/2
		a = (a * a) % p
	}
	return res
}

func gcd(a, b uint64) uint64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}


// This function is called for all k trials. It returns
// false if n is composite and returns true if n is
// probably prime.
// d is an odd number such that  d*2<sup>r</sup> = n-1
// for some r >= 1
func millerTest(d, n uint64) bool {
	// Pick a random number in [2..n-2]
	// Corner cases make sure that n > 4
	a := 2 + uint64(rand.Int())%(n-4)

	// Compute a^d % n
	x := power(a, d, n)

	if x == 1 || x == n-1 {
		return true
	}

	// Keep squaring x while one of the following doesn't
	// happen
	// (i)   d does not reach n-1
	// (ii)  (x^2) % n is not 1
	// (iii) (x^2) % n is not n-1
	for d != n-1 {
		x = (x * x) % n
		d *= 2

		if x == 1 {
			return false
		}
		if x == n-1 {
			return true
		}
	}

	// Return composite
	return false
}
