package primes_test

import (
	"testing"

	"github.com/otiai10/primes"
)

func BenchmarkUntilWithoutKnowledge(b *testing.B) {
	primes.Ask(new(primes.EmptyKnowledge))
	for i := 0; i < b.N; i++ {
		for n := 1234; n < 1240; n++ {
			primes.Until(int64(n))
		}
	}
	primes.Forget()
}

func BenchmarkUntilWithOnMemoryKnowledge(b *testing.B) {
	primes.Ask(primes.NewOnMemoryKnowledge())
	for i := 0; i < b.N; i++ {
		for n := 1234; n < 1240; n++ {
			primes.Until(int64(n))
		}
	}
	primes.Forget()
}

func BenchmarkFactorizeWithoutKnowledge(b *testing.B) {
	primes.Ask(new(primes.EmptyKnowledge))
	for i := 0; i < b.N; i++ {
		for n := 1234; n < 1240; n++ {
			primes.Factorize(int64(n))
		}
	}
	primes.Forget()
}

func BenchmarkFactorizeWithOnmemoryKnowledge(b *testing.B) {
	primes.Ask(primes.NewOnMemoryKnowledge())
	for i := 0; i < b.N; i++ {
		for n := 1234; n < 1240; n++ {
			primes.Factorize(int64(n))
		}
	}
	primes.Forget()
}
