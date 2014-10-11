package sprime_test

import (
	. "github.com/otiai10/mint"
	"github.com/otiai10/sprime"
	"testing"
)

func TestFindPrimesUntil(t *testing.T) {
	Expect(t, sprime.FindPrimesUntil(10)).TypeOf("[]int")
	Expect(t, sprime.FindPrimesUntil(10)).ToBe(
		[]int{2, 3, 5, 7},
	)
}

func TestFactorize(t *testing.T) {
	Expect(t, sprime.Factorize(10).List()).ToBe(
		[]int{2, 5},
	)
	Expect(t, sprime.Factorize(10).Dict()).ToBe(
		map[int]int{2: 1, 5: 1},
	)

	Expect(t, sprime.Factorize(351).List()).ToBe(
		[]int{3, 13},
	)
}

func TestReduce(t *testing.T) {
	Expect(t, sprime.Reduce(144, 360).String()).ToBe("2/5")
}

func TestFraction_Reduce(t *testing.T) {
	Expect(t, sprime.Fraction{144, 360}.Reduce().String()).ToBe("2/5")
}
