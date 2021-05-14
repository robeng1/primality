// This package contains functions for checking the primality of a number
// and finding the highest prime number lower than a given input
package prime

import (
	"math"
	"math/rand"
)

//OptimizedRepeatedDivisonPrimality checks the primality of a number using repeated division
func OptimizedRepeatedDivisonPrimality(x uint64) bool {
	// imediately return the input is 2
	if x == 2 {
		return true
	}

	// 2 is the only even prime number, so we immediately exit when x is even
	if x%2 == 0 {
		return false
	}

	var i uint64 // counter

	// Since 2 is the only even prime, once we verify that x isn’t even we only need try the odd
	// numbers as candidate factors, hence i starting from 3, 5, ....
	i = 3

	// We terminate the loop at √x because if x is composite but has a smallest non-trivial prime factor p which is greater than √x.
	// Then x/p must also divide x, and must be larger than p, or else we would have seen
	// it earlier. But the product of two numbers greater than √n must be larger than n, which makes it a
	// contradiction

	// Run math.Sqrt() with an extra iteration since it is a numerical function with imperfect precision
	// It is being run outside the main loop to avoid calling math.Sqrt() on each iteration
	root := math.Sqrt(float64(x)) + 1

	for float64(i) < root {
		if x%i == 0 {
			return false
		} else {
			// we only need try the odd numbers as candidate factors
			i = i + 2
		}
	}
	return true
}

//If n is a prime number, then for every a, 1 < a < n-1,
//a^(n-1) ≡ 1 (mod n)
// OR
//a^(n-1) % n = 1
// This test is also probalistic based of Fermat's Little Theorem
// More and detailed proof https://en.wikipedia.org/wiki/Proofs_of_Fermat's_little_theorem
func FermatPrimality(n uint64) bool {
	// we set k to 7 iterations
	k := 7
	// Corner cases
	if n <= 1 || n == 4 {
		return false
	}
	if n <= 3 {
		return true
	}

	// Try k times
	for k > 0 {
		// Pick a random number in [2..n-2]
		// Above corner cases make sure that n > 4
		a := 2 + uint64(rand.Int())%(n-4)

		// Checking if a and n are co-prime
		if gcd(n, a) != 1 {
			return false
		}

		// Fermat's little theorem
		if power(a, n-1, n) != 1 {
			return false
		}

		k--
	}

	return true

}

// It returns false if n is composite and returns true if n
// is probably prime.  k is an input parameter that determines
// accuracy level. Higher value of k indicates more accuracy.
// This test is probabilistic
// References https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
func MillerRabinPrimality(x uint64) bool {
	// best known tests of this type involve 3 rounds of the Miller-Rabin test
	// for 32-bit integers and 7 rounds for 64-bit integer
	// can be checked with benchmarking
	k := 40

	// Corner cases
	if x <= 1 || x == 4 {
		return false
	}
	if x <= 3 {
		return true
	}
	// Find r such that n = 2^d * r + 1 for some r >= 1
	d := x - 1
	for d%2 == 0 {
		d /= 2
	}

	// Iterate given nber of 'k' times
	for i := 0; i < k; i++ {
		if !millerTest(d, x) {
			return false
		}

	}

	return true
}

// Finds the highest prime number lower the input number x
func FindPrimeToLeft(x uint64) uint64 {
	// since the prime number we're looking for has to be lower than
	// the input number we start checking from x-1
	for i := x - 1; i > 1; i-- {
		if OptimizedRepeatedDivisonPrimality(i) {
			return i
		}
	}
	// return -1 if no prime number is found
	return 0
}
