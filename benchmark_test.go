package primes

import "testing"

func BenchmarkUntil(b *testing.B) {
	Globally.Clear()
	for i := 0; i < b.N; i++ {
		for n := 123450; n < 123456; n++ {
			Until(int64(n))
		}
	}
}

func BenchmarkGlobally(b *testing.B) {
	Globally.Clear()
	for i := 0; i < b.N; i++ {
		for n := 123450; n < 123456; n++ {
			Globally.Until(int64(n))
		}
	}
}

func BenchmarkFactorize(b *testing.B) {
	Globally.Clear()
	for i := 0; i < b.N; i++ {
		for n := 123450; n < 123456; n++ {
			Factorize(int64(n))
		}
	}
}
