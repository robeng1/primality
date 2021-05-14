package prime

import (
	"math/big"
	"math/rand"
)

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

// modExpGoBigIntegerExp calculates modular exponentiation using native Exp method from math/big package.
func modExpGoBigIntegerExp(base, exponent, modulus int64) int64 {
	return new(big.Int).Exp(big.NewInt(base), big.NewInt(exponent), big.NewInt(modulus)).Int64()
}




// This function is called for all k trials. It returns
// false if n is composite and returns true if n is
// probably prime.
// d is an odd number such that  d*2<sup>r</sup> = n-1
// for some r >= 1
func millerTest(d, n int64) bool {
	// Pick a random number in [2..n-2]
	// Corner cases make sure that n > 4
	a := 2 + int64(rand.Int())%(n-4)

	// Compute a^d % n
	x := modExpGoBigIntegerExp(a, d, n)

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
