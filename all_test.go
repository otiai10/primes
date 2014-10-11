package srime_test

import (
	. "github.com/otiai10/mint"
	"github.com/otiai10/srime"
	"testing"
)

func TestPrime(t *testing.T) {
	Expect(t, srime.FindPrimesUntil(10)).TypeOf("[]int")
	Expect(t, srime.FindPrimesUntil(10)).ToBe(
		[]int{2, 3, 5, 7},
	)
}
