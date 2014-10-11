package srime_test

import (
	. "github.com/otiai10/mint"
	"github.com/otiai10/srime"
	"testing"
)

func TestFindPrimesUntil(t *testing.T) {
	Expect(t, srime.FindPrimesUntil(10)).TypeOf("[]int")
	Expect(t, srime.FindPrimesUntil(10)).ToBe(
		[]int{2, 3, 5, 7},
	)
}

func TestFactorize(t *testing.T) {
	Expect(t, srime.Factorize(10).List()).ToBe(
		[]int{2, 5},
	)
	Expect(t, srime.Factorize(10).Dict()).ToBe(
		map[int]int{2: 2, 5: 5},
	)

	Expect(t, srime.Factorize(351).List()).ToBe(
		[]int{3, 13},
	)
}
