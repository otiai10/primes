package primes

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkUntil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range fixtureTargets(true) {
			Until(n)
		}
	}
}

func BenchmarkGlobally(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range fixtureTargets(true) {
			Globally.Until(n)
		}
	}
}

func BenchmarkFactorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorize(123456)
	}
}

var sharedFixture []int64

func fixtureTargets(refresh ...bool) []int64 {
	refresh = append(refresh, false)
	if sharedFixture != nil || refresh[0] {
		rand.Seed(time.Now().Unix())
		for i := 0; i < 10; i++ {
			sharedFixture = append(sharedFixture, int64(rand.Intn(123456)))
		}
	}
	return sharedFixture
}
