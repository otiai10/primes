package primes

import "testing"

func BenchmarkUntil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Globally.Clear()
		// UseKnowledge(new(OnMemoryKnowledge))
		if err := UseKnowledge(new(OnSQLiteKnowledge)); err != nil {
			b.Fatal(err)
		}
		// UseKnowledge(new(OnMemoryKnowledge))
		for n := 4000; n < 4010; n++ {
			Until(int64(n))
		}
	}
}

func BenchmarkGlobally(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Globally.Clear()
		// UseKnowledge(new(OnMemoryKnowledge))
		for n := 1234; n < 1240; n++ {
			// Globally.Until(int64(n))
			knowledge.Until(int64(n))
		}
	}
}

func BenchmarkFactorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Globally.Clear()
		// knowledge.Clear()
		for n := 1234; n < 1240; n++ {
			Factorize(int64(n))
		}
	}
}
