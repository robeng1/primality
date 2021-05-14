// This package contains functions for checking the primality of a number
package prime

import (
	"testing"
)

func TestOptimizedRepeatedDivisonPrimality(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Check 2", args: args{x: 2}, want: true},
		{name: "Check 3", args: args{x: 3}, want: true},
		{name: "Check 4", args: args{x: 4}, want: false},
		{name: "Check 5", args: args{x: 5}, want: true},
		{name: "Check 6", args: args{x: 6}, want: false},
		{name: "Check 7", args: args{x: 7}, want: true},
		{name: "Check 2,147,483,647", args: args{2147483647}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptimizedRepeatedDivisonPrimality(tt.args.x); got != tt.want {
				t.Errorf("OptimizedRepeatedDivisonPrimality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOptimizedRepeatedDivisonPrimality(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptimizedRepeatedDivisonPrimality(2147483647)
	}
}

func TestFermatPrimality(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Check 2", args: args{n: 2}, want: true},
		{name: "Check 3", args: args{n: 3}, want: true},
		{name: "Check 4", args: args{n: 4}, want: false},
		{name: "Check 5", args: args{n: 5}, want: true},
		{name: "Check 6", args: args{n: 6}, want: false},
		{name: "Check 7", args: args{n: 7}, want: true},
		{name: "Check 2,147,483,647", args: args{2147483647}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FermatPrimality(tt.args.n); got != tt.want {
				t.Errorf("FermatPrimality() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkFermatPrimality(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FermatPrimality(2147483647)
	}
}

func TestMillerRabinPrimality(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Check 2", args: args{x: 2}, want: true},
		{name: "Check 3", args: args{x: 3}, want: true},
		{name: "Check 4", args: args{x: 4}, want: false},
		{name: "Check 5", args: args{x: 5}, want: true},
		{name: "Check 6", args: args{x: 6}, want: false},
		{name: "Check 7", args: args{x: 7}, want: true},
		{name: "Check 2,147,483,647", args: args{2147483647}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MillerRabinPrimality(tt.args.x); got != tt.want {
				t.Errorf("MillerRabinPrimality() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkMillerRabinPrimality(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MillerRabinPrimality(2147483647)
	}
}

func TestFindPrimeToLeft(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "Check 55", args: args{x: 55}, want: 53},
		{name: "Check 2", args: args{x: 2}, want: 0},
		{name: "Check 3", args: args{x: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPrimeToLeft(tt.args.x); got != tt.want {
				t.Errorf("FindPrimeToLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
